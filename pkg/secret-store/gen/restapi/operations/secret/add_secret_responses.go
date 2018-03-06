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

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/secret-store/gen/models"
)

// AddSecretCreatedCode is the HTTP code returned for type AddSecretCreated
const AddSecretCreatedCode int = 201

/*AddSecretCreated The created secret.

swagger:response addSecretCreated
*/
type AddSecretCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Secret `json:"body,omitempty"`
}

// NewAddSecretCreated creates AddSecretCreated with default headers values
func NewAddSecretCreated() *AddSecretCreated {
	return &AddSecretCreated{}
}

// WithPayload adds the payload to the add secret created response
func (o *AddSecretCreated) WithPayload(payload *models.Secret) *AddSecretCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add secret created response
func (o *AddSecretCreated) SetPayload(payload *models.Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSecretCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddSecretBadRequestCode is the HTTP code returned for type AddSecretBadRequest
const AddSecretBadRequestCode int = 400

/*AddSecretBadRequest Bad Request

swagger:response addSecretBadRequest
*/
type AddSecretBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddSecretBadRequest creates AddSecretBadRequest with default headers values
func NewAddSecretBadRequest() *AddSecretBadRequest {
	return &AddSecretBadRequest{}
}

// WithPayload adds the payload to the add secret bad request response
func (o *AddSecretBadRequest) WithPayload(payload *models.Error) *AddSecretBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add secret bad request response
func (o *AddSecretBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSecretBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddSecretConflictCode is the HTTP code returned for type AddSecretConflict
const AddSecretConflictCode int = 409

/*AddSecretConflict Already Exists

swagger:response addSecretConflict
*/
type AddSecretConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddSecretConflict creates AddSecretConflict with default headers values
func NewAddSecretConflict() *AddSecretConflict {
	return &AddSecretConflict{}
}

// WithPayload adds the payload to the add secret conflict response
func (o *AddSecretConflict) WithPayload(payload *models.Error) *AddSecretConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add secret conflict response
func (o *AddSecretConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSecretConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddSecretDefault Standard error

swagger:response addSecretDefault
*/
type AddSecretDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddSecretDefault creates AddSecretDefault with default headers values
func NewAddSecretDefault(code int) *AddSecretDefault {
	if code <= 0 {
		code = 500
	}

	return &AddSecretDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add secret default response
func (o *AddSecretDefault) WithStatusCode(code int) *AddSecretDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add secret default response
func (o *AddSecretDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add secret default response
func (o *AddSecretDefault) WithPayload(payload *models.Error) *AddSecretDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add secret default response
func (o *AddSecretDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSecretDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
