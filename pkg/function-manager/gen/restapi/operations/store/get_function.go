///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetFunctionHandlerFunc turns a function with the right signature into a get function handler
type GetFunctionHandlerFunc func(GetFunctionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFunctionHandlerFunc) Handle(params GetFunctionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetFunctionHandler interface for that can handle valid get function params
type GetFunctionHandler interface {
	Handle(GetFunctionParams, interface{}) middleware.Responder
}

// NewGetFunction creates a new http.Handler for the get function operation
func NewGetFunction(ctx *middleware.Context, handler GetFunctionHandler) *GetFunction {
	return &GetFunction{Context: ctx, Handler: handler}
}

/*GetFunction swagger:route GET /{functionName} Store getFunction

Find function by Name

Returns a single function

*/
type GetFunction struct {
	Context *middleware.Context
	Handler GetFunctionHandler
}

func (o *GetFunction) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFunctionParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}