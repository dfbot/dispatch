///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package runner

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

// NewGetRunsParams creates a new GetRunsParams object
// with the default values initialized.
func NewGetRunsParams() *GetRunsParams {

	return &GetRunsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunsParamsWithTimeout creates a new GetRunsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunsParamsWithTimeout(timeout time.Duration) *GetRunsParams {

	return &GetRunsParams{

		timeout: timeout,
	}
}

// NewGetRunsParamsWithContext creates a new GetRunsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunsParamsWithContext(ctx context.Context) *GetRunsParams {

	return &GetRunsParams{

		Context: ctx,
	}
}

// NewGetRunsParamsWithHTTPClient creates a new GetRunsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunsParamsWithHTTPClient(client *http.Client) *GetRunsParams {

	return &GetRunsParams{
		HTTPClient: client,
	}
}

/*GetRunsParams contains all the parameters to send to the API endpoint
for the get runs operation typically these are written to a http.Request
*/
type GetRunsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get runs params
func (o *GetRunsParams) WithTimeout(timeout time.Duration) *GetRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runs params
func (o *GetRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runs params
func (o *GetRunsParams) WithContext(ctx context.Context) *GetRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runs params
func (o *GetRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runs params
func (o *GetRunsParams) WithHTTPClient(client *http.Client) *GetRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runs params
func (o *GetRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
