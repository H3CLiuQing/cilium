// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetLrpHandlerFunc turns a function with the right signature into a get lrp handler
type GetLrpHandlerFunc func(GetLrpParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLrpHandlerFunc) Handle(params GetLrpParams) middleware.Responder {
	return fn(params)
}

// GetLrpHandler interface for that can handle valid get lrp params
type GetLrpHandler interface {
	Handle(GetLrpParams) middleware.Responder
}

// NewGetLrp creates a new http.Handler for the get lrp operation
func NewGetLrp(ctx *middleware.Context, handler GetLrpHandler) *GetLrp {
	return &GetLrp{Context: ctx, Handler: handler}
}

/*GetLrp swagger:route GET /lrp service getLrp

Retrieve list of all local redirect policies

*/
type GetLrp struct {
	Context *middleware.Context
	Handler GetLrpHandler
}

func (o *GetLrp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetLrpParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
