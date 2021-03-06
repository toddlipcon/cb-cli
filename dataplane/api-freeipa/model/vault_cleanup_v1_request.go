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

// VaultCleanupV1Request vault cleanup v1 request
// swagger:model VaultCleanupV1Request
type VaultCleanupV1Request struct {

	// CRN of the cluster
	ClusterCrn string `json:"clusterCrn,omitempty"`

	// CRN of the environment
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`
}

// Validate validates this vault cleanup v1 request
func (m *VaultCleanupV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VaultCleanupV1Request) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VaultCleanupV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VaultCleanupV1Request) UnmarshalBinary(b []byte) error {
	var res VaultCleanupV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
