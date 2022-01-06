// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetEndpointOKCode is the HTTP code returned for type GetEndpointOK
const GetEndpointOKCode int = 200

/*GetEndpointOK Success

swagger:response getEndpointOK
*/
type GetEndpointOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Endpoint `json:"body,omitempty"`
}

// NewGetEndpointOK creates GetEndpointOK with default headers values
func NewGetEndpointOK() *GetEndpointOK {

	return &GetEndpointOK{}
}

// WithPayload adds the payload to the get endpoint o k response
func (o *GetEndpointOK) WithPayload(payload []*models.Endpoint) *GetEndpointOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get endpoint o k response
func (o *GetEndpointOK) SetPayload(payload []*models.Endpoint) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEndpointOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Endpoint, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetEndpointNotFoundCode is the HTTP code returned for type GetEndpointNotFound
const GetEndpointNotFoundCode int = 404

/*GetEndpointNotFound Endpoints with provided parameters not found

swagger:response getEndpointNotFound
*/
type GetEndpointNotFound struct {
}

// NewGetEndpointNotFound creates GetEndpointNotFound with default headers values
func NewGetEndpointNotFound() *GetEndpointNotFound {

	return &GetEndpointNotFound{}
}

// WriteResponse to the client
func (o *GetEndpointNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetEndpointTooManyRequestsCode is the HTTP code returned for type GetEndpointTooManyRequests
const GetEndpointTooManyRequestsCode int = 429

/*GetEndpointTooManyRequests Rate-limiting too many requests in the given time frame

swagger:response getEndpointTooManyRequests
*/
type GetEndpointTooManyRequests struct {
}

// NewGetEndpointTooManyRequests creates GetEndpointTooManyRequests with default headers values
func NewGetEndpointTooManyRequests() *GetEndpointTooManyRequests {

	return &GetEndpointTooManyRequests{}
}

// WriteResponse to the client
func (o *GetEndpointTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(429)
}
