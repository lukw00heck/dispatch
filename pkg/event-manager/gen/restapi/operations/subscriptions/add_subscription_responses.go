///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/event-manager/gen/models"
)

// AddSubscriptionCreatedCode is the HTTP code returned for type AddSubscriptionCreated
const AddSubscriptionCreatedCode int = 201

/*AddSubscriptionCreated Subscription created

swagger:response addSubscriptionCreated
*/
type AddSubscriptionCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Subscription `json:"body,omitempty"`
}

// NewAddSubscriptionCreated creates AddSubscriptionCreated with default headers values
func NewAddSubscriptionCreated() *AddSubscriptionCreated {
	return &AddSubscriptionCreated{}
}

// WithPayload adds the payload to the add subscription created response
func (o *AddSubscriptionCreated) WithPayload(payload *models.Subscription) *AddSubscriptionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add subscription created response
func (o *AddSubscriptionCreated) SetPayload(payload *models.Subscription) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSubscriptionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddSubscriptionBadRequestCode is the HTTP code returned for type AddSubscriptionBadRequest
const AddSubscriptionBadRequestCode int = 400

/*AddSubscriptionBadRequest Invalid input

swagger:response addSubscriptionBadRequest
*/
type AddSubscriptionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddSubscriptionBadRequest creates AddSubscriptionBadRequest with default headers values
func NewAddSubscriptionBadRequest() *AddSubscriptionBadRequest {
	return &AddSubscriptionBadRequest{}
}

// WithPayload adds the payload to the add subscription bad request response
func (o *AddSubscriptionBadRequest) WithPayload(payload *models.Error) *AddSubscriptionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add subscription bad request response
func (o *AddSubscriptionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSubscriptionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddSubscriptionUnauthorizedCode is the HTTP code returned for type AddSubscriptionUnauthorized
const AddSubscriptionUnauthorizedCode int = 401

/*AddSubscriptionUnauthorized Unauthorized Request

swagger:response addSubscriptionUnauthorized
*/
type AddSubscriptionUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddSubscriptionUnauthorized creates AddSubscriptionUnauthorized with default headers values
func NewAddSubscriptionUnauthorized() *AddSubscriptionUnauthorized {
	return &AddSubscriptionUnauthorized{}
}

// WithPayload adds the payload to the add subscription unauthorized response
func (o *AddSubscriptionUnauthorized) WithPayload(payload *models.Error) *AddSubscriptionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add subscription unauthorized response
func (o *AddSubscriptionUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSubscriptionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddSubscriptionConflictCode is the HTTP code returned for type AddSubscriptionConflict
const AddSubscriptionConflictCode int = 409

/*AddSubscriptionConflict Already Exists

swagger:response addSubscriptionConflict
*/
type AddSubscriptionConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddSubscriptionConflict creates AddSubscriptionConflict with default headers values
func NewAddSubscriptionConflict() *AddSubscriptionConflict {
	return &AddSubscriptionConflict{}
}

// WithPayload adds the payload to the add subscription conflict response
func (o *AddSubscriptionConflict) WithPayload(payload *models.Error) *AddSubscriptionConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add subscription conflict response
func (o *AddSubscriptionConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSubscriptionConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddSubscriptionInternalServerErrorCode is the HTTP code returned for type AddSubscriptionInternalServerError
const AddSubscriptionInternalServerErrorCode int = 500

/*AddSubscriptionInternalServerError Internal server error

swagger:response addSubscriptionInternalServerError
*/
type AddSubscriptionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddSubscriptionInternalServerError creates AddSubscriptionInternalServerError with default headers values
func NewAddSubscriptionInternalServerError() *AddSubscriptionInternalServerError {
	return &AddSubscriptionInternalServerError{}
}

// WithPayload adds the payload to the add subscription internal server error response
func (o *AddSubscriptionInternalServerError) WithPayload(payload *models.Error) *AddSubscriptionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add subscription internal server error response
func (o *AddSubscriptionInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSubscriptionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddSubscriptionDefault Unknown error

swagger:response addSubscriptionDefault
*/
type AddSubscriptionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddSubscriptionDefault creates AddSubscriptionDefault with default headers values
func NewAddSubscriptionDefault(code int) *AddSubscriptionDefault {
	if code <= 0 {
		code = 500
	}

	return &AddSubscriptionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add subscription default response
func (o *AddSubscriptionDefault) WithStatusCode(code int) *AddSubscriptionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add subscription default response
func (o *AddSubscriptionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add subscription default response
func (o *AddSubscriptionDefault) WithPayload(payload *models.Error) *AddSubscriptionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add subscription default response
func (o *AddSubscriptionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSubscriptionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
