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

// PatchEndpointIDLabelsOKCode is the HTTP code returned for type PatchEndpointIDLabelsOK
const PatchEndpointIDLabelsOKCode int = 200

/*PatchEndpointIDLabelsOK Success

swagger:response patchEndpointIdLabelsOK
*/
type PatchEndpointIDLabelsOK struct {
}

// NewPatchEndpointIDLabelsOK creates PatchEndpointIDLabelsOK with default headers values
func NewPatchEndpointIDLabelsOK() *PatchEndpointIDLabelsOK {

	return &PatchEndpointIDLabelsOK{}
}

// WriteResponse to the client
func (o *PatchEndpointIDLabelsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PatchEndpointIDLabelsNotFoundCode is the HTTP code returned for type PatchEndpointIDLabelsNotFound
const PatchEndpointIDLabelsNotFoundCode int = 404

/*PatchEndpointIDLabelsNotFound Endpoint not found

swagger:response patchEndpointIdLabelsNotFound
*/
type PatchEndpointIDLabelsNotFound struct {
}

// NewPatchEndpointIDLabelsNotFound creates PatchEndpointIDLabelsNotFound with default headers values
func NewPatchEndpointIDLabelsNotFound() *PatchEndpointIDLabelsNotFound {

	return &PatchEndpointIDLabelsNotFound{}
}

// WriteResponse to the client
func (o *PatchEndpointIDLabelsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// PatchEndpointIDLabelsTooManyRequestsCode is the HTTP code returned for type PatchEndpointIDLabelsTooManyRequests
const PatchEndpointIDLabelsTooManyRequestsCode int = 429

/*PatchEndpointIDLabelsTooManyRequests Rate-limiting too many requests in the given time frame

swagger:response patchEndpointIdLabelsTooManyRequests
*/
type PatchEndpointIDLabelsTooManyRequests struct {
}

// NewPatchEndpointIDLabelsTooManyRequests creates PatchEndpointIDLabelsTooManyRequests with default headers values
func NewPatchEndpointIDLabelsTooManyRequests() *PatchEndpointIDLabelsTooManyRequests {

	return &PatchEndpointIDLabelsTooManyRequests{}
}

// WriteResponse to the client
func (o *PatchEndpointIDLabelsTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(429)
}

// PatchEndpointIDLabelsUpdateFailedCode is the HTTP code returned for type PatchEndpointIDLabelsUpdateFailed
const PatchEndpointIDLabelsUpdateFailedCode int = 500

/*PatchEndpointIDLabelsUpdateFailed Error while updating labels

swagger:response patchEndpointIdLabelsUpdateFailed
*/
type PatchEndpointIDLabelsUpdateFailed struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPatchEndpointIDLabelsUpdateFailed creates PatchEndpointIDLabelsUpdateFailed with default headers values
func NewPatchEndpointIDLabelsUpdateFailed() *PatchEndpointIDLabelsUpdateFailed {

	return &PatchEndpointIDLabelsUpdateFailed{}
}

// WithPayload adds the payload to the patch endpoint Id labels update failed response
func (o *PatchEndpointIDLabelsUpdateFailed) WithPayload(payload models.Error) *PatchEndpointIDLabelsUpdateFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch endpoint Id labels update failed response
func (o *PatchEndpointIDLabelsUpdateFailed) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchEndpointIDLabelsUpdateFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
