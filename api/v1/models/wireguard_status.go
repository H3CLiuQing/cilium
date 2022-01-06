// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WireguardStatus Status of the Wireguard agent
//
// +k8s:deepcopy-gen=true
//
// swagger:model WireguardStatus
type WireguardStatus struct {

	// Wireguard interfaces managed by this Cilium instance
	Interfaces []*WireguardInterface `json:"interfaces"`
}

// Validate validates this wireguard status
func (m *WireguardStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WireguardStatus) validateInterfaces(formats strfmt.Registry) error {

	if swag.IsZero(m.Interfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Interfaces); i++ {
		if swag.IsZero(m.Interfaces[i]) { // not required
			continue
		}

		if m.Interfaces[i] != nil {
			if err := m.Interfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WireguardStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WireguardStatus) UnmarshalBinary(b []byte) error {
	var res WireguardStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
