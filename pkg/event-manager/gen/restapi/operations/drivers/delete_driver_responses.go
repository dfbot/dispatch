///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/event-manager/gen/models"
)

// DeleteDriverOKCode is the HTTP code returned for type DeleteDriverOK
const DeleteDriverOKCode int = 200

/*DeleteDriverOK successful operation

swagger:response deleteDriverOK
*/
type DeleteDriverOK struct {

	/*
	  In: Body
	*/
	Payload *models.Driver `json:"body,omitempty"`
}

// NewDeleteDriverOK creates DeleteDriverOK with default headers values
func NewDeleteDriverOK() *DeleteDriverOK {
	return &DeleteDriverOK{}
}

// WithPayload adds the payload to the delete driver o k response
func (o *DeleteDriverOK) WithPayload(payload *models.Driver) *DeleteDriverOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete driver o k response
func (o *DeleteDriverOK) SetPayload(payload *models.Driver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDriverOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDriverBadRequestCode is the HTTP code returned for type DeleteDriverBadRequest
const DeleteDriverBadRequestCode int = 400

/*DeleteDriverBadRequest Invalid ID supplied

swagger:response deleteDriverBadRequest
*/
type DeleteDriverBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDriverBadRequest creates DeleteDriverBadRequest with default headers values
func NewDeleteDriverBadRequest() *DeleteDriverBadRequest {
	return &DeleteDriverBadRequest{}
}

// WithPayload adds the payload to the delete driver bad request response
func (o *DeleteDriverBadRequest) WithPayload(payload *models.Error) *DeleteDriverBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete driver bad request response
func (o *DeleteDriverBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDriverBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDriverNotFoundCode is the HTTP code returned for type DeleteDriverNotFound
const DeleteDriverNotFoundCode int = 404

/*DeleteDriverNotFound Driver not found

swagger:response deleteDriverNotFound
*/
type DeleteDriverNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDriverNotFound creates DeleteDriverNotFound with default headers values
func NewDeleteDriverNotFound() *DeleteDriverNotFound {
	return &DeleteDriverNotFound{}
}

// WithPayload adds the payload to the delete driver not found response
func (o *DeleteDriverNotFound) WithPayload(payload *models.Error) *DeleteDriverNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete driver not found response
func (o *DeleteDriverNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDriverNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDriverInternalServerErrorCode is the HTTP code returned for type DeleteDriverInternalServerError
const DeleteDriverInternalServerErrorCode int = 500

/*DeleteDriverInternalServerError Internal server error

swagger:response deleteDriverInternalServerError
*/
type DeleteDriverInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDriverInternalServerError creates DeleteDriverInternalServerError with default headers values
func NewDeleteDriverInternalServerError() *DeleteDriverInternalServerError {
	return &DeleteDriverInternalServerError{}
}

// WithPayload adds the payload to the delete driver internal server error response
func (o *DeleteDriverInternalServerError) WithPayload(payload *models.Error) *DeleteDriverInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete driver internal server error response
func (o *DeleteDriverInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDriverInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteDriverDefault Generic error response

swagger:response deleteDriverDefault
*/
type DeleteDriverDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDriverDefault creates DeleteDriverDefault with default headers values
func NewDeleteDriverDefault(code int) *DeleteDriverDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteDriverDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete driver default response
func (o *DeleteDriverDefault) WithStatusCode(code int) *DeleteDriverDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete driver default response
func (o *DeleteDriverDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete driver default response
func (o *DeleteDriverDefault) WithPayload(payload *models.Error) *DeleteDriverDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete driver default response
func (o *DeleteDriverDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDriverDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
