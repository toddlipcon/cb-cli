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

// SimpleEnvironmentV1Response simple environment v1 response
// swagger:model SimpleEnvironmentV1Response
type SimpleEnvironmentV1Response struct {

	// SSH key for accessing cluster node instances.
	Authentication *EnvironmentAuthenticationV1Response `json:"authentication,omitempty"`

	// Cloud platform of the environment.
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// Create freeipa in environment
	CreateFreeIpa bool `json:"createFreeIpa,omitempty"`

	// created
	Created int64 `json:"created,omitempty"`

	// crn of the creator
	Creator string `json:"creator,omitempty"`

	// Credential of the environment.
	Credential *CredentialV1Response `json:"credential,omitempty"`

	// id of the resource
	Crn string `json:"crn,omitempty"`

	// description of the resource
	Description string `json:"description,omitempty"`

	// Status of the environment.
	// Enum: [CREATION_INITIATED DELETE_INITIATED UPDATE_INITIATED NETWORK_CREATION_IN_PROGRESS NETWORK_DELETE_IN_PROGRESS RDBMS_DELETE_IN_PROGRESS FREEIPA_CREATION_IN_PROGRESS FREEIPA_DELETE_IN_PROGRESS AVAILABLE ARCHIVED CREATE_FAILED DELETE_FAILED UPDATE_FAILED]
	EnvironmentStatus string `json:"environmentStatus,omitempty"`

	// Location of the environment.
	Location *LocationV1Response `json:"location,omitempty"`

	// log cloud storage
	LogCloudStorage *CloudStorageV1Response `json:"logCloudStorage,omitempty"`

	// name of the resource
	Name string `json:"name,omitempty"`

	// Network related specifics of the environment.
	Network *EnvironmentNetworkV1Response `json:"network,omitempty"`

	// Regions of the environment.
	Regions *CompactRegionV1Response `json:"regions,omitempty"`

	// Security control for FreeIPA and Datalake deployment.
	SecurityAccess *SecurityAccessV1Response `json:"securityAccess,omitempty"`

	// status reason
	StatusReason string `json:"statusReason,omitempty"`

	// Telemetry related specifics of the environment.
	Telemetry *TelemetryV1Response `json:"telemetry,omitempty"`

	// Configuration that the connection going directly or with ccm.
	// Enum: [DIRECT CCM]
	Tunnel string `json:"tunnel,omitempty"`
}

// Validate validates this simple environment v1 response
func (m *SimpleEnvironmentV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogCloudStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTelemetry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTunnel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimpleEnvironmentV1Response) validateAuthentication(formats strfmt.Registry) error {

	if swag.IsZero(m.Authentication) { // not required
		return nil
	}

	if m.Authentication != nil {
		if err := m.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *SimpleEnvironmentV1Response) validateCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.Credential) { // not required
		return nil
	}

	if m.Credential != nil {
		if err := m.Credential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

var simpleEnvironmentV1ResponseTypeEnvironmentStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATION_INITIATED","DELETE_INITIATED","UPDATE_INITIATED","NETWORK_CREATION_IN_PROGRESS","NETWORK_DELETE_IN_PROGRESS","RDBMS_DELETE_IN_PROGRESS","FREEIPA_CREATION_IN_PROGRESS","FREEIPA_DELETE_IN_PROGRESS","AVAILABLE","ARCHIVED","CREATE_FAILED","DELETE_FAILED","UPDATE_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simpleEnvironmentV1ResponseTypeEnvironmentStatusPropEnum = append(simpleEnvironmentV1ResponseTypeEnvironmentStatusPropEnum, v)
	}
}

