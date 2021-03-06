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

	"github.com/go-openapi/runtime"

	models "github.com/vmware/dispatch/pkg/api-manager/gen/models"
)

// AddAPIOKCode is the HTTP code returned for type AddAPIOK
const AddAPIOKCode int = 200

/*AddAPIOK API created

swagger:response addApiOK
*/
type AddAPIOK struct {

	/*
	  In: Body
	*/
	Payload *models.API `json:"body,omitempty"`
}

// NewAddAPIOK creates AddAPIOK with default headers values
func NewAddAPIOK() *AddAPIOK {

	return &AddAPIOK{}
}

// WithPayload adds the payload to the add Api o k response
func (o *AddAPIOK) WithPayload(payload *models.API) *AddAPIOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add Api o k response
func (o *AddAPIOK) SetPayload(payload *models.API) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAPIOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAPIBadRequestCode is the HTTP code returned for type AddAPIBadRequest
const AddAPIBadRequestCode int = 400

/*AddAPIBadRequest Invalid Input

swagger:response addApiBadRequest
*/
type AddAPIBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddAPIBadRequest creates AddAPIBadRequest with default headers values
func NewAddAPIBadRequest() *AddAPIBadRequest {

	return &AddAPIBadRequest{}
}

// WithPayload adds the payload to the add Api bad request response
func (o *AddAPIBadRequest) WithPayload(payload *models.Error) *AddAPIBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add Api bad request response
func (o *AddAPIBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAPIBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAPIUnauthorizedCode is the HTTP code returned for type AddAPIUnauthorized
const AddAPIUnauthorizedCode int = 401

/*AddAPIUnauthorized Unauthorized Request

swagger:response addApiUnauthorized
*/
type AddAPIUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddAPIUnauthorized creates AddAPIUnauthorized with default headers values
func NewAddAPIUnauthorized() *AddAPIUnauthorized {

	return &AddAPIUnauthorized{}
}

// WithPayload adds the payload to the add Api unauthorized response
func (o *AddAPIUnauthorized) WithPayload(payload *models.Error) *AddAPIUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add Api unauthorized response
func (o *AddAPIUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAPIUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAPIConflictCode is the HTTP code returned for type AddAPIConflict
const AddAPIConflictCode int = 409

/*AddAPIConflict Already Exists

swagger:response addApiConflict
*/
type AddAPIConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddAPIConflict creates AddAPIConflict with default headers values
func NewAddAPIConflict() *AddAPIConflict {

	return &AddAPIConflict{}
}

// WithPayload adds the payload to the add Api conflict response
func (o *AddAPIConflict) WithPayload(payload *models.Error) *AddAPIConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add Api conflict response
func (o *AddAPIConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAPIConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAPIInternalServerErrorCode is the HTTP code returned for type AddAPIInternalServerError
const AddAPIInternalServerErrorCode int = 500

/*AddAPIInternalServerError Internal Error

swagger:response addApiInternalServerError
*/
type AddAPIInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddAPIInternalServerError creates AddAPIInternalServerError with default headers values
func NewAddAPIInternalServerError() *AddAPIInternalServerError {

	return &AddAPIInternalServerError{}
}

// WithPayload adds the payload to the add Api internal server error response
func (o *AddAPIInternalServerError) WithPayload(payload *models.Error) *AddAPIInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add Api internal server error response
func (o *AddAPIInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAPIInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
