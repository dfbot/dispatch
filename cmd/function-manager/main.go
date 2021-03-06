///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package main

import (
	"os"
	"time"

	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/boltdb"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/loads/fmts"
	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"
	"github.com/justinas/alice"
	log "github.com/sirupsen/logrus"

	"github.com/vmware/dispatch/pkg/config"
	"github.com/vmware/dispatch/pkg/entity-store"
	"github.com/vmware/dispatch/pkg/function-manager"
	"github.com/vmware/dispatch/pkg/function-manager/gen/restapi"
	"github.com/vmware/dispatch/pkg/function-manager/gen/restapi/operations"
	"github.com/vmware/dispatch/pkg/functions"
	"github.com/vmware/dispatch/pkg/functions/openfaas"
	"github.com/vmware/dispatch/pkg/functions/openwhisk"
	"github.com/vmware/dispatch/pkg/functions/riff"
	"github.com/vmware/dispatch/pkg/functions/runner"
	"github.com/vmware/dispatch/pkg/functions/secretinjector"
	"github.com/vmware/dispatch/pkg/functions/validator"
	"github.com/vmware/dispatch/pkg/middleware"
	"github.com/vmware/dispatch/pkg/trace"
)

var drivers = map[string]func() functions.FaaSDriver{
	"openfaas": func() functions.FaaSDriver {
		faas, err := openfaas.New(&openfaas.Config{
			Gateway:       config.Global.OpenFaas.Gateway,
			ImageRegistry: config.Global.OpenFaas.ImageRegistry,
			RegistryAuth:  config.Global.OpenFaas.RegistryAuth,
		})
		if err != nil {
			log.Fatalf("Error starting OpenFaaS driver: %+v", err)
		}
		return faas
	},
	"riff": func() functions.FaaSDriver {
		faas, err := riff.New(&riff.Config{
			ImageRegistry: config.Global.Riff.ImageRegistry,
			Gateway:       config.Global.Riff.Gateway,
			RegistryAuth:  config.Global.Riff.RegistryAuth,
			K8sConfig:     config.Global.Riff.K8sConfig,
			RiffNamespace: config.Global.Riff.RiffNamespace,
		})
		if err != nil {
			log.Fatalf("Error starting OpenFaaS driver: %+v", err)
		}
		return faas
	},
	"openwhisk": func() functions.FaaSDriver {
		faas, err := openwhisk.New(&openwhisk.Config{
			AuthToken: config.Global.Openwhisk.AuthToken,
			Host:      config.Global.Openwhisk.Host,
			Insecure:  true,
		})
		if err != nil {
			log.Fatalf("Error getting OpenWhisk client: %+v", err)
		}
		return faas
	},
}

func init() {
	loads.AddLoader(fmts.YAMLMatcher, fmts.YAMLDoc)
	boltdb.Register()
}

var debugFlags = struct {
	DebugEnabled   bool `long:"debug" description:"Enable debugging messages"`
	TracingEnabled bool `long:"trace" description:"Enable tracing messages (enables debugging)"`
}{}

func configureFlags() []swag.CommandLineOptionsGroup {
	return []swag.CommandLineOptionsGroup{{
		ShortDescription: "Function manager Flags",
		LongDescription:  "",
		Options:          &functionmanager.FunctionManagerFlags,
	}, {
		ShortDescription: "Debug options",
		LongDescription:  "",
		Options:          &debugFlags,
	}}
}

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "2.0")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewFunctionManagerAPI(swaggerSpec)
	server := restapi.NewServer(api)

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Function Manager"
	parser.LongDescription = "This is the API server for the Dispatch Function Manager service.\n"

	optsGroups := configureFlags()
	for _, optsGroup := range optsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	if debugFlags.DebugEnabled {
		log.SetLevel(log.DebugLevel)
	}
	if debugFlags.TracingEnabled {
		log.SetLevel(log.DebugLevel)
		trace.Enable()
	}

	config.Global = config.LoadConfiguration(functionmanager.FunctionManagerFlags.Config)

	kv, err := libkv.NewStore(
		store.BOLTDB,
		[]string{functionmanager.FunctionManagerFlags.DbFile},
		&store.Config{
			Bucket:            "function",
			ConnectionTimeout: 1 * time.Second,
			PersistConnection: true,
		},
	)
	if err != nil {
		log.Fatalf("Error creating/opening the entity store: %v", err)
	}
	faas := drivers[functionmanager.FunctionManagerFlags.Faas]()

	es := entitystore.New(kv)

	c := &functionmanager.ControllerConfig{
		ResyncPeriod:   time.Duration(functionmanager.FunctionManagerFlags.ResyncPeriod) * time.Second,
		OrganizationID: functionmanager.FunctionManagerFlags.OrgID,
	}
	r := runner.New(&runner.Config{
		Faas:           faas,
		Validator:      validator.New(),
		SecretInjector: secretinjector.New(functionmanager.SecretStoreClient()),
	})

	controller := functionmanager.NewController(c, es, faas, r, functionmanager.ImageManagerClient())
	defer controller.Shutdown()
	controller.Start()

	handlers := functionmanager.NewHandlers(controller.Watcher(), es)
	handlers.ConfigureHandlers(api)

	healthChecker := func() error {
		// TODO: implement service-specific healthchecking
		return nil
	}

	handler := alice.New(
		middleware.NewHealthCheckMW("", healthChecker),
	).Then(api.Serve(nil))

	server.SetHandler(handler)

	defer server.Shutdown()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
