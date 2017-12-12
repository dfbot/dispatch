///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// LoginPasswordHandlerFunc turns a function with the right signature into a login password handler
type LoginPasswordHandlerFunc func(LoginPasswordParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LoginPasswordHandlerFunc) Handle(params LoginPasswordParams) middleware.Responder {
	return fn(params)
}

// LoginPasswordHandler interface for that can handle valid login password params
type LoginPasswordHandler interface {
	Handle(LoginPasswordParams) middleware.Responder
}

// NewLoginPassword creates a new http.Handler for the login password operation
func NewLoginPassword(ctx *middleware.Context, handler LoginPasswordHandler) *LoginPassword {
	return &LoginPassword{Context: ctx, Handler: handler}
}

/*LoginPassword swagger:route GET /v1/iam/login/password authentication loginPassword

user logs in with username and password, the credientials are forwarded to the external identity provider to exchange for auth token

*/
type LoginPassword struct {
	Context *middleware.Context
	Handler LoginPasswordHandler
}

func (o *LoginPassword) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewLoginPasswordParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
