// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FeaturesRequest features request
// swagger:model FeaturesRequest
type FeaturesRequest struct {

	// enable cluster deployment log reporting.
	ReportDeploymentLogs *FeatureSetting `json:"reportDeploymentLogs,omitempty"`

	// enable shared Altus credential usage
	UseSharedAltusCredential *FeatureSetting `json:"useSharedAltusCredential,omitempty"`

	// Workload analytics (telemetry) settings.
	WorkloadAnalytics *FeatureSetting `json:"workloadAnalytics,omitempty"`
}

// Validate validates this features request
func (m *FeaturesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReportDeploymentLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseSharedAltusCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadAnalytics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeaturesRequest) validateReportDeploymentLogs(formats strfmt.Registry) error {

	if swag.IsZero(m.ReportDeploymentLogs) { // not required
		return nil
	}

	if m.ReportDeploymentLogs != nil {
		if err := m.ReportDeploymentLogs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reportDeploymentLogs")
			}
			return err
		}
	}

	return nil
}

func (m *FeaturesRequest) validateUseSharedAltusCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.UseSharedAltusCredential) { // not required
		return nil
	}

	if m.UseSharedAltusCredential != nil {
		if err := m.UseSharedAltusCredential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("useSharedAltusCredential")
			}
			return err
		}
	}

	return nil
}

func (m *FeaturesRequest) validateWorkloadAnalytics(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkloadAnalytics) { // not required
		return nil
	}

	if m.WorkloadAnalytics != nil {
		if err := m.WorkloadAnalytics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workloadAnalytics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeaturesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeaturesRequest) UnmarshalBinary(b []byte) error {
	var res FeaturesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
