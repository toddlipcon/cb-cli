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

// KeystoneV2Parameters keystone v2 parameters
// swagger:model KeystoneV2Parameters
type KeystoneV2Parameters struct {

	// tenant name
	// Required: true
	TenantName *string `json:"tenantName"`
}

// Validate validates this keystone v2 parameters
func (m *KeystoneV2Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTenantName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeystoneV2Parameters) validateTenantName(formats strfmt.Registry) error {

	if err := validate.Required("tenantName", "body", m.TenantName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeystoneV2Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeystoneV2Parameters) UnmarshalBinary(b []byte) error {
	var res KeystoneV2Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
