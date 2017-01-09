package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*HostGroupRequest host group request

swagger:model HostGroupRequest
*/
type HostGroupRequest struct {

	/* instance group or resource constraint for a hostgroup

	Required: true
	*/
	Constraint *Constraint `json:"constraint"`

	/* name of the resource

	Required: true
	*/
	Name string `json:"name"`

	/* referenced recipe ids

	Unique: true
	*/
	RecipeIds []int64 `json:"recipeIds,omitempty"`

	/* recovery mode of the hostgroup's nodes
	 */
	RecoveryMode *string `json:"recoveryMode,omitempty"`
}

// Validate validates this host group request
func (m *HostGroupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstraint(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecipeIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecoveryMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostGroupRequest) validateConstraint(formats strfmt.Registry) error {

	if m.Constraint != nil {

		if err := m.Constraint.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *HostGroupRequest) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *HostGroupRequest) validateRecipeIds(formats strfmt.Registry) error {

	if swag.IsZero(m.RecipeIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("recipeIds", "body", m.RecipeIds); err != nil {
		return err
	}

	for i := 0; i < len(m.RecipeIds); i++ {

		if err := validate.Required("recipeIds"+"."+strconv.Itoa(i), "body", int64(m.RecipeIds[i])); err != nil {
			return err
		}

	}

	return nil
}

var hostGroupRequestTypeRecoveryModePropEnum []interface{}

func (m *HostGroupRequest) validateRecoveryModeEnum(path, location string, value string) error {
	if hostGroupRequestTypeRecoveryModePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["MANUAL","AUTO"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			hostGroupRequestTypeRecoveryModePropEnum = append(hostGroupRequestTypeRecoveryModePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, hostGroupRequestTypeRecoveryModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HostGroupRequest) validateRecoveryMode(formats strfmt.Registry) error {

	if swag.IsZero(m.RecoveryMode) { // not required
		return nil
	}

	if err := m.validateRecoveryModeEnum("recoveryMode", "body", *m.RecoveryMode); err != nil {
		return err
	}

	return nil
}
