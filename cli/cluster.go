package cli

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/sequenceiq/hdc-cli/client/stacks"
	"github.com/sequenceiq/hdc-cli/models"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sync"
)

type ClusterSkeleton struct {
	ClusterName              string `json:"ClusterName" yaml:"ClusterName"`
	HDPVersion               string `json:"HDPVersion" yaml:"HDPVersion"`
	ClusterType              string `json:"ClusterType" yaml:"ClusterType"`
	MasterInstanceType       string `json:"MasterInstanceType" yaml:"MasterInstanceType"`
	WorkerInstanceType       string `json:"WorkerInstanceType" yaml:"WorkerInstanceType"`
	InstanceCount            int32  `json:"InstanceCount" yaml:"InstanceCount"`
	SSHKeyName               string `json:"SSHKeyName" yaml:"SSHKeyName"`
	RemoteAccess             string `json:"RemoteAccess" yaml:"RemoteAccess"`
	WebAccess                bool   `json:"WebAccess" yaml:"WebAccess"`
	ClusterAndAmbariUser     string `json:"ClusterAndAmbariUser" yaml:"ClusterAndAmbariUser"`
	ClusterAndAmbariPassword string `json:"ClusterAndAmbariPassword" yaml:"ClusterAndAmbariPassword"`

	//InstanceRole             string `json:"InstanceRole" yaml:"InstanceRole"`
	//HiveMetastoreUrl         string `json:"HiveMetastoreUrl" yaml:"HiveMetastoreUrl"`
	//HiveMetastoreUser        string `json:"HiveMetastoreUser" yaml:"HiveMetastoreUser"`
	//HiveMetastorePassword    string `json:"HiveMetastorePassword" yaml:"HiveMetastorePassword"`
}

func (c ClusterSkeleton) Json() string {
	j, _ := json.Marshal(c)
	return string(j)
}

func (c ClusterSkeleton) JsonPretty() string {
	j, _ := json.MarshalIndent(c, "", "  ")
	return string(j)
}

func (c ClusterSkeleton) Yaml() string {
	j, _ := yaml.Marshal(c)
	return string(j)
}

func ListClusters(c *cli.Context) error {
	client := NewOAuth2HTTPClient(c.String(FlCBServer.Name), c.String(FlCBUsername.Name), c.String(FlCBPassword.Name))

	resp, err := client.Cloudbreak.Stacks.GetStacksUser(&stacks.GetStacksUserParams{})
	if err != nil {
		log.Error(err)
	}

	for _, v := range resp.Payload {
		fmt.Printf("%s\n", v.Name)
	}
	return nil
}

func CreateCluster(c *cli.Context) error {
	path := c.String(FlCBInputJson.Name)
	if len(path) == 0 {
		log.Errorf("[CreateCluster] missing parameter: %s", FlCBInputJson.Name)
		return newExitError()
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Errorf("[CreateCluster] %s", err.Error())
		return newExitError()
	}

	log.Infof("[CreateCluster] read cluster create json from file: %s", path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Errorf("[CreateCluster] %s", err.Error())
		return newExitError()
	}

	var skeleton ClusterSkeleton
	err = json.Unmarshal(content, &skeleton)
	if err != nil {
		log.Errorf("[CreateCluster] %s", err.Error())
		return newExitError()
	}
	log.Infof("[CreateCluster] assemble cluster based on skeleton: %s", skeleton.Json())

	client := NewOAuth2HTTPClient(c.String(FlCBServer.Name), c.String(FlCBUsername.Name), c.String(FlCBPassword.Name))

	var wg sync.WaitGroup
	wg.Add(4)

	credentialId := make(chan int64, 1)
	go client.CreateCredential(skeleton, credentialId, &wg)

	templateIds := make(chan int64, 2)
	go client.CreateTemplate(skeleton, templateIds, &wg)

	var secGroupId = make(chan int64, 1)
	if skeleton.WebAccess {
		go client.CreateSecurityGroup(skeleton, secGroupId, &wg)
	} else {
		wg.Done()
	}

	networkId := make(chan int64, 1)
	go client.CreateNetwork(skeleton, networkId, &wg)

	wg.Wait()

	// create stack

	func() {
		failurePolicy := models.FailurePolicy{AdjustmentType: "BEST_EFFORT"}
		failureAction := "DO_NOTHING"
		masterType := "GATEWAY"
		workerType := "CORE"
		ambariVersion := "2.4"
		secGroupId := <-secGroupId
		platform := "AWS"

		instanceGroups := []*models.InstanceGroup{
			{
				Group:           "master",
				TemplateID:      <-templateIds,
				NodeCount:       1,
				Type:            &masterType,
				SecurityGroupID: secGroupId,
			},
			{
				Group:           "worker",
				TemplateID:      <-templateIds,
				NodeCount:       skeleton.InstanceCount - 1,
				Type:            &workerType,
				SecurityGroupID: secGroupId,
			},
		}

		var stackParameters = make(map[string]string)
		stackParameters["instanceProfileStrategy"] = "CREATE"

		orchestrator := models.OrchestratorRequest{Type: "SALT"}

		stackReq := models.StackRequest{
			Name:            skeleton.ClusterName,
			CredentialID:    <-credentialId,
			Region:          "eu-central-1",
			FailurePolicy:   &failurePolicy,
			OnFailureAction: &failureAction,
			InstanceGroups:  instanceGroups,
			Parameters:      stackParameters,
			CloudPlatform:   platform,
			PlatformVariant: &platform,
			NetworkID:       <-networkId,
			AmbariVersion:   &ambariVersion,
			HdpVersion:      &skeleton.HDPVersion,
			Orchestrator:    &orchestrator,
		}

		log.Infof("[CreateStack] sending stack create request with name: %s", skeleton.ClusterName)
		resp, err := client.Cloudbreak.Stacks.PostStacksUser(&stacks.PostStacksUserParams{&stackReq})

		if err != nil {
			log.Errorf("[CreateStack] %s", err.Error())
			newExitError()
		}

		log.Infof("[CreateStack] stack created, id: %d", resp.Payload.ID)
	}()

	return nil
}

func GenerateCreateClusterSkeleton(c *cli.Context) error {
	fmt.Println(ClusterSkeleton{}.JsonPretty())
	return nil
}