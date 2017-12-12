///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAPIParams creates a new GetAPIParams object
// with the default values initialized.
func NewGetAPIParams() GetAPIParams {
	var ()
	return GetAPIParams{}
}

// GetAPIParams contains all the bound params for the get API operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAPI
type GetAPIParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Name of API to work on
	  Required: true
	  Pattern: ^[\w\d\-]+$
	  In: path
	*/
	API string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetAPIParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rAPI, rhkAPI, _ := route.Params.GetOK("api")
	if err := o.bindAPI(rAPI, rhkAPI, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAPIParams) bindAPI(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.API = raw

	if err := o.validateAPI(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetAPIParams) validateAPI(formats strfmt.Registry) error {

	if err := validate.Pattern("api", "path", o.API, `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}
