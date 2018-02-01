///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/application-manager/gen/models"
)

// GetAppsReader is a Reader for the GetApps structure.
type GetAppsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAppsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetAppsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetAppsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAppsOK creates a GetAppsOK with default headers values
func NewGetAppsOK() *GetAppsOK {
	return &GetAppsOK{}
}

/*GetAppsOK handles this case with default header values.

Successful operation
*/
type GetAppsOK struct {
	Payload models.GetAppsOKBody
}

func (o *GetAppsOK) Error() string {
	return fmt.Sprintf("[GET /][%d] getAppsOK  %+v", 200, o.Payload)
}

func (o *GetAppsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppsInternalServerError creates a GetAppsInternalServerError with default headers values
func NewGetAppsInternalServerError() *GetAppsInternalServerError {
	return &GetAppsInternalServerError{}
}

/*GetAppsInternalServerError handles this case with default header values.

Internal Error
*/
type GetAppsInternalServerError struct {
	Payload *models.Error
}

func (o *GetAppsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /][%d] getAppsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppsDefault creates a GetAppsDefault with default headers values
func NewGetAppsDefault(code int) *GetAppsDefault {
	return &GetAppsDefault{
		_statusCode: code,
	}
}

/*GetAppsDefault handles this case with default header values.

Unexpected Error
*/
type GetAppsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get apps default response
func (o *GetAppsDefault) Code() int {
	return o._statusCode
}

func (o *GetAppsDefault) Error() string {
	return fmt.Sprintf("[GET /][%d] getApps default  %+v", o._statusCode, o.Payload)
}

func (o *GetAppsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
