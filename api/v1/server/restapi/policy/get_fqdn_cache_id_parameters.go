// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetFqdnCacheIDParams creates a new GetFqdnCacheIDParams object
// no default values defined in spec.
func NewGetFqdnCacheIDParams() GetFqdnCacheIDParams {

	return GetFqdnCacheIDParams{}
}

// GetFqdnCacheIDParams contains all the bound params for the get fqdn cache ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetFqdnCacheID
type GetFqdnCacheIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A CIDR range of IPs
	  In: query
	*/
	Cidr *string
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
	/*A toFQDNs compatible matchPattern expression
	  In: query
	*/
	Matchpattern *string
	/*Source from which FQDN entries come from
	  In: query
	*/
	Source *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetFqdnCacheIDParams() beforehand.
func (o *GetFqdnCacheIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCidr, qhkCidr, _ := qs.GetOK("cidr")
	if err := o.bindCidr(qCidr, qhkCidr, route.Formats); err != nil {
		res = append(res, err)
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qMatchpattern, qhkMatchpattern, _ := qs.GetOK("matchpattern")
	if err := o.bindMatchpattern(qMatchpattern, qhkMatchpattern, route.Formats); err != nil {
		res = append(res, err)
	}

	qSource, qhkSource, _ := qs.GetOK("source")
	if err := o.bindSource(qSource, qhkSource, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCidr binds and validates parameter Cidr from query.
func (o *GetFqdnCacheIDParams) bindCidr(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Cidr = &raw

	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GetFqdnCacheIDParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}

// bindMatchpattern binds and validates parameter Matchpattern from query.
func (o *GetFqdnCacheIDParams) bindMatchpattern(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Matchpattern = &raw

	return nil
}

// bindSource binds and validates parameter Source from query.
func (o *GetFqdnCacheIDParams) bindSource(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Source = &raw

	return nil
}
