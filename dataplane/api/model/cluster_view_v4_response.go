// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterViewV4Response cluster view v4 response
// swagger:model ClusterViewV4Response
type ClusterViewV4Response struct {

	// cluster definition for the cluster
	ClusterDefinition *ClusterDefinitionV4ViewResponse `json:"clusterDefinition,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// collection of hostgroups
	// Unique: true
	HostGroups []*HostGroupViewV4Response `json:"hostGroups"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// Kerberos config name for the cluster
	KerberosName string `json:"kerberosName,omitempty"`

	// name of the resource
	// Required: true
	// Max Length: 100
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// tells wether the cluster is secured or not
	Secure *bool `json:"secure,omitempty"`

	// public ambari ip of the stack
	ServerIP string `json:"serverIp,omitempty"`

	// shared service for a specific stack
	SharedServiceResponse *SharedServiceV4Response `json:"sharedServiceResponse,omitempty"`

	// status of the cluster
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED CREATE_FAILED ENABLE_SECURITY_FAILED PRE_DELETE_IN_PROGRESS DELETE_IN_PROGRESS DELETE_FAILED DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED]
	Status string `json:"status,omitempty"`
}

// Validate validates this cluster view v4 response
func (m *ClusterViewV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedServiceResponse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterViewV4Response) validateClusterDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterDefinition) { // not required
		return nil
	}

	if m.ClusterDefinition != nil {
		if err := m.ClusterDefinition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterDefinition")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterViewV4Response) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *ClusterViewV4Response) validateHostGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.HostGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("hostGroups", "body", m.HostGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.HostGroups); i++ {
		if swag.IsZero(m.HostGroups[i]) { // not required
			continue
		}

		if m.HostGroups[i] != nil {
			if err := m.HostGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterViewV4Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterViewV4Response) validateSharedServiceResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.SharedServiceResponse) { // not required
		return nil
	}

	if m.SharedServiceResponse != nil {
		if err := m.SharedServiceResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharedServiceResponse")
			}
			return err
		}
	}

	return nil
}

var clusterViewV4ResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterViewV4ResponseTypeStatusPropEnum = append(clusterViewV4ResponseTypeStatusPropEnum, v)
	}
}

const (

	// ClusterViewV4ResponseStatusREQUESTED captures enum value "REQUESTED"
	ClusterViewV4ResponseStatusREQUESTED string = "REQUESTED"

	// ClusterViewV4ResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	ClusterViewV4ResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// ClusterViewV4ResponseStatusAVAILABLE captures enum value "AVAILABLE"
	ClusterViewV4ResponseStatusAVAILABLE string = "AVAILABLE"

	// ClusterViewV4ResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	ClusterViewV4ResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// ClusterViewV4ResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	ClusterViewV4ResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// ClusterViewV4ResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	ClusterViewV4ResponseStatusUPDATEFAILED string = "UPDATE_FAILED"

	// ClusterViewV4ResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	ClusterViewV4ResponseStatusCREATEFAILED string = "CREATE_FAILED"

	// ClusterViewV4ResponseStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	ClusterViewV4ResponseStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"

	// ClusterViewV4ResponseStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	ClusterViewV4ResponseStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"

	// ClusterViewV4ResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	ClusterViewV4ResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// ClusterViewV4ResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	ClusterViewV4ResponseStatusDELETEFAILED string = "DELETE_FAILED"

	// ClusterViewV4ResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	ClusterViewV4ResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// ClusterViewV4ResponseStatusSTOPPED captures enum value "STOPPED"
	ClusterViewV4ResponseStatusSTOPPED string = "STOPPED"

	// ClusterViewV4ResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	ClusterViewV4ResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// ClusterViewV4ResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	ClusterViewV4ResponseStatusSTARTREQUESTED string = "START_REQUESTED"

	// ClusterViewV4ResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	ClusterViewV4ResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// ClusterViewV4ResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	ClusterViewV4ResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// ClusterViewV4ResponseStatusSTARTFAILED captures enum value "START_FAILED"
	ClusterViewV4ResponseStatusSTARTFAILED string = "START_FAILED"

	// ClusterViewV4ResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	ClusterViewV4ResponseStatusSTOPFAILED string = "STOP_FAILED"

	// ClusterViewV4ResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	ClusterViewV4ResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// ClusterViewV4ResponseStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	ClusterViewV4ResponseStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"
)

// prop value enum
func (m *ClusterViewV4Response) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterViewV4ResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterViewV4Response) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterViewV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterViewV4Response) UnmarshalBinary(b []byte) error {
	var res ClusterViewV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
