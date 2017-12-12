///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/event-manager/gen/models"
)

// GetSubscriptionOKCode is the HTTP code returned for type GetSubscriptionOK
const GetSubscriptionOKCode int = 200

/*GetSubscriptionOK Successful operation

swagger:response getSubscriptionOK
*/
type GetSubscriptionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Subscription `json:"body,omitempty"`
}

// NewGetSubscriptionOK creates GetSubscriptionOK with default headers values
func NewGetSubscriptionOK() *GetSubscriptionOK {
	return &GetSubscriptionOK{}
}

// WithPayload adds the payload to the get subscription o k response
func (o *GetSubscriptionOK) WithPayload(payload *models.Subscription) *GetSubscriptionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get subscription o k response
func (o *GetSubscriptionOK) SetPayload(payload *models.Subscription) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSubscriptionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSubscriptionBadRequestCode is the HTTP code returned for type GetSubscriptionBadRequest
const GetSubscriptionBadRequestCode int = 400

/*GetSubscriptionBadRequest Invalid Name supplied

swagger:response getSubscriptionBadRequest
*/
type GetSubscriptionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSubscriptionBadRequest creates GetSubscriptionBadRequest with default headers values
func NewGetSubscriptionBadRequest() *GetSubscriptionBadRequest {
	return &GetSubscriptionBadRequest{}
}

// WithPayload adds the payload to the get subscription bad request response
func (o *GetSubscriptionBadRequest) WithPayload(payload *models.Error) *GetSubscriptionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get subscription bad request response
func (o *GetSubscriptionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSubscriptionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSubscriptionNotFoundCode is the HTTP code returned for type GetSubscriptionNotFound
const GetSubscriptionNotFoundCode int = 404

/*GetSubscriptionNotFound Subscription not found

swagger:response getSubscriptionNotFound
*/
type GetSubscriptionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSubscriptionNotFound creates GetSubscriptionNotFound with default headers values
func NewGetSubscriptionNotFound() *GetSubscriptionNotFound {
	return &GetSubscriptionNotFound{}
}

// WithPayload adds the payload to the get subscription not found response
func (o *GetSubscriptionNotFound) WithPayload(payload *models.Error) *GetSubscriptionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get subscription not found response
func (o *GetSubscriptionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSubscriptionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSubscriptionInternalServerErrorCode is the HTTP code returned for type GetSubscriptionInternalServerError
const GetSubscriptionInternalServerErrorCode int = 500

/*GetSubscriptionInternalServerError Internal server error

swagger:response getSubscriptionInternalServerError
*/
type GetSubscriptionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSubscriptionInternalServerError creates GetSubscriptionInternalServerError with default headers values
func NewGetSubscriptionInternalServerError() *GetSubscriptionInternalServerError {
	return &GetSubscriptionInternalServerError{}
}

// WithPayload adds the payload to the get subscription internal server error response
func (o *GetSubscriptionInternalServerError) WithPayload(payload *models.Error) *GetSubscriptionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get subscription internal server error response
func (o *GetSubscriptionInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSubscriptionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSubscriptionDefault Unknown error

swagger:response getSubscriptionDefault
*/
type GetSubscriptionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSubscriptionDefault creates GetSubscriptionDefault with default headers values
func NewGetSubscriptionDefault(code int) *GetSubscriptionDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSubscriptionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get subscription default response
func (o *GetSubscriptionDefault) WithStatusCode(code int) *GetSubscriptionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get subscription default response
func (o *GetSubscriptionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get subscription default response
func (o *GetSubscriptionDefault) WithPayload(payload *models.Error) *GetSubscriptionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get subscription default response
func (o *GetSubscriptionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSubscriptionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}