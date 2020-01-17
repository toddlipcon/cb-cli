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

// StackMatrixV4Response stack matrix v4 response
// swagger:model StackMatrixV4Response
type StackMatrixV4Response struct {

	// cdh
	Cdh map[string]ClouderaManagerStackDescriptorV4Response `json:"cdh,omitempty"`
}

// Validate validates this stack matrix v4 response
func (m *StackMatrixV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCdh(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackMatrixV4Response) validateCdh(formats strfmt.Registry) error {

	if swag.IsZero(m.Cdh) { // not required
		return nil
	}

	for k := range m.Cdh {

		if err := validate.Required("cdh"+"."+k, "body", m.Cdh[k]); err != nil {
			return err
		}
		if val, ok := m.Cdh[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackMatrixV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackMatrixV4Response) UnmarshalBinary(b []byte) error {
	var res StackMatrixV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
