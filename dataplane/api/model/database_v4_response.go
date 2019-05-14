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

// DatabaseV4Response database v4 response
// swagger:model DatabaseV4Response
type DatabaseV4Response struct {

	// list of clusters which use config
	// Unique: true
	ClusterNames []string `json:"clusterNames"`

	// Name of the JDBC connection driver (for example: 'org.postgresql.Driver')
	// Required: true
	ConnectionDriver *string `json:"connectionDriver"`

	// Password to use for the jdbc connection
	ConnectionPassword *SecretResponse `json:"connectionPassword,omitempty"`

	// JDBC connection URL in the form of jdbc:<db-type>://<address>:<port>/<db>
	// Required: true
	ConnectionURL *string `json:"connectionURL"`

	// Username to use for the jdbc connection
	ConnectionUserName *SecretResponse `json:"connectionUserName,omitempty"`

	// URL that points to the jar of the connection driver(connector)
	ConnectorJarURL string `json:"connectorJarUrl,omitempty"`

	// creation time of the resource in long
	CreationDate int64 `json:"creationDate,omitempty"`

	// Name of the external database engine (MYSQL, POSTGRES...)
	// Required: true
	DatabaseEngine *string `json:"databaseEngine"`

	// Display name of the external database engine (Mysql, PostgreSQL...)
	// Required: true
	DatabaseEngineDisplayName *string `json:"databaseEngineDisplayName"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// Name of the RDS configuration resource
	// Required: true
	// Max Length: 100
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// Type of RDS, aka the service name that will use the RDS like HIVE, DRUID, SUPERSET, RANGER, etc.
	// Required: true
	Type *string `json:"type"`

	// workspace of the resource
	Workspace *WorkspaceResourceV4Response `json:"workspace,omitempty"`
}

// Validate validates this database v4 response
func (m *DatabaseV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionUserName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabaseEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabaseEngineDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseV4Response) validateClusterNames(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("clusterNames", "body", m.ClusterNames); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Response) validateConnectionDriver(formats strfmt.Registry) error {

	if err := validate.Required("connectionDriver", "body", m.ConnectionDriver); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Response) validateConnectionPassword(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionPassword) { // not required
		return nil
	}

	if m.ConnectionPassword != nil {
		if err := m.ConnectionPassword.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionPassword")
			}
			return err
		}
	}

	return nil
}

func (m *DatabaseV4Response) validateConnectionURL(formats strfmt.Registry) error {

	if err := validate.Required("connectionURL", "body", m.ConnectionURL); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Response) validateConnectionUserName(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionUserName) { // not required
		return nil
	}

	if m.ConnectionUserName != nil {
		if err := m.ConnectionUserName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionUserName")
			}
			return err
		}
	}

	return nil
}

func (m *DatabaseV4Response) validateDatabaseEngine(formats strfmt.Registry) error {

	if err := validate.Required("databaseEngine", "body", m.DatabaseEngine); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Response) validateDatabaseEngineDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("databaseEngineDisplayName", "body", m.DatabaseEngineDisplayName); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Response) validateDescription(formats strfmt.Registry) error {

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

func (m *DatabaseV4Response) validateName(formats strfmt.Registry) error {

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

func (m *DatabaseV4Response) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseV4Response) validateWorkspace(formats strfmt.Registry) error {

	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {
		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseV4Response) UnmarshalBinary(b []byte) error {
	var res DatabaseV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
