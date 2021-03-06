// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DistroXScaleV1Request distro x scale v1 request
// swagger:model DistroXScaleV1Request
type DistroXScaleV1Request struct {

	// scaling adjustment of the instance groups
	// Required: true
	DesiredCount *int32 `json:"desiredCount"`

	// name of the instance group
	// Required: true
	Group *string `json:"group"`
}

// Validate validates this distro x scale v1 request
func (m *DistroXScaleV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDesiredCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistroXScaleV1Request) validateDesiredCount(formats strfmt.Registry) error {

	if err := validate.Required("desiredCount", "body", m.DesiredCount); err != nil {
		return err
	}

	return nil
}

func (m *DistroXScaleV1Request) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DistroXScaleV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistroXScaleV1Request) UnmarshalBinary(b []byte) error {
	var res DistroXScaleV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
