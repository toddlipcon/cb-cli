// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SdxClusterRequest sdx cluster request
// swagger:model SdxClusterRequest
type SdxClusterRequest struct {

	// cloud storage
	CloudStorage *SdxCloudStorageRequest `json:"cloudStorage,omitempty"`

	// cluster shape
	// Required: true
	// Enum: [CUSTOM LIGHT_DUTY MEDIUM_DUTY_HA]
	ClusterShape *string `json:"clusterShape"`

	// environment
	// Required: true
	Environment *string `json:"environment"`

	// external database
	ExternalDatabase *SdxDatabaseRequest `json:"externalDatabase,omitempty"`

	// runtime
	Runtime string `json:"runtime,omitempty"`

	// tags
	Tags map[string]string `json:"tags,omitempty"`
}

// Validate validates this sdx cluster request
func (m *SdxClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterShape(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalDatabase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SdxClusterRequest) validateCloudStorage(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudStorage) { // not required
		return nil
	}

	if m.CloudStorage != nil {
		if err := m.CloudStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudStorage")
			}
			return err
		}
	}

	return nil
}

var sdxClusterRequestTypeClusterShapePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CUSTOM","LIGHT_DUTY","MEDIUM_DUTY_HA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sdxClusterRequestTypeClusterShapePropEnum = append(sdxClusterRequestTypeClusterShapePropEnum, v)
	}
}

const (

	// SdxClusterRequestClusterShapeCUSTOM captures enum value "CUSTOM"
	SdxClusterRequestClusterShapeCUSTOM string = "CUSTOM"

	// SdxClusterRequestClusterShapeLIGHTDUTY captures enum value "LIGHT_DUTY"
	SdxClusterRequestClusterShapeLIGHTDUTY string = "LIGHT_DUTY"

	// SdxClusterRequestClusterShapeMEDIUMDUTYHA captures enum value "MEDIUM_DUTY_HA"
	SdxClusterRequestClusterShapeMEDIUMDUTYHA string = "MEDIUM_DUTY_HA"
)

// prop value enum
func (m *SdxClusterRequest) validateClusterShapeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sdxClusterRequestTypeClusterShapePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SdxClusterRequest) validateClusterShape(formats strfmt.Registry) error {

	if err := validate.Required("clusterShape", "body", m.ClusterShape); err != nil {
		return err
	}

	// value enum
	if err := m.validateClusterShapeEnum("clusterShape", "body", *m.ClusterShape); err != nil {
		return err
	}

	return nil
}

func (m *SdxClusterRequest) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

func (m *SdxClusterRequest) validateExternalDatabase(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalDatabase) { // not required
		return nil
	}

	if m.ExternalDatabase != nil {
		if err := m.ExternalDatabase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalDatabase")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SdxClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SdxClusterRequest) UnmarshalBinary(b []byte) error {
	var res SdxClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
