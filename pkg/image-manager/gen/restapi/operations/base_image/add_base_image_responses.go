///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package base_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/image-manager/gen/models"
)

// AddBaseImageCreatedCode is the HTTP code returned for type AddBaseImageCreated
const AddBaseImageCreatedCode int = 201

/*AddBaseImageCreated created

swagger:response addBaseImageCreated
*/
type AddBaseImageCreated struct {

	/*
	  In: Body
	*/
	Payload *models.BaseImage `json:"body,omitempty"`
}

// NewAddBaseImageCreated creates AddBaseImageCreated with default headers values
func NewAddBaseImageCreated() *AddBaseImageCreated {
	return &AddBaseImageCreated{}
}

// WithPayload adds the payload to the add base image created response
func (o *AddBaseImageCreated) WithPayload(payload *models.BaseImage) *AddBaseImageCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add base image created response
func (o *AddBaseImageCreated) SetPayload(payload *models.BaseImage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddBaseImageCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddBaseImageBadRequestCode is the HTTP code returned for type AddBaseImageBadRequest
const AddBaseImageBadRequestCode int = 400

/*AddBaseImageBadRequest Invalid input

swagger:response addBaseImageBadRequest
*/
type AddBaseImageBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddBaseImageBadRequest creates AddBaseImageBadRequest with default headers values
func NewAddBaseImageBadRequest() *AddBaseImageBadRequest {
	return &AddBaseImageBadRequest{}
}

// WithPayload adds the payload to the add base image bad request response
func (o *AddBaseImageBadRequest) WithPayload(payload *models.Error) *AddBaseImageBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add base image bad request response
func (o *AddBaseImageBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddBaseImageBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddBaseImageDefault Generic error response

swagger:response addBaseImageDefault
*/
type AddBaseImageDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddBaseImageDefault creates AddBaseImageDefault with default headers values
func NewAddBaseImageDefault(code int) *AddBaseImageDefault {
	if code <= 0 {
		code = 500
	}

	return &AddBaseImageDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add base image default response
func (o *AddBaseImageDefault) WithStatusCode(code int) *AddBaseImageDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add base image default response
func (o *AddBaseImageDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add base image default response
func (o *AddBaseImageDefault) WithPayload(payload *models.Error) *AddBaseImageDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add base image default response
func (o *AddBaseImageDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddBaseImageDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
