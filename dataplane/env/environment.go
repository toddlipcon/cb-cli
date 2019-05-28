package env

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/dataplane/api-environment/model"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/oauth"
	"github.com/hortonworks/dp-cli-common/utils"
	"github.com/urfave/cli"

	v1env "github.com/hortonworks/cb-cli/dataplane/api-environment/client/v1env"
)

var EnvironmentHeader = []string{"Name", "Description", "CloudPlatform", "Credential", "Regions", "LocationName", "Longitude", "Latitude"}

type environment struct {
	Name          string   `json:"Name" yaml:"Name"`
	Description   string   `json:"Description" yaml:"Description"`
	CloudPlatform string   `json:"CloudPlatform" yaml:"CloudPlatform"`
	Credential    string   `json:"Credential" yaml:"Credential"`
	Regions       []string `json:"Regions" yaml:"Regions"`
	LocationName  string   `json:"LocationName" yaml:"LocationName"`
	Longitude     float64  `json:"Longitude" yaml:"Longitude"`
	Latitude      float64  `json:"Latitude" yaml:"Latitude"`
}

type environmentOutTableDescribe struct {
	*environment
	ID string `json:"ID" yaml:"ID"`
}

type environmentOutJsonDescribe struct {
	*environment
	LdapConfigs     []string                           `json:"LdapConfigs" yaml:"LdapConfigs"`
	ProxyConfigs    []string                           `json:"ProxyConfigs" yaml:"ProxyConfigs"`
	KerberosConfigs []string                           `json:"KerberosConfigs" yaml:"KerberosConfigs"`
	RdsConfigs      []string                           `json:"RdsConfigs" yaml:"RdsConfigs"`
	ID              string                             `json:"ID" yaml:"ID"`
	Network         model.EnvironmentNetworkV1Response `json:"Network" yaml:"Network"`
}

type environmentListJsonDescribe struct {
	*environment
	Network model.EnvironmentNetworkV1Response `json:"Network" yaml:"Network"`
}

type environmentClient interface {
	CreateEnvironmentV1(params *v1env.CreateEnvironmentV1Params) (*v1env.CreateEnvironmentV1OK, error)
	ListEnvironmentV1(params *v1env.ListEnvironmentV1Params) (*v1env.ListEnvironmentV1OK, error)
}

func (e *environment) DataAsStringArray() []string {
	return []string{e.Name, e.Description, e.CloudPlatform, e.Credential, strings.Join(e.Regions, ","), e.LocationName, utils.FloatToString(e.Longitude), utils.FloatToString(e.Latitude)}
}

func (e *environmentOutJsonDescribe) DataAsStringArray() []string {
	return append(e.environment.DataAsStringArray(), e.ID)
}

func (e *environmentOutTableDescribe) DataAsStringArray() []string {
	return append(e.environment.DataAsStringArray(), e.ID)
}

func CreateEnvironment(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "create environment")

	name := c.String(fl.FlName.Name)
	description := c.String(fl.FlDescriptionOptional.Name)
	credentialName := c.String(fl.FlEnvironmentCredential.Name)
	regions := utils.DelimitedStringToArray(c.String(fl.FlEnvironmentRegions.Name), ",")
	locationName := c.String(fl.FlEnvironmentLocationName.Name)
	longitude := c.Float64(fl.FlEnvironmentLongitudeOptional.Name)
	latitude := c.Float64(fl.FlEnvironmentLatitudeOptional.Name)

	EnvironmentV1Request := &model.EnvironmentV1Request{
		Name:           &name,
		Description:    &description,
		CredentialName: credentialName,
		Regions:        regions,
		Location: &model.LocationV1Request{
			Name:      &locationName,
			Longitude: longitude,
			Latitude:  latitude,
		},
	}

	createEnvironmentImpl(c, EnvironmentV1Request)
}

