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

// SyncOperationV1Status sync operation v1 status
// swagger:model SyncOperationV1Status
type SyncOperationV1Status struct {

	// User synchronization operation end time
	EndTime int64 `json:"endTime,omitempty"`

	// error information about operation failure
	Error string `json:"error,omitempty"`

	// details about environments where operation failed
	Failure []*FailureDetailsV1 `json:"failure"`

	// User synchronization operation id
	// Required: true
	OperationID *string `json:"operationId"`

	// User synchronization operation start time
	StartTime int64 `json:"startTime,omitempty"`

	// User synchronization operation status
	// Enum: [REQUESTED RUNNING COMPLETED FAILED REJECTED TIMEDOUT]
	Status string `json:"status,omitempty"`

	// details about environments where operation succeeded
	Success []*SuccessDetailsV1 `json:"success"`

	// Operation type
	// Required: true
	// Enum: [USER_SYNC SET_PASSWORD]
	SyncOperationType *string `json:"syncOperationType"`
}

// Validate validates this sync operation v1 status
func (m *SyncOperationV1Status) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncOperationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SyncOperationV1Status) validateFailure(formats strfmt.Registry) error {

	if swag.IsZero(m.Failure) { // not required
		return nil
	}

	for i := 0; i < len(m.Failure); i++ {
		if swag.IsZero(m.Failure[i]) { // not required
			continue
		}

		if m.Failure[i] != nil {
			if err := m.Failure[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failure" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SyncOperationV1Status) validateOperationID(formats strfmt.Registry) error {

	if err := validate.Required("operationId", "body", m.OperationID); err != nil {
		return err
	}

	return nil
}

var syncOperationV1StatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","RUNNING","COMPLETED","FAILED","REJECTED","TIMEDOUT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		syncOperationV1StatusTypeStatusPropEnum = append(syncOperationV1StatusTypeStatusPropEnum, v)
	}
}

const (

	// SyncOperationV1StatusStatusREQUESTED captures enum value "REQUESTED"
	SyncOperationV1StatusStatusREQUESTED string = "REQUESTED"

	// SyncOperationV1StatusStatusRUNNING captures enum value "RUNNING"
	SyncOperationV1StatusStatusRUNNING string = "RUNNING"

	// SyncOperationV1StatusStatusCOMPLETED captures enum value "COMPLETED"
	SyncOperationV1StatusStatusCOMPLETED string = "COMPLETED"

	// SyncOperationV1StatusStatusFAILED captures enum value "FAILED"
	SyncOperationV1StatusStatusFAILED string = "FAILED"

	// SyncOperationV1StatusStatusREJECTED captures enum value "REJECTED"
	SyncOperationV1StatusStatusREJECTED string = "REJECTED"

	// SyncOperationV1StatusStatusTIMEDOUT captures enum value "TIMEDOUT"
	SyncOperationV1StatusStatusTIMEDOUT string = "TIMEDOUT"
)

// prop value enum
func (m *SyncOperationV1Status) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, syncOperationV1StatusTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SyncOperationV1Status) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *SyncOperationV1Status) validateSuccess(formats strfmt.Registry) error {

	if swag.IsZero(m.Success) { // not required
		return nil
	}

	for i := 0; i < len(m.Success); i++ {
		if swag.IsZero(m.Success[i]) { // not required
			continue
		}

		if m.Success[i] != nil {
			if err := m.Success[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("success" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var syncOperationV1StatusTypeSyncOperationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["USER_SYNC","SET_PASSWORD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		syncOperationV1StatusTypeSyncOperationTypePropEnum = append(syncOperationV1StatusTypeSyncOperationTypePropEnum, v)
	}
}

const (

	// SyncOperationV1StatusSyncOperationTypeUSERSYNC captures enum value "USER_SYNC"
	SyncOperationV1StatusSyncOperationTypeUSERSYNC string = "USER_SYNC"

	// SyncOperationV1StatusSyncOperationTypeSETPASSWORD captures enum value "SET_PASSWORD"
	SyncOperationV1StatusSyncOperationTypeSETPASSWORD string = "SET_PASSWORD"
)

// prop value enum
func (m *SyncOperationV1Status) validateSyncOperationTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, syncOperationV1StatusTypeSyncOperationTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SyncOperationV1Status) validateSyncOperationType(formats strfmt.Registry) error {

	if err := validate.Required("syncOperationType", "body", m.SyncOperationType); err != nil {
		return err
	}

	// value enum
	if err := m.validateSyncOperationTypeEnum("syncOperationType", "body", *m.SyncOperationType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SyncOperationV1Status) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncOperationV1Status) UnmarshalBinary(b []byte) error {
	var res SyncOperationV1Status
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
