// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetConfigOKCode is the HTTP code returned for type GetConfigOK
const GetConfigOKCode int = 200

/*GetConfigOK Success

swagger:response getConfigOK
*/
type GetConfigOK struct {

	/*
	  In: Body
	*/
	Payload *models.DaemonConfiguration `json:"body,omitempty"`
}

// NewGetConfigOK creates GetConfigOK with default headers values
func NewGetConfigOK() *GetConfigOK {

	return &GetConfigOK{}
}

// WithPayload adds the payload to the get config o k response
func (o *GetConfigOK) WithPayload(payload *models.DaemonConfiguration) *GetConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config o k response
func (o *GetConfigOK) SetPayload(payload *models.DaemonConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
