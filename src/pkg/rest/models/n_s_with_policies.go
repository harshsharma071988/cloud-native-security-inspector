// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NSWithPolicies n s with policies
//
// swagger:model NSWithPolicies
type NSWithPolicies struct {

	// name of inspection policies applied to the namespace
	Name string `json:"name,omitempty"`

	// policies
	Policies []string `json:"policies"`
}

// Validate validates this n s with policies
func (m *NSWithPolicies) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NSWithPolicies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NSWithPolicies) UnmarshalBinary(b []byte) error {
	var res NSWithPolicies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}