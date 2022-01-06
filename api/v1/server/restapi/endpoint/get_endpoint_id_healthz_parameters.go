// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetEndpointIDHealthzParams creates a new GetEndpointIDHealthzParams object
// no default values defined in spec.
func NewGetEndpointIDHealthzParams() GetEndpointIDHealthzParams {

	return GetEndpointIDHealthzParams{}
}

// GetEndpointIDHealthzParams contains all the bound params for the get endpoint ID healthz operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetEndpointIDHealthz
type GetEndpointIDHealthzParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*String describing an endpoint with the format ``[prefix:]id``. If no prefix
	is specified, a prefix of ``cilium-local:`` is assumed. Not all endpoints
	will be addressable by all endpoint ID prefixes with the exception of the
	local Cilium UUID which is assigned to all endpoints.

	Supported endpoint id prefixes:
	  - cilium-local: Local Cilium endpoint UUID, e.g. cilium-local:3389595
	  - cilium-global: Global Cilium endpoint UUID, e.g. cilium-global:cluster1:nodeX:452343
	  - container-id: Container runtime ID, e.g. container-id:22222
	  - container-name: Container name, e.g. container-name:foobar
	  - pod-name: pod name for this container if K8s is enabled, e.g. pod-name:default:foobar
	  - docker-endpoint: Docker libnetwork endpoint ID, e.g. docker-endpoint:4444

	  Required: true
	  In: path
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetEndpointIDHealthzParams() beforehand.
func (o *GetEndpointIDHealthzParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GetEndpointIDHealthzParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}
