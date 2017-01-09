package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*ClusterRequest cluster request

swagger:model ClusterRequest
*/
type ClusterRequest struct {

	/* details of the external Ambari database
	 */
	AmbariDatabaseDetails *AmbariDatabaseDetails `json:"ambariDatabaseDetails,omitempty"`

	/* details of the Ambari package repository
	 */
	AmbariRepoDetailsJSON *AmbariRepoDetails `json:"ambariRepoDetailsJson,omitempty"`

	/* details of the Ambari stack
	 */
	AmbariStackDetails *AmbariStackDetails `json:"ambariStackDetails,omitempty"`

	// TODO WARNING: do not replace it with string, otherwise it cannot be serialized
	/* blueprint custom properties
	 */
	BlueprintCustomProperties []Configurations `json:"blueprintCustomProperties,omitempty"`

	/* blueprint id for the cluster

	Required: true
	*/
	BlueprintID int64 `json:"blueprintId"`

	/* blueprint inputs in the cluster

	Unique: true
	*/
	BlueprintInputs []*BlueprintInput `json:"blueprintInputs,omitempty"`

	/* config recommendation strategy
	 */
	ConfigStrategy *string `json:"configStrategy,omitempty"`

	/* description of the resource

	Max Length: 1000
	Min Length: 0
	*/
	Description *string `json:"description,omitempty"`

	/* send email about the result of the cluster installation
	 */
	EmailNeeded *bool `json:"emailNeeded,omitempty"`

	/* send email to the requested address
	 */
	EmailTo *string `json:"emailTo,omitempty"`

	/* enable Kerberos security
	 */
	EnableSecurity *bool `json:"enableSecurity,omitempty"`

	/* shipyard service enabled in the cluster
	 */
	EnableShipyard *bool `json:"enableShipyard,omitempty"`

	/* external file system configuration
	 */
	FileSystem *FileSystem `json:"fileSystem,omitempty"`

	/* collection of hostgroups

	Unique: true
	*/
	HostGroups []*HostGroupRequest `json:"hostGroups,omitempty"`

	/* kerberos
	 */
	Kerberos *KerberosRequest `json:"kerberos,omitempty"`

	/* LDAP config id for the cluster
	 */
	LdapConfigID *int64 `json:"ldapConfigId,omitempty"`

	/* flag for default LDAP support
	 */
	LdapRequired *bool `json:"ldapRequired,omitempty"`

	/* name of the resource

	Required: true
	Max Length: 40
	Min Length: 5
	Pattern: ([a-z][-a-z0-9]*[a-z0-9])
	*/
	Name string `json:"name"`

	/* ambari password

	Required: true
	Max Length: 100
	Min Length: 5
	*/
	Password string `json:"password"`

	/* RDS configuration id for the cluster
	 */
	RdsConfigID *int64 `json:"rdsConfigId,omitempty"`

	/* details of the external database for Hadoop components
	 */
	RdsConfigJSON *RDSConfig `json:"rdsConfigJson,omitempty"`

	/* SSSD config id for the cluster
	 */
	SssdConfigID *int64 `json:"sssdConfigId,omitempty"`

	/* ambari username

	Required: true
	Max Length: 15
	Min Length: 5
	Pattern: ([a-z][-a-z0-9]*[a-z0-9])
	*/
	UserName string `json:"userName"`

	/* validate blueprint
	 */
	ValidateBlueprint *bool `json:"validateBlueprint,omitempty"`
}

// Validate validates this cluster request
func (m *ClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlueprintID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBlueprintInputs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConfigStrategy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHostGroups(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterRequest) validateBlueprintID(formats strfmt.Registry) error {

	if err := validate.Required("blueprintId", "body", int64(m.BlueprintID)); err != nil {
		return err
	}

	return nil
}

func (m *ClusterRequest) validateBlueprintInputs(formats strfmt.Registry) error {

	if swag.IsZero(m.BlueprintInputs) { // not required
		return nil
	}

	if err := validate.UniqueItems("blueprintInputs", "body", m.BlueprintInputs); err != nil {
		return err
	}

	for i := 0; i < len(m.BlueprintInputs); i++ {

		if m.BlueprintInputs[i] != nil {

			if err := m.BlueprintInputs[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

var clusterRequestTypeConfigStrategyPropEnum []interface{}

func (m *ClusterRequest) validateConfigStrategyEnum(path, location string, value string) error {
	if clusterRequestTypeConfigStrategyPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["NEVER_APPLY","ONLY_STACK_DEFAULTS_APPLY","ALWAYS_APPLY","ALWAYS_APPLY_DONT_OVERRIDE_CUSTOM_VALUES"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			clusterRequestTypeConfigStrategyPropEnum = append(clusterRequestTypeConfigStrategyPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, clusterRequestTypeConfigStrategyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterRequest) validateConfigStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigStrategy) { // not required
		return nil
	}

	if err := m.validateConfigStrategyEnum("configStrategy", "body", *m.ConfigStrategy); err != nil {
		return err
	}

	return nil
}

func (m *ClusterRequest) validateDescription(formats strfmt.Registry) error {

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

func (m *ClusterRequest) validateHostGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.HostGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("hostGroups", "body", m.HostGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.HostGroups); i++ {

		if m.HostGroups[i] != nil {

			if err := m.HostGroups[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ClusterRequest) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 40); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `([a-z][-a-z0-9]*[a-z0-9])`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterRequest) validatePassword(formats strfmt.Registry) error {

	if err := validate.RequiredString("password", "body", string(m.Password)); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", string(m.Password), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(m.Password), 100); err != nil {
		return err
	}

	return nil
}

func (m *ClusterRequest) validateUserName(formats strfmt.Registry) error {

	if err := validate.RequiredString("userName", "body", string(m.UserName)); err != nil {
		return err
	}

	if err := validate.MinLength("userName", "body", string(m.UserName), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("userName", "body", string(m.UserName), 15); err != nil {
		return err
	}

	if err := validate.Pattern("userName", "body", string(m.UserName), `([a-z][-a-z0-9]*[a-z0-9])`); err != nil {
		return err
	}

	return nil
}
