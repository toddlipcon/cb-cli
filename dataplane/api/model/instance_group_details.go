// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InstanceGroupDetails instance group details
// swagger:model InstanceGroupDetails
type InstanceGroupDetails struct {

	// group name
	GroupName string `json:"groupName,omitempty"`

	// group type
	GroupType string `json:"groupType,omitempty"`

	// instance type
	InstanceType string `json:"instanceType,omitempty"`

	// node count
	NodeCount int32 `json:"nodeCount,omitempty"`

	// security group
	SecurityGroup *SecurityGroupDetails `json:"securityGroup,omitempty"`

	// volumes
	Volumes []*VolumeDetails `json:"volumes"`
}

// Validate validates this instance group details
func (m *InstanceGroupDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecurityGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceGroupDetails) validateSecurityGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityGroup) { // not required
		return nil
	}

	if m.SecurityGroup != nil {
		if err := m.SecurityGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityGroup")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceGroupDetails) validateVolumes(formats strfmt.Registry) error {

	if swag.IsZero(m.Volumes) { // not required
		return nil
	}

	for i := 0; i < len(m.Volumes); i++ {
		if swag.IsZero(m.Volumes[i]) { // not required
			continue
		}

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceGroupDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceGroupDetails) UnmarshalBinary(b []byte) error {
	var res InstanceGroupDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
