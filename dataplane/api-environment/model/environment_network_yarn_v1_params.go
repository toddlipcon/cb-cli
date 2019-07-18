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

// EnvironmentNetworkYarnV1Params environment network yarn v1 params
// swagger:model EnvironmentNetworkYarnV1Params
type EnvironmentNetworkYarnV1Params struct {

	// Subnet ids of the specified networks
	// Required: true
	Queue *string `json:"queue"`
}

// Validate validates this environment network yarn v1 params
func (m *EnvironmentNetworkYarnV1Params) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvironmentNetworkYarnV1Params) validateQueue(formats strfmt.Registry) error {

	if err := validate.Required("queue", "body", m.Queue); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentNetworkYarnV1Params) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentNetworkYarnV1Params) UnmarshalBinary(b []byte) error {
	var res EnvironmentNetworkYarnV1Params
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