func CreateEnvironmentFromTemplate(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "create environment from template")
	fileLocation := c.String(fl.FlEnvironmentTemplateFile.Name)
	log.Infof("[assembleStackTemplate] read environment template JSON from file: %s", fileLocation)
	content := utils.ReadFile(fileLocation)

	var req model.EnvironmentV1Request
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid JSON format: %s. Please make sure that the json is valid (check for commas and double quotes).`, err.Error())
		utils.LogErrorMessageAndExit(msg)
	}

	if name := c.String(fl.FlEnvironmentNameOptional.Name); len(name) != 0 {
		req.Name = &name
	} else if req.Name == nil {
		utils.LogErrorMessageAndExit("Name of the environment must be set either in the template or with the --name command line option.")
	}
	createEnvironmentImpl(c, &req)
}

func createEnvironmentImpl(c *cli.Context, EnvironmentV1Request *model.EnvironmentV1Request) {
	log.Infof("[createEnvironmentImpl] create environment with name: %s", *EnvironmentV1Request.Name)

	envClient := oauth.NewEnvironmentClientFromContext(c)
	resp, err := envClient.Environment.V1env.CreateEnvironmentV1(v1env.NewCreateEnvironmentV1Params().WithBody(EnvironmentV1Request))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	environment := resp.Payload

	log.Infof("[createEnvironmentImpl] environment created with name: %s, id: %d", *EnvironmentV1Request.Name, environment.ID)
}

func GenerateAwsEnvironmentTemplate(c *cli.Context) error {
	template := createEnvironmentWithNetwork(c)
	template.Network.Aws = &model.EnvironmentNetworkAwsV1Params{
		VpcID: new(string),
	}
	return printTemplate(template)
}

func GenerateAzureEnvironmentTemplate(c *cli.Context) error {
	template := createEnvironmentWithNetwork(c)
	template.Network.Azure = &model.EnvironmentNetworkAzureV1Params{
		NetworkID:         new(string),
		NoFirewallRules:   *new(bool),
		NoPublicIP:        *new(bool),
		ResourceGroupName: new(string),
	}
	return printTemplate(template)
}

func createEnvironmentWithNetwork(c *cli.Context) model.EnvironmentV1Request {
	template := model.EnvironmentV1Request{
		Name:           new(string),
		Description:    new(string),
		CredentialName: "____",
		Regions:        make([]string, 0),
		Location: &model.LocationV1Request{
			Name:      new(string),
			Longitude: 0,
			Latitude:  0,
		},
		Network: &model.EnvironmentNetworkV1Request{
			SubnetIds: make([]string, 0),
		},
	}
	if credName := c.String(fl.FlEnvironmentCredentialOptional.Name); len(credName) != 0 {
		template.CredentialName = credName
	}
	if locationName := c.String(fl.FlEnvironmentLocationNameOptional.Name); len(locationName) != 0 {
		template.Location.Name = &locationName
	}
	if regions := utils.DelimitedStringToArray(c.String(fl.FlEnvironmentRegions.Name), ","); len(regions) != 0 {
		template.Regions = regions
	}
	return template
}

func printTemplate(template model.EnvironmentV1Request) error {
	resp, err := json.MarshalIndent(template, "", "\t")
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	fmt.Printf("%s\n", string(resp))
	return nil
}

func ListEnvironments(c *cli.Context) error {
	defer utils.TimeTrack(time.Now(), "list environments")

	envClient := oauth.NewEnvironmentClientFromContext(c)

	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	return listEnvironmentsImpl(envClient.Environment.V1env, output)
}

func listEnvironmentsImpl(envClient environmentClient, output utils.Output) error {
	resp, err := envClient.ListEnvironmentV1(v1env.NewListEnvironmentV1Params())
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	var tableRows []utils.Row
	for _, e := range resp.Payload.Responses {
		row := &environment{
			Name:          e.Name,
			Description:   e.Description,
			CloudPlatform: e.CloudPlatform,
			Credential:    e.CredentialName,
			Regions:       getRegionNames(e.Regions),
			LocationName:  e.Location.Name,
			Longitude:     e.Location.Longitude,
			Latitude:      e.Location.Latitude,
		}

		if output.Format != "table" && output.Format != "yaml" && e.Network != nil {
			tableRows = append(tableRows, &environmentListJsonDescribe{
				environment: row,
				Network:     *e.Network,
			})
		} else {
			tableRows = append(tableRows, row)
		}
	}

	if output.Format != "table" && output.Format != "yaml" {
		output.WriteList(append(EnvironmentHeader, "Network"), tableRows)
	} else {
		output.WriteList(EnvironmentHeader, tableRows)
	}
	return nil
}

func DescribeEnvironment(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "describe an environment")

	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	envName := c.String(fl.FlName.Name)
	log.Infof("[DescribeEnvironment] describe environment by name: %s", envName)
	envClient := oauth.NewEnvironmentClientFromContext(c)

	resp, err := envClient.Environment.V1env.GetEnvironmentV1(v1env.NewGetEnvironmentV1Params().WithName(envName))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	env := resp.Payload
	if output.Format != "table" && output.Format != "yaml" {
		output.Write(append(EnvironmentHeader, "ID", "Network"), convertResponseToJsonOutput(env))
	} else {
		output.Write(append(EnvironmentHeader, "ID"), convertResponseToTableOutput(env))
	}
}

func DeleteEnvironment(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "delete an environment")
	envName := c.String(fl.FlName.Name)
	log.Infof("[DeleteEnvironment] delete environment by name: %s", envName)
	envClient := oauth.NewEnvironmentClientFromContext(c)
	_, err := envClient.Environment.V1env.DeleteEnvironmentV1(v1env.NewDeleteEnvironmentV1Params().WithName(envName))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
}

func ChangeCredential(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "change credential of environment")
	envName := c.String(fl.FlName.Name)
	credential := c.String(fl.FlCredential.Name)

	requestBody := &model.EnvironmentChangeCredentialV1Request{
		CredentialName: credential,
	}
	request := v1env.NewChangeCredentialInEnvironmentV1Params().WithName(envName).WithBody(requestBody)
	envClient := oauth.NewEnvironmentClientFromContext(c)
	resp, err := envClient.Environment.V1env.ChangeCredentialInEnvironmentV1(request)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	environment := resp.Payload
	log.Infof("[ChangeCredential] credential of environment %s changed to: %s", environment.Name, environment.CredentialName)
}

func EditEnvironment(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "edit environment")
	envName := c.String(fl.FlName.Name)
	description := c.String(fl.FlDescriptionOptional.Name)
	regions := utils.DelimitedStringToArray(c.String(fl.FlEnvironmentRegions.Name), ",")
	locationName := c.String(fl.FlEnvironmentLocationNameOptional.Name)
	longitude := c.Float64(fl.FlEnvironmentLongitudeOptional.Name)
	latitude := c.Float64(fl.FlEnvironmentLatitudeOptional.Name)

	requestBody := &model.EnvironmentEditV1Request{
		Description: &description,
		Regions:     regions,
		Location: &model.LocationV1Request{
			Name:      &locationName,
			Longitude: longitude,
			Latitude:  latitude,
		},
	}
	request := v1env.NewEditEnvironmentV1Params().WithName(envName).WithBody(requestBody)
	envClient := oauth.NewEnvironmentClientFromContext(c)
	resp, err := envClient.Environment.V1env.EditEnvironmentV1(request)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	environment := resp.Payload
	log.Infof("[Edit] Environment %s was edited.", environment.Name)
}

func convertResponseToTableOutput(env *model.DetailedEnvironmentV1Response) *environmentOutTableDescribe {
	return &environmentOutTableDescribe{
		environment: &environment{
			Name:          env.Name,
			Description:   env.Description,
			CloudPlatform: env.CloudPlatform,
			Credential:    env.CredentialName,
			Regions:       getRegionNames(env.Regions),
			LocationName:  env.Location.Name,
			Longitude:     env.Location.Longitude,
			Latitude:      env.Location.Latitude,
		},
		ID: env.ID,
	}
}

func convertResponseToJsonOutput(env *model.DetailedEnvironmentV1Response) *environmentOutJsonDescribe {
	result := &environmentOutJsonDescribe{
		environment: &environment{
			Name:          env.Name,
			Description:   env.Description,
			CloudPlatform: env.CloudPlatform,
			Credential:    env.CredentialName,
			Regions:       getRegionNames(env.Regions),
			LocationName:  env.Location.Name,
			Longitude:     env.Location.Longitude,
			Latitude:      env.Location.Latitude,
		},
		ID: env.ID,
	}
	if env.Network != nil {
		result.Network = *env.Network
	}
	return result
}

func getRegionNames(region *model.CompactRegionV1Response) []string {
	var regions []string
	for _, v := range region.Regions {
		regions = append(regions, v)
	}
	return regions
}

func getProxyConfigNames(configs []*model.ProxyResponse) []string {
	var names []string
	for _, c := range configs {
		names = append(names, *c.Name)
	}
	return names
}
