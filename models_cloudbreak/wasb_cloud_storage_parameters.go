// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WasbCloudStorageParameters wasb cloud storage parameters
// swagger:model WasbCloudStorageParameters

type WasbCloudStorageParameters struct {

	// account key
	// Required: true
	AccountKey *string `json:"accountKey"`

	// account name
	// Required: true
	AccountName *string `json:"accountName"`

	// secure
	Secure *bool `json:"secure,omitempty"`
}

/* polymorph WasbCloudStorageParameters accountKey false */

/* polymorph WasbCloudStorageParameters accountName false */

/* polymorph WasbCloudStorageParameters secure false */

// Validate validates this wasb cloud storage parameters
func (m *WasbCloudStorageParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountKey(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAccountName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WasbCloudStorageParameters) validateAccountKey(formats strfmt.Registry) error {

	if err := validate.Required("accountKey", "body", m.AccountKey); err != nil {
		return err
	}

	return nil
}

func (m *WasbCloudStorageParameters) validateAccountName(formats strfmt.Registry) error {

	if err := validate.Required("accountName", "body", m.AccountName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WasbCloudStorageParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WasbCloudStorageParameters) UnmarshalBinary(b []byte) error {
	var res WasbCloudStorageParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}