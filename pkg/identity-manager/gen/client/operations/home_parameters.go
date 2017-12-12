///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewHomeParams creates a new HomeParams object
// with the default values initialized.
func NewHomeParams() *HomeParams {

	return &HomeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHomeParamsWithTimeout creates a new HomeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHomeParamsWithTimeout(timeout time.Duration) *HomeParams {

	return &HomeParams{

		timeout: timeout,
	}
}

// NewHomeParamsWithContext creates a new HomeParams object
// with the default values initialized, and the ability to set a context for a request
func NewHomeParamsWithContext(ctx context.Context) *HomeParams {

	return &HomeParams{

		Context: ctx,
	}
}

// NewHomeParamsWithHTTPClient creates a new HomeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHomeParamsWithHTTPClient(client *http.Client) *HomeParams {

	return &HomeParams{
		HTTPClient: client,
	}
}

/*HomeParams contains all the parameters to send to the API endpoint
for the home operation typically these are written to a http.Request
*/
type HomeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the home params
func (o *HomeParams) WithTimeout(timeout time.Duration) *HomeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the home params
func (o *HomeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the home params
func (o *HomeParams) WithContext(ctx context.Context) *HomeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the home params
func (o *HomeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the home params
func (o *HomeParams) WithHTTPClient(client *http.Client) *HomeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the home params
func (o *HomeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HomeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
