// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpgradeOptionV4Response upgrade option v4 response
// swagger:model UpgradeOptionV4Response
type UpgradeOptionV4Response struct {

	// current
	Current *ImageInfoV4Response `json:"current,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// upgrade
	Upgrade *ImageInfoV4Response `json:"upgrade,omitempty"`
}

// Validate validates this upgrade option v4 response
func (m *UpgradeOptionV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgrade(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpgradeOptionV4Response) validateCurrent(formats strfmt.Registry) error {

	if swag.IsZero(m.Current) { // not required
		return nil
	}

	if m.Current != nil {
		if err := m.Current.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

func (m *UpgradeOptionV4Response) validateUpgrade(formats strfmt.Registry) error {

	if swag.IsZero(m.Upgrade) { // not required
		return nil
	}

	if m.Upgrade != nil {
		if err := m.Upgrade.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upgrade")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpgradeOptionV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeOptionV4Response) UnmarshalBinary(b []byte) error {
	var res UpgradeOptionV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
