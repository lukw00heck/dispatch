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

// AddDriverTypeCreatedCode is the HTTP code returned for type AddDriverTypeCreated
const AddDriverTypeCreatedCode int = 201

/*AddDriverTypeCreated Driver Type created

swagger:response addDriverTypeCreated
*/
type AddDriverTypeCreated struct {

	/*
	  In: Body
	*/
	Payload *models.DriverType `json:"body,omitempty"`
}

// NewAddDriverTypeCreated creates AddDriverTypeCreated with default headers values
func NewAddDriverTypeCreated() *AddDriverTypeCreated {
	return &AddDriverTypeCreated{}
}

// WithPayload adds the payload to the add driver type created response
func (o *AddDriverTypeCreated) WithPayload(payload *models.DriverType) *AddDriverTypeCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add driver type created response
func (o *AddDriverTypeCreated) SetPayload(payload *models.DriverType) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddDriverTypeCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddDriverTypeBadRequestCode is the HTTP code returned for type AddDriverTypeBadRequest
const AddDriverTypeBadRequestCode int = 400

/*AddDriverTypeBadRequest Invalid input

swagger:response addDriverTypeBadRequest
*/
type AddDriverTypeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddDriverTypeBadRequest creates AddDriverTypeBadRequest with default headers values
func NewAddDriverTypeBadRequest() *AddDriverTypeBadRequest {
	return &AddDriverTypeBadRequest{}
}

// WithPayload adds the payload to the add driver type bad request response
func (o *AddDriverTypeBadRequest) WithPayload(payload *models.Error) *AddDriverTypeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add driver type bad request response
func (o *AddDriverTypeBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddDriverTypeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddDriverTypeUnauthorizedCode is the HTTP code returned for type AddDriverTypeUnauthorized
const AddDriverTypeUnauthorizedCode int = 401

/*AddDriverTypeUnauthorized Unauthorized Request

swagger:response addDriverTypeUnauthorized
*/
type AddDriverTypeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddDriverTypeUnauthorized creates AddDriverTypeUnauthorized with default headers values
func NewAddDriverTypeUnauthorized() *AddDriverTypeUnauthorized {
	return &AddDriverTypeUnauthorized{}
}

// WithPayload adds the payload to the add driver type unauthorized response
func (o *AddDriverTypeUnauthorized) WithPayload(payload *models.Error) *AddDriverTypeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add driver type unauthorized response
func (o *AddDriverTypeUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddDriverTypeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddDriverTypeConflictCode is the HTTP code returned for type AddDriverTypeConflict
const AddDriverTypeConflictCode int = 409

/*AddDriverTypeConflict Already Exists

swagger:response addDriverTypeConflict
*/
type AddDriverTypeConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddDriverTypeConflict creates AddDriverTypeConflict with default headers values
func NewAddDriverTypeConflict() *AddDriverTypeConflict {
	return &AddDriverTypeConflict{}
}

// WithPayload adds the payload to the add driver type conflict response
func (o *AddDriverTypeConflict) WithPayload(payload *models.Error) *AddDriverTypeConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add driver type conflict response
func (o *AddDriverTypeConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddDriverTypeConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddDriverTypeInternalServerErrorCode is the HTTP code returned for type AddDriverTypeInternalServerError
const AddDriverTypeInternalServerErrorCode int = 500

/*AddDriverTypeInternalServerError Internal server error

swagger:response addDriverTypeInternalServerError
*/
type AddDriverTypeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddDriverTypeInternalServerError creates AddDriverTypeInternalServerError with default headers values
func NewAddDriverTypeInternalServerError() *AddDriverTypeInternalServerError {
	return &AddDriverTypeInternalServerError{}
}

// WithPayload adds the payload to the add driver type internal server error response
func (o *AddDriverTypeInternalServerError) WithPayload(payload *models.Error) *AddDriverTypeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add driver type internal server error response
func (o *AddDriverTypeInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddDriverTypeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddDriverTypeDefault Unknown error

swagger:response addDriverTypeDefault
*/
type AddDriverTypeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddDriverTypeDefault creates AddDriverTypeDefault with default headers values
func NewAddDriverTypeDefault(code int) *AddDriverTypeDefault {
	if code <= 0 {
		code = 500
	}

	return &AddDriverTypeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add driver type default response
func (o *AddDriverTypeDefault) WithStatusCode(code int) *AddDriverTypeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add driver type default response
func (o *AddDriverTypeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add driver type default response
func (o *AddDriverTypeDefault) WithPayload(payload *models.Error) *AddDriverTypeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add driver type default response
func (o *AddDriverTypeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddDriverTypeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
