///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package eventmanager

// NO TEST

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-openapi/swag"
	ewrapper "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	corev1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/vmware/dispatch/pkg/errors"
)

type DriverBackend interface {
	Deploy(*Driver) error
	Update(*Driver) error
	Delete(*Driver) error
}

// K8sBackendConfig specify options for Kubernetes driver backend
type K8sBackendConfig struct {
	Replicas             int32
	EnableReadinessProbe bool
	Namespace            string
}

type k8sBackend struct {
	clientset *kubernetes.Clientset
	config    K8sBackendConfig
}

func NewK8sBackend() (DriverBackend, error) {

	var err error
	var config *rest.Config
	if EventManagerFlags.K8sConfig == "" {
		// creates the in-cluster config
		config, err = rest.InClusterConfig()
	} else {
		// create from a configuration
		config, err = clientcmd.BuildConfigFromFlags("", EventManagerFlags.K8sConfig)
	}
	if err != nil {
		return nil, ewrapper.Wrap(err, "Error getting kubernetes config")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, ewrapper.Wrap(err, "Error getting kubernetes clientset")
	}

	return &k8sBackend{
		clientset: clientset,
		config: K8sBackendConfig{
			EnableReadinessProbe: false,
			Replicas:             1,
			Namespace:            EventManagerFlags.K8sNamespace,
		},
	}, nil
}

func getDriverFullName(driver *Driver) string {
	return fmt.Sprintf("event-driver-%s-%s", driver.Type, driver.Name)
}

func (k *k8sBackend) makeDeploymentSpec(driver *Driver) (*v1beta1.Deployment, error) {
	fullname := getDriverFullName(driver)

	// TODO: readiness check
	path := filepath.Join(os.TempDir(), ".lock")
	probe := &corev1.Probe{
		Handler: corev1.Handler{
			Exec: &corev1.ExecAction{
				Command: []string{"cat", path},
			},
		},
		InitialDelaySeconds: 3,
		TimeoutSeconds:      1,
		PeriodSeconds:       10,
		SuccessThreshold:    1,
		FailureThreshold:    3,
	}
	if !k.config.EnableReadinessProbe {
		probe = nil
	}

	args := []string{
		driver.Type,
		fmt.Sprintf("--%s=%s", "amqpurl", EventManagerFlags.AMQPURL),
	}
	for key, val := range driver.Config {
		args = append(args, fmt.Sprintf("--%s=%s", key, val))
	}

	deploymentSpec := &v1beta1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "extensions/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: fullname,
		},
		Spec: v1beta1.DeploymentSpec{
			Replicas: swag.Int32(k.config.Replicas),
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: fullname,
					Labels: map[string]string{
						"app": "event-driver",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            fullname,
							ImagePullPolicy: corev1.PullIfNotPresent,
							Image:           EventManagerFlags.EventDriverImage,
							Ports:           []corev1.ContainerPort{},
							LivenessProbe:   probe,
							Args:            args,
						},
					},
					RestartPolicy: corev1.RestartPolicyAlways,
				},
			},
		},
	}
	return deploymentSpec, nil
}

func (k *k8sBackend) Deploy(driver *Driver) error {

	deploymentSpec, err := k.makeDeploymentSpec(driver)
	if err != nil {
		err = &errors.DriverError{
			Err: ewrapper.Wrapf(err, "k8s: error making a deployment"),
		}
		log.Errorln(err)
		return err
	}

	result, err := k.clientset.Extensions().Deployments(k.config.Namespace).Create(deploymentSpec)
	if err != nil {
		err = &errors.DriverError{
			Err: ewrapper.Wrapf(err, "k8s: error creating a deployment"),
		}
		log.Errorln(err)
		return err
	}

	if output, err := json.MarshalIndent(result, "", "  "); err == nil {
		log.Debugf("k8s: creating deployment\n%s\n", output)
	} else {
		log.Debugf("k8s: json marshal error")
	}

	log.Debugf("k8s: deployment=%s created", getDriverFullName(driver))
	return nil
}

func isEventDriver(deployment *v1beta1.Deployment) bool {
	if deployment != nil {
		val, ok := deployment.Labels["app"]
		if ok && val == "event-driver" {
			return true
		}
	}
	return false
}

func (k *k8sBackend) Delete(driver *Driver) error {

	fullname := getDriverFullName(driver)

	deployment, err := k.clientset.Extensions().Deployments(k.config.Namespace).Get(fullname, metav1.GetOptions{})
	if err != nil {
		if k8serrors.IsNotFound(err) {
			err = &errors.ObjectNotFoundError{
				Err: ewrapper.Wrapf(err, "k8s: deployment=%s not found", fullname),
			}
		} else {
			err = &errors.DriverError{
				Err: ewrapper.Wrapf(err, "k8s: deployment=%s unexpected error"),
			}
		}
		log.Errorln(err)
		return err
	}

	if !isEventDriver(deployment) {
		err = &errors.DriverError{
			Err: ewrapper.Wrapf(err, "k8s: deployment=%s: deleting a NON-event-driver deployment"),
		}
		return err
	}

	foreground := metav1.DeletePropagationForeground
	if err := k.clientset.Extensions().Deployments(k.config.Namespace).Delete(fullname,
		&metav1.DeleteOptions{
			PropagationPolicy: &foreground,
		}); err != nil {

		if k8serrors.IsNotFound(err) {
			err = &errors.ObjectNotFoundError{
				Err: ewrapper.Wrapf(err, "k8s: deployment=%s not found", fullname),
			}
		} else {
			err = &errors.DriverError{
				Err: ewrapper.Wrapf(err, "k8s: deployment=%s unexpected error"),
			}
		}
		log.Errorln(err)
		return err
	}
	return nil
}

func (k *k8sBackend) Update(driver *Driver) error {
	// TODO:
	return fmt.Errorf("Update not implemented yet")
}
