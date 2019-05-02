// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SdxClusterRequest sdx cluster request
// swagger:model SdxClusterRequest
type SdxClusterRequest struct {

	// access cidr
	AccessCidr string `json:"accessCidr,omitempty"`

	// cluster shape
	ClusterShape string `json:"clusterShape,omitempty"`

	// tags
	Tags map[string]string `json:"tags,omitempty"`
}

// Validate validates this sdx cluster request
func (m *SdxClusterRequest) Validate(formats strfmt.Registry) error {
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
