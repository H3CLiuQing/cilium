// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// GetMapNameReader is a Reader for the GetMapName structure.
type GetMapNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMapNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMapNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetMapNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMapNameOK creates a GetMapNameOK with default headers values
func NewGetMapNameOK() *GetMapNameOK {
	return &GetMapNameOK{}
}

/*GetMapNameOK handles this case with default header values.

Success
*/
type GetMapNameOK struct {
	Payload *models.BPFMap
}

func (o *GetMapNameOK) Error() string {
	return fmt.Sprintf("[GET /map/{name}][%d] getMapNameOK  %+v", 200, o.Payload)
}

func (o *GetMapNameOK) GetPayload() *models.BPFMap {
	return o.Payload
}

func (o *GetMapNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BPFMap)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMapNameNotFound creates a GetMapNameNotFound with default headers values
func NewGetMapNameNotFound() *GetMapNameNotFound {
	return &GetMapNameNotFound{}
}

/*GetMapNameNotFound handles this case with default header values.

Map not found
*/
type GetMapNameNotFound struct {
}

func (o *GetMapNameNotFound) Error() string {
	return fmt.Sprintf("[GET /map/{name}][%d] getMapNameNotFound ", 404)
}

func (o *GetMapNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
