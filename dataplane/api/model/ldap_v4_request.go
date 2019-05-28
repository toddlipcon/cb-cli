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

// LdapV4Request ldap v4 request
// swagger:model LdapV4Request
type LdapV4Request struct {

	// LDAP group for administrators
	AdminGroup string `json:"adminGroup,omitempty"`

	// bind distinguished name for connection test and group search (e.g. cn=admin,dc=example,dc=org)
	// Required: true
	BindDn *string `json:"bindDn"`

	// password for the provided bind DN
	// Required: true
	BindPassword *string `json:"bindPassword"`

	// Self-signed certificate of LDAPS server
	Certificate string `json:"certificate,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// directory type of server LDAP or ACTIVE_DIRECTORY and the default is ACTIVE_DIRECTORY
	// Enum: [LDAP ACTIVE_DIRECTORY]
	DirectoryType string `json:"directoryType,omitempty"`

	// domain in LDAP server (e.g. ad.seq.com).
	Domain string `json:"domain,omitempty"`

	// Group Member Attribute (defaults to member)
	GroupMemberAttribute string `json:"groupMemberAttribute,omitempty"`

	// Group Id Attribute (defaults to cn)
	GroupNameAttribute string `json:"groupNameAttribute,omitempty"`

	// Group Object Class (defaults to groupOfNames)
	GroupObjectClass string `json:"groupObjectClass,omitempty"`

	// template for group search for authorization (e.g. dc=hadoop,dc=apache,dc=org)
	GroupSearchBase string `json:"groupSearchBase,omitempty"`

	// public host or IP address of LDAP server
	// Required: true
	Host *string `json:"host"`

	// name of the resource
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// port of LDAP server (typically: 389 or 636 for LDAPS)
	// Required: true
	// Maximum: 65535
	// Minimum: 1
	Port *int32 `json:"port"`

	// determines the protocol (LDAP or LDAP over SSL)
	Protocol string `json:"protocol,omitempty"`

	// template for pattern based user search for authentication (e.g. cn={0},dc=hadoop,dc=apache,dc=org)
	// Required: true
	UserDnPattern *string `json:"userDnPattern"`

	// attribute name for simplified search filter (e.g. sAMAccountName in case of AD, UID or cn for LDAP).
	UserNameAttribute string `json:"userNameAttribute,omitempty"`

	// User Object Class (defaults to person)
	UserObjectClass string `json:"userObjectClass,omitempty"`

	// template for user search for authentication (e.g. dc=hadoop,dc=apache,dc=org)
	// Required: true
	UserSearchBase *string `json:"userSearchBase"`
}

// Validate validates this ldap v4 request
func (m *LdapV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBindDn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBindPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectoryType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserDnPattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserSearchBase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapV4Request) validateBindDn(formats strfmt.Registry) error {

	if err := validate.Required("bindDn", "body", m.BindDn); err != nil {
		return err
	}

	return nil
}

func (m *LdapV4Request) validateBindPassword(formats strfmt.Registry) error {

	if err := validate.Required("bindPassword", "body", m.BindPassword); err != nil {
		return err
	}

	return nil
}

func (m *LdapV4Request) validateDescription(formats strfmt.Registry) error {

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

var ldapV4RequestTypeDirectoryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LDAP","ACTIVE_DIRECTORY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapV4RequestTypeDirectoryTypePropEnum = append(ldapV4RequestTypeDirectoryTypePropEnum, v)
	}
}

const (

	// LdapV4RequestDirectoryTypeLDAP captures enum value "LDAP"
	LdapV4RequestDirectoryTypeLDAP string = "LDAP"

	// LdapV4RequestDirectoryTypeACTIVEDIRECTORY captures enum value "ACTIVE_DIRECTORY"
	LdapV4RequestDirectoryTypeACTIVEDIRECTORY string = "ACTIVE_DIRECTORY"
)

// prop value enum
func (m *LdapV4Request) validateDirectoryTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ldapV4RequestTypeDirectoryTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LdapV4Request) validateDirectoryType(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectoryType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectoryTypeEnum("directoryType", "body", m.DirectoryType); err != nil {
		return err
	}

	return nil
}

func (m *LdapV4Request) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *LdapV4Request) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
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

func (m *LdapV4Request) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *LdapV4Request) validateUserDnPattern(formats strfmt.Registry) error {

	if err := validate.Required("userDnPattern", "body", m.UserDnPattern); err != nil {
		return err
	}

	return nil
}

func (m *LdapV4Request) validateUserSearchBase(formats strfmt.Registry) error {

	if err := validate.Required("userSearchBase", "body", m.UserSearchBase); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapV4Request) UnmarshalBinary(b []byte) error {
	var res LdapV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
