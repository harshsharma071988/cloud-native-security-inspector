// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Details details
//
// swagger:model Details
type Details struct {

	// baseline
	Baseline string `json:"baseline,omitempty"`

	// container
	// Required: true
	Container *string `json:"container"`

	// dimension
	Dimension string `json:"dimension,omitempty"`

	// failure
	// Required: true
	Failure *string `json:"failure"`

	// workload
	// Required: true
	Workload *string `json:"workload"`
}

// Validate validates this details
func (m *Details) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Details) validateContainer(formats strfmt.Registry) error {

	if err := validate.Required("container", "body", m.Container); err != nil {
		return err
	}

	return nil
}

func (m *Details) validateFailure(formats strfmt.Registry) error {

	if err := validate.Required("failure", "body", m.Failure); err != nil {
		return err
	}

	return nil
}

func (m *Details) validateWorkload(formats strfmt.Registry) error {

	if err := validate.Required("workload", "body", m.Workload); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Details) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Details) UnmarshalBinary(b []byte) error {
	var res Details
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}