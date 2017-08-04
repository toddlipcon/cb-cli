package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*LdapValidationRequest ldap validation request

swagger:model LdapValidationRequest
*/
type LdapValidationRequest struct {

	/* bind distinguished name for connection test and group search (e.g. cn=admin,dc=example,dc=org)

	Required: true
	*/
	BindDn string `json:"bindDn"`

	/* password for the provided bind DN

	Required: true
	*/
	BindPassword string `json:"bindPassword"`

	/* determines the protocol (LDAP or LDAP over SSL)
	 */
	Protocol *string `json:"protocol,omitempty"`

	/* public host or IP address of LDAP server

	Required: true
	*/
	ServerHost string `json:"serverHost"`

	/* port of LDAP server (typically: 389 or 636 for LDAPS)

	Required: true
	Maximum: 65535
	Minimum: 1
	*/
	ServerPort int32 `json:"serverPort"`
}

// Validate validates this ldap validation request
func (m *LdapValidationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBindDn(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBindPassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServerHost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServerPort(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapValidationRequest) validateBindDn(formats strfmt.Registry) error {

	if err := validate.RequiredString("bindDn", "body", string(m.BindDn)); err != nil {
		return err
	}

	return nil
}

func (m *LdapValidationRequest) validateBindPassword(formats strfmt.Registry) error {

	if err := validate.RequiredString("bindPassword", "body", string(m.BindPassword)); err != nil {
		return err
	}

	return nil
}

func (m *LdapValidationRequest) validateServerHost(formats strfmt.Registry) error {

	if err := validate.RequiredString("serverHost", "body", string(m.ServerHost)); err != nil {
		return err
	}

	return nil
}

func (m *LdapValidationRequest) validateServerPort(formats strfmt.Registry) error {

	if err := validate.Required("serverPort", "body", int32(m.ServerPort)); err != nil {
		return err
	}

	if err := validate.MinimumInt("serverPort", "body", int64(m.ServerPort), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("serverPort", "body", int64(m.ServerPort), 65535, false); err != nil {
		return err
	}

	return nil
}
