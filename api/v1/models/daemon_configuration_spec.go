// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2022 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DaemonConfigurationSpec The controllable configuration of the daemon.
//
// swagger:model DaemonConfigurationSpec
type DaemonConfigurationSpec struct {

	// Changeable configuration
	Options ConfigurationMap `json:"options,omitempty"`

	// The policy-enforcement mode
	// Enum: [default always never]
	PolicyEnforcement string `json:"policy-enforcement,omitempty"`
}

// Validate validates this daemon configuration spec
func (m *DaemonConfigurationSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyEnforcement(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DaemonConfigurationSpec) validateOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Options) { // not required
		return nil
	}

	if err := m.Options.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("options")
		}
		return err
	}

	return nil
}

var daemonConfigurationSpecTypePolicyEnforcementPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","always","never"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		daemonConfigurationSpecTypePolicyEnforcementPropEnum = append(daemonConfigurationSpecTypePolicyEnforcementPropEnum, v)
	}
}

const (

	// DaemonConfigurationSpecPolicyEnforcementDefault captures enum value "default"
	DaemonConfigurationSpecPolicyEnforcementDefault string = "default"

	// DaemonConfigurationSpecPolicyEnforcementAlways captures enum value "always"
	DaemonConfigurationSpecPolicyEnforcementAlways string = "always"

	// DaemonConfigurationSpecPolicyEnforcementNever captures enum value "never"
	DaemonConfigurationSpecPolicyEnforcementNever string = "never"
)

// prop value enum
func (m *DaemonConfigurationSpec) validatePolicyEnforcementEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, daemonConfigurationSpecTypePolicyEnforcementPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DaemonConfigurationSpec) validatePolicyEnforcement(formats strfmt.Registry) error {

	if swag.IsZero(m.PolicyEnforcement) { // not required
		return nil
	}

	// value enum
	if err := m.validatePolicyEnforcementEnum("policy-enforcement", "body", m.PolicyEnforcement); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DaemonConfigurationSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DaemonConfigurationSpec) UnmarshalBinary(b []byte) error {
	var res DaemonConfigurationSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
