///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package secret

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

// NewDeleteSecretParams creates a new DeleteSecretParams object
// with the default values initialized.
func NewDeleteSecretParams() *DeleteSecretParams {
	var ()
	return &DeleteSecretParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSecretParamsWithTimeout creates a new DeleteSecretParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSecretParamsWithTimeout(timeout time.Duration) *DeleteSecretParams {
	var ()
	return &DeleteSecretParams{

		timeout: timeout,
	}
}

// NewDeleteSecretParamsWithContext creates a new DeleteSecretParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSecretParamsWithContext(ctx context.Context) *DeleteSecretParams {
	var ()
	return &DeleteSecretParams{

		Context: ctx,
	}
}

// NewDeleteSecretParamsWithHTTPClient creates a new DeleteSecretParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSecretParamsWithHTTPClient(client *http.Client) *DeleteSecretParams {
	var ()
	return &DeleteSecretParams{
		HTTPClient: client,
	}
}

/*DeleteSecretParams contains all the parameters to send to the API endpoint
for the delete secret operation typically these are written to a http.Request
*/
type DeleteSecretParams struct {

	/*SecretName*/
	SecretName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete secret params
func (o *DeleteSecretParams) WithTimeout(timeout time.Duration) *DeleteSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete secret params
func (o *DeleteSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete secret params
func (o *DeleteSecretParams) WithContext(ctx context.Context) *DeleteSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete secret params
func (o *DeleteSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete secret params
func (o *DeleteSecretParams) WithHTTPClient(client *http.Client) *DeleteSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete secret params
func (o *DeleteSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSecretName adds the secretName to the delete secret params
func (o *DeleteSecretParams) WithSecretName(secretName string) *DeleteSecretParams {
	o.SetSecretName(secretName)
	return o
}

// SetSecretName adds the secretName to the delete secret params
func (o *DeleteSecretParams) SetSecretName(secretName string) {
	o.SecretName = secretName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param secretName
	if err := r.SetPathParam("secretName", o.SecretName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
