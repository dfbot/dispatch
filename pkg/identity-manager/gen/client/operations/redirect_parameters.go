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

// NewRedirectParams creates a new RedirectParams object
// with the default values initialized.
func NewRedirectParams() *RedirectParams {
	var ()
	return &RedirectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRedirectParamsWithTimeout creates a new RedirectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRedirectParamsWithTimeout(timeout time.Duration) *RedirectParams {
	var ()
	return &RedirectParams{

		timeout: timeout,
	}
}

// NewRedirectParamsWithContext creates a new RedirectParams object
// with the default values initialized, and the ability to set a context for a request
func NewRedirectParamsWithContext(ctx context.Context) *RedirectParams {
	var ()
	return &RedirectParams{

		Context: ctx,
	}
}

// NewRedirectParamsWithHTTPClient creates a new RedirectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRedirectParamsWithHTTPClient(client *http.Client) *RedirectParams {
	var ()
	return &RedirectParams{
		HTTPClient: client,
	}
}

/*RedirectParams contains all the parameters to send to the API endpoint
for the redirect operation typically these are written to a http.Request
*/
type RedirectParams struct {

	/*Redirect
	  the local server url redirecting to

	*/
	Redirect *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the redirect params
func (o *RedirectParams) WithTimeout(timeout time.Duration) *RedirectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the redirect params
func (o *RedirectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the redirect params
func (o *RedirectParams) WithContext(ctx context.Context) *RedirectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the redirect params
func (o *RedirectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the redirect params
func (o *RedirectParams) WithHTTPClient(client *http.Client) *RedirectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the redirect params
func (o *RedirectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRedirect adds the redirect to the redirect params
func (o *RedirectParams) WithRedirect(redirect *string) *RedirectParams {
	o.SetRedirect(redirect)
	return o
}

// SetRedirect adds the redirect to the redirect params
func (o *RedirectParams) SetRedirect(redirect *string) {
	o.Redirect = redirect
}

// WriteToRequest writes these params to a swagger request
func (o *RedirectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Redirect != nil {

		// query param redirect
		var qrRedirect string
		if o.Redirect != nil {
			qrRedirect = *o.Redirect
		}
		qRedirect := qrRedirect
		if qRedirect != "" {
			if err := r.SetQueryParam("redirect", qRedirect); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
