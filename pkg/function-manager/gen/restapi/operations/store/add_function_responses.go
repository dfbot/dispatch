///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/function-manager/gen/models"
)

// AddFunctionOKCode is the HTTP code returned for type AddFunctionOK
const AddFunctionOKCode int = 200

/*AddFunctionOK Function created

swagger:response addFunctionOK
*/
type AddFunctionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Function `json:"body,omitempty"`
}

// NewAddFunctionOK creates AddFunctionOK with default headers values
func NewAddFunctionOK() *AddFunctionOK {
	return &AddFunctionOK{}
}

// WithPayload adds the payload to the add function o k response
func (o *AddFunctionOK) WithPayload(payload *models.Function) *AddFunctionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add function o k response
func (o *AddFunctionOK) SetPayload(payload *models.Function) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddFunctionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddFunctionBadRequestCode is the HTTP code returned for type AddFunctionBadRequest
const AddFunctionBadRequestCode int = 400

/*AddFunctionBadRequest Invalid input

swagger:response addFunctionBadRequest
*/
type AddFunctionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddFunctionBadRequest creates AddFunctionBadRequest with default headers values
func NewAddFunctionBadRequest() *AddFunctionBadRequest {
	return &AddFunctionBadRequest{}
}

// WithPayload adds the payload to the add function bad request response
func (o *AddFunctionBadRequest) WithPayload(payload *models.Error) *AddFunctionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add function bad request response
func (o *AddFunctionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddFunctionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddFunctionUnauthorizedCode is the HTTP code returned for type AddFunctionUnauthorized
const AddFunctionUnauthorizedCode int = 401

/*AddFunctionUnauthorized Unauthorized Request

swagger:response addFunctionUnauthorized
*/
type AddFunctionUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddFunctionUnauthorized creates AddFunctionUnauthorized with default headers values
func NewAddFunctionUnauthorized() *AddFunctionUnauthorized {
	return &AddFunctionUnauthorized{}
}

// WithPayload adds the payload to the add function unauthorized response
func (o *AddFunctionUnauthorized) WithPayload(payload *models.Error) *AddFunctionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add function unauthorized response
func (o *AddFunctionUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddFunctionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddFunctionInternalServerErrorCode is the HTTP code returned for type AddFunctionInternalServerError
const AddFunctionInternalServerErrorCode int = 500

/*AddFunctionInternalServerError Internal error

swagger:response addFunctionInternalServerError
*/
type AddFunctionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddFunctionInternalServerError creates AddFunctionInternalServerError with default headers values
func NewAddFunctionInternalServerError() *AddFunctionInternalServerError {
	return &AddFunctionInternalServerError{}
}

// WithPayload adds the payload to the add function internal server error response
func (o *AddFunctionInternalServerError) WithPayload(payload *models.Error) *AddFunctionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add function internal server error response
func (o *AddFunctionInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddFunctionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