const (

	// SimpleEnvironmentV1ResponseEnvironmentStatusCREATIONINITIATED captures enum value "CREATION_INITIATED"
	SimpleEnvironmentV1ResponseEnvironmentStatusCREATIONINITIATED string = "CREATION_INITIATED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusDELETEINITIATED captures enum value "DELETE_INITIATED"
	SimpleEnvironmentV1ResponseEnvironmentStatusDELETEINITIATED string = "DELETE_INITIATED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusUPDATEINITIATED captures enum value "UPDATE_INITIATED"
	SimpleEnvironmentV1ResponseEnvironmentStatusUPDATEINITIATED string = "UPDATE_INITIATED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusNETWORKCREATIONINPROGRESS captures enum value "NETWORK_CREATION_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusNETWORKCREATIONINPROGRESS string = "NETWORK_CREATION_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusNETWORKDELETEINPROGRESS captures enum value "NETWORK_DELETE_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusNETWORKDELETEINPROGRESS string = "NETWORK_DELETE_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusRDBMSDELETEINPROGRESS captures enum value "RDBMS_DELETE_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusRDBMSDELETEINPROGRESS string = "RDBMS_DELETE_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusFREEIPACREATIONINPROGRESS captures enum value "FREEIPA_CREATION_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusFREEIPACREATIONINPROGRESS string = "FREEIPA_CREATION_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusFREEIPADELETEINPROGRESS captures enum value "FREEIPA_DELETE_IN_PROGRESS"
	SimpleEnvironmentV1ResponseEnvironmentStatusFREEIPADELETEINPROGRESS string = "FREEIPA_DELETE_IN_PROGRESS"

	// SimpleEnvironmentV1ResponseEnvironmentStatusAVAILABLE captures enum value "AVAILABLE"
	SimpleEnvironmentV1ResponseEnvironmentStatusAVAILABLE string = "AVAILABLE"

	// SimpleEnvironmentV1ResponseEnvironmentStatusARCHIVED captures enum value "ARCHIVED"
	SimpleEnvironmentV1ResponseEnvironmentStatusARCHIVED string = "ARCHIVED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusCREATEFAILED captures enum value "CREATE_FAILED"
	SimpleEnvironmentV1ResponseEnvironmentStatusCREATEFAILED string = "CREATE_FAILED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusDELETEFAILED captures enum value "DELETE_FAILED"
	SimpleEnvironmentV1ResponseEnvironmentStatusDELETEFAILED string = "DELETE_FAILED"

	// SimpleEnvironmentV1ResponseEnvironmentStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	SimpleEnvironmentV1ResponseEnvironmentStatusUPDATEFAILED string = "UPDATE_FAILED"
)

// prop value enum
func (m *SimpleEnvironmentV1Response) validateEnvironmentStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, simpleEnvironmentV1ResponseTypeEnvironmentStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SimpleEnvironmentV1Response) validateEnvironmentStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvironmentStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnvironmentStatusEnum("environmentStatus", "body", m.EnvironmentStatus); err != nil {
		return err
	}

	return nil
}

func (m *SimpleEnvironmentV1Response) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *SimpleEnvironmentV1Response) validateLogCloudStorage(formats strfmt.Registry) error {

	if swag.IsZero(m.LogCloudStorage) { // not required
		return nil
	}

	if m.LogCloudStorage != nil {
		if err := m.LogCloudStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logCloudStorage")
			}
			return err
		}
	}

	return nil
}

func (m *SimpleEnvironmentV1Response) validateNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *SimpleEnvironmentV1Response) validateRegions(formats strfmt.Registry) error {

	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	if m.Regions != nil {
		if err := m.Regions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("regions")
			}
			return err
		}
	}

	return nil
}

func (m *SimpleEnvironmentV1Response) validateSecurityAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityAccess) { // not required
		return nil
	}

	if m.SecurityAccess != nil {
		if err := m.SecurityAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityAccess")
			}
			return err
		}
	}

	return nil
}

func (m *SimpleEnvironmentV1Response) validateTelemetry(formats strfmt.Registry) error {

	if swag.IsZero(m.Telemetry) { // not required
		return nil
	}

	if m.Telemetry != nil {
		if err := m.Telemetry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("telemetry")
			}
			return err
		}
	}

	return nil
}

var simpleEnvironmentV1ResponseTypeTunnelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DIRECT","CCM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simpleEnvironmentV1ResponseTypeTunnelPropEnum = append(simpleEnvironmentV1ResponseTypeTunnelPropEnum, v)
	}
}

const (

	// SimpleEnvironmentV1ResponseTunnelDIRECT captures enum value "DIRECT"
	SimpleEnvironmentV1ResponseTunnelDIRECT string = "DIRECT"

	// SimpleEnvironmentV1ResponseTunnelCCM captures enum value "CCM"
	SimpleEnvironmentV1ResponseTunnelCCM string = "CCM"
)

// prop value enum
func (m *SimpleEnvironmentV1Response) validateTunnelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, simpleEnvironmentV1ResponseTypeTunnelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SimpleEnvironmentV1Response) validateTunnel(formats strfmt.Registry) error {

	if swag.IsZero(m.Tunnel) { // not required
		return nil
	}

	// value enum
	if err := m.validateTunnelEnum("tunnel", "body", m.Tunnel); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SimpleEnvironmentV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimpleEnvironmentV1Response) UnmarshalBinary(b []byte) error {
	var res SimpleEnvironmentV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
