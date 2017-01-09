package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*SecurityRuleRequest security rule request

swagger:model SecurityRuleRequest
*/
type SecurityRuleRequest struct {

	/* flag for making the rule modifiable
	 */
	Modifiable *bool `json:"modifiable,omitempty"`

	/* comma separated list of accessible ports

	Required: true
	Pattern: ^[1-9][0-9]{0,4}(-[1-9][0-9]{0,4}){0,1}(,[1-9][0-9]{0,4}(-[1-9][0-9]{0,4}){0,1})*$
	*/
	Ports string `json:"ports"`

	/* protocol of the rule

	Required: true
	*/
	Protocol string `json:"protocol"`

	/* definition of allowed subnet in CIDR format

	Required: true
	Pattern: ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/([0-9]|[1-2][0-9]|3[0-2]))$
	*/
	Subnet string `json:"subnet"`
}

// Validate validates this security rule request
func (m *SecurityRuleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePorts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityRuleRequest) validatePorts(formats strfmt.Registry) error {

	if err := validate.RequiredString("ports", "body", string(m.Ports)); err != nil {
		return err
	}

	if err := validate.Pattern("ports", "body", string(m.Ports), `^[1-9][0-9]{0,4}(-[1-9][0-9]{0,4}){0,1}(,[1-9][0-9]{0,4}(-[1-9][0-9]{0,4}){0,1})*$`); err != nil {
		return err
	}

	return nil
}

func (m *SecurityRuleRequest) validateProtocol(formats strfmt.Registry) error {

	if err := validate.RequiredString("protocol", "body", string(m.Protocol)); err != nil {
		return err
	}

	return nil
}

func (m *SecurityRuleRequest) validateSubnet(formats strfmt.Registry) error {

	if err := validate.RequiredString("subnet", "body", string(m.Subnet)); err != nil {
		return err
	}

	if err := validate.Pattern("subnet", "body", string(m.Subnet), `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/([0-9]|[1-2][0-9]|3[0-2]))$`); err != nil {
		return err
	}

	return nil
}
