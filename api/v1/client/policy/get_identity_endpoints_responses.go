// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// GetIdentityEndpointsReader is a Reader for the GetIdentityEndpoints structure.
type GetIdentityEndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentityEndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIdentityEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetIdentityEndpointsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIdentityEndpointsOK creates a GetIdentityEndpointsOK with default headers values
func NewGetIdentityEndpointsOK() *GetIdentityEndpointsOK {
	return &GetIdentityEndpointsOK{}
}

/*GetIdentityEndpointsOK handles this case with default header values.

Success
*/
type GetIdentityEndpointsOK struct {
	Payload []*models.IdentityEndpoints
}

func (o *GetIdentityEndpointsOK) Error() string {
	return fmt.Sprintf("[GET /identity/endpoints][%d] getIdentityEndpointsOK  %+v", 200, o.Payload)
}

func (o *GetIdentityEndpointsOK) GetPayload() []*models.IdentityEndpoints {
	return o.Payload
}

func (o *GetIdentityEndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityEndpointsNotFound creates a GetIdentityEndpointsNotFound with default headers values
func NewGetIdentityEndpointsNotFound() *GetIdentityEndpointsNotFound {
	return &GetIdentityEndpointsNotFound{}
}

/*GetIdentityEndpointsNotFound handles this case with default header values.

Set of identities which are being used by local endpoints could not be found.
*/
type GetIdentityEndpointsNotFound struct {
}

func (o *GetIdentityEndpointsNotFound) Error() string {
	return fmt.Sprintf("[GET /identity/endpoints][%d] getIdentityEndpointsNotFound ", 404)
}

func (o *GetIdentityEndpointsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
