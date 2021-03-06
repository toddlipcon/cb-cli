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

// ProjectKeystoneV3Parameters project keystone v3 parameters
// swagger:model ProjectKeystoneV3Parameters
type ProjectKeystoneV3Parameters struct {
	KeystoneV3Base

	// project domain name
	// Required: true
	ProjectDomainName *string `json:"projectDomainName"`

	// project name
	// Required: true
	ProjectName *string `json:"projectName"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ProjectKeystoneV3Parameters) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 KeystoneV3Base
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.KeystoneV3Base = aO0

	// AO1
	var dataAO1 struct {
		ProjectDomainName *string `json:"projectDomainName"`

		ProjectName *string `json:"projectName"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ProjectDomainName = dataAO1.ProjectDomainName

	m.ProjectName = dataAO1.ProjectName

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProjectKeystoneV3Parameters) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.KeystoneV3Base)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ProjectDomainName *string `json:"projectDomainName"`

		ProjectName *string `json:"projectName"`
	}

	dataAO1.ProjectDomainName = m.ProjectDomainName

	dataAO1.ProjectName = m.ProjectName

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this project keystone v3 parameters
func (m *ProjectKeystoneV3Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with KeystoneV3Base
	if err := m.KeystoneV3Base.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectDomainName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectKeystoneV3Parameters) validateProjectDomainName(formats strfmt.Registry) error {

	if err := validate.Required("projectDomainName", "body", m.ProjectDomainName); err != nil {
		return err
	}

	return nil
}

func (m *ProjectKeystoneV3Parameters) validateProjectName(formats strfmt.Registry) error {

	if err := validate.Required("projectName", "body", m.ProjectName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectKeystoneV3Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectKeystoneV3Parameters) UnmarshalBinary(b []byte) error {
	var res ProjectKeystoneV3Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
