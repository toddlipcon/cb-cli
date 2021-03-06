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

// ListFreeIpaV1Response list free ipa v1 response
// swagger:model ListFreeIpaV1Response
type ListFreeIpaV1Response struct {

	// crn
	// Required: true
	Crn *string `json:"crn"`

	// domain
	Domain string `json:"domain,omitempty"`

	// CRN of the environment
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// name of the freeipa stack
	// Required: true
	Name *string `json:"name"`

	// status
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE STACK_AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED CREATE_FAILED DELETE_IN_PROGRESS DELETE_FAILED DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED UNREACHABLE UNHEALTHY DELETED_ON_PROVIDER_SIDE UNKNOWN]
	Status string `json:"status,omitempty"`

	// status string
	StatusString string `json:"statusString,omitempty"`
}

// Validate validates this list free ipa v1 response
func (m *ListFreeIpaV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *ListFreeIpaV1Response) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *ListFreeIpaV1Response) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

func (m *ListFreeIpaV1Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var listFreeIpaV1ResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","STACK_AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED","UNREACHABLE","UNHEALTHY","DELETED_ON_PROVIDER_SIDE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listFreeIpaV1ResponseTypeStatusPropEnum = append(listFreeIpaV1ResponseTypeStatusPropEnum, v)
	}
}

const (

	// ListFreeIpaV1ResponseStatusREQUESTED captures enum value "REQUESTED"
	ListFreeIpaV1ResponseStatusREQUESTED string = "REQUESTED"

	// ListFreeIpaV1ResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	ListFreeIpaV1ResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// ListFreeIpaV1ResponseStatusAVAILABLE captures enum value "AVAILABLE"
	ListFreeIpaV1ResponseStatusAVAILABLE string = "AVAILABLE"

	// ListFreeIpaV1ResponseStatusSTACKAVAILABLE captures enum value "STACK_AVAILABLE"
	ListFreeIpaV1ResponseStatusSTACKAVAILABLE string = "STACK_AVAILABLE"

	// ListFreeIpaV1ResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	ListFreeIpaV1ResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// ListFreeIpaV1ResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	ListFreeIpaV1ResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// ListFreeIpaV1ResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	ListFreeIpaV1ResponseStatusUPDATEFAILED string = "UPDATE_FAILED"

	// ListFreeIpaV1ResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	ListFreeIpaV1ResponseStatusCREATEFAILED string = "CREATE_FAILED"

	// ListFreeIpaV1ResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	ListFreeIpaV1ResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// ListFreeIpaV1ResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	ListFreeIpaV1ResponseStatusDELETEFAILED string = "DELETE_FAILED"

	// ListFreeIpaV1ResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	ListFreeIpaV1ResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// ListFreeIpaV1ResponseStatusSTOPPED captures enum value "STOPPED"
	ListFreeIpaV1ResponseStatusSTOPPED string = "STOPPED"

	// ListFreeIpaV1ResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	ListFreeIpaV1ResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// ListFreeIpaV1ResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	ListFreeIpaV1ResponseStatusSTARTREQUESTED string = "START_REQUESTED"

	// ListFreeIpaV1ResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	ListFreeIpaV1ResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// ListFreeIpaV1ResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	ListFreeIpaV1ResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// ListFreeIpaV1ResponseStatusSTARTFAILED captures enum value "START_FAILED"
	ListFreeIpaV1ResponseStatusSTARTFAILED string = "START_FAILED"

	// ListFreeIpaV1ResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	ListFreeIpaV1ResponseStatusSTOPFAILED string = "STOP_FAILED"

	// ListFreeIpaV1ResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	ListFreeIpaV1ResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// ListFreeIpaV1ResponseStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	ListFreeIpaV1ResponseStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"

	// ListFreeIpaV1ResponseStatusUNREACHABLE captures enum value "UNREACHABLE"
	ListFreeIpaV1ResponseStatusUNREACHABLE string = "UNREACHABLE"

	// ListFreeIpaV1ResponseStatusUNHEALTHY captures enum value "UNHEALTHY"
	ListFreeIpaV1ResponseStatusUNHEALTHY string = "UNHEALTHY"

	// ListFreeIpaV1ResponseStatusDELETEDONPROVIDERSIDE captures enum value "DELETED_ON_PROVIDER_SIDE"
	ListFreeIpaV1ResponseStatusDELETEDONPROVIDERSIDE string = "DELETED_ON_PROVIDER_SIDE"

	// ListFreeIpaV1ResponseStatusUNKNOWN captures enum value "UNKNOWN"
	ListFreeIpaV1ResponseStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *ListFreeIpaV1Response) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, listFreeIpaV1ResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ListFreeIpaV1Response) validateStatus(formats strfmt.Registry) error {

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
func (m *ListFreeIpaV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListFreeIpaV1Response) UnmarshalBinary(b []byte) error {
	var res ListFreeIpaV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
