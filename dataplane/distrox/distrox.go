package distrox

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hortonworks/cb-cli/dataplane/api-environment/client/v1env"

	envmodel "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
	"github.com/hortonworks/cb-cli/dataplane/api/client/v1distrox"
	"github.com/hortonworks/cb-cli/dataplane/api/model"
	"github.com/hortonworks/cb-cli/dataplane/common"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/oauth"
	"github.com/hortonworks/cb-cli/dataplane/types"
	"github.com/hortonworks/cb-cli/dataplane/utils"
	commonutils "github.com/hortonworks/dp-cli-common/utils"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var stackHeader = []string{"Name", "Crn", "CloudPlatform", "Environment", "DistroXStatus", "ClusterStatus", "NodeCount"}

type dxOut struct {
	common.CloudResourceOut
	Crn           string `json:"Crn" yaml:"Crn"`
	Environment   string `json:"Environment" yaml:"Environment"`
	DistroXStatus string `json:"DistroXStatus" yaml:"DistroXStatus"`
	ClusterStatus string `json:"ClusterStatus" yaml:"ClusterStatus"`
	NodeCount     string `json:"NodeCount" yaml:"NodeCount"`
}

type stackOutDescribe struct {
	Stack       model.StackV4Response
	Environment envmodel.DetailedEnvironmentV1Response
	Crn         string
}

func (s *dxOut) DataAsStringArray() []string {
	arr := []string{s.Name, s.Crn}
	arr = append(arr, s.CloudPlatform)
	arr = append(arr, s.Environment)
	arr = append(arr, s.DistroXStatus)
	arr = append(arr, s.ClusterStatus)
	arr = append(arr, s.NodeCount)
	return arr
}

func (s *stackOutDescribe) DataAsStringArray() []string {
	stack := convertResponseToDx(s)
	arr := append(stack.DataAsStringArray(), s.Stack.StatusReason)
	return arr
}

func convertResponseToDx(s *stackOutDescribe) *dxOut {
	return &dxOut{
		CloudResourceOut: common.CloudResourceOut{
			Name:          *s.Stack.Name,
			Description:   utils.SafeClusterDescriptionConvert(&s.Stack),
			CloudPlatform: s.Environment.CloudPlatform,
		},
		Crn:           s.Crn,
		Environment:   s.Environment.Name,
		DistroXStatus: s.Stack.Status,
		ClusterStatus: utils.SafeClusterStatusConvert(&s.Stack),
		NodeCount:     strconv.FormatInt(int64(s.Stack.NodeCount), 10),
	}
}

func convertViewResponseToStack(s *model.StackViewV4Response) *dxOut {
	return &dxOut{
		CloudResourceOut: common.CloudResourceOut{
			Name:          *s.Name,
			CloudPlatform: s.CloudPlatform,
		},
		Crn:           s.Crn,
		Environment:   s.EnvironmentName,
		DistroXStatus: s.Status,
		ClusterStatus: utils.SafeClusterViewStatusConvert(s),
		NodeCount:     strconv.FormatInt(int64(s.NodeCount), 10),
	}
}

func DeleteDistroX(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "delete DistroX cluster")

	dxClient := DistroX(*oauth.NewCloudbreakHTTPClientFromContext(c))
	name := c.String(fl.FlName.Name)
	dxClient.deleteDistroX(name, c.Bool(fl.FlForceOptional.Name))
	log.Infof("[DeleteDistroX] DistroX deleted, name: %s", name)

	if c.Bool(fl.FlWaitOptional.Name) {
		dxClient.WaitForDistroXOperationToFinish(name, DELETE_COMPLETED, SKIP)
	}
}

func DeleteMultipleDistroxClusters(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "delete multiple DistroX clusters by name")

	dxClient := DistroX(*oauth.NewCloudbreakHTTPClientFromContext(c))
	val := c.String(fl.FlNames.Name)
	clusterNames := strings.Split(val[1:len(val)-1], ",")
	forced := c.Bool(fl.FlForceOptional.Name)
	log.Infof("[DeleteMultipleDistroxClusters] delete clusters(s) by names: %s", clusterNames)
	request := model.DistroXMultiDeleteV1Request{
		Names: clusterNames,
	}
	err := dxClient.Cloudbreak.V1distrox.DeleteMultipleDistroXClustersByNamesV1(
		v1distrox.NewDeleteMultipleDistroXClustersByNamesV1Params().WithForced(&forced).WithBody(&request))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
}

func assembleDistroXRequest(c *cli.Context) *model.DistroXV1Request {
	path := c.String(fl.FlInputJson.Name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		commonutils.LogErrorAndExit(err)
	}

	log.Infof("[assembleDistroXTemplate] read DistroX create json from file: %s", path)
	content := commonutils.ReadFile(path)

	var req model.DistroXV1Request
	err := json.Unmarshal(content, &req)
	if err != nil {
		msg := fmt.Sprintf(`Invalid json format: %s. Please make sure that the json is valid (check for commas and double quotes).`, err.Error())
		commonutils.LogErrorMessageAndExit(msg)
	}

	name := c.String(fl.FlName.Name)
	if len(name) != 0 {
		req.Name = &name
		if req.InstanceGroups != nil {
			for _, ig := range req.InstanceGroups {
				if ig.Azure != nil && ig.Azure.AvailabilitySet != nil {
					as := ig.Azure.AvailabilitySet
					as.Name = fmt.Sprintf("%s-%s-as", *req.Name, *req.Name)
				}
			}
		}
	}
	if req.Name == nil || len(*req.Name) == 0 {
		commonutils.LogErrorMessageAndExit("Name of the DistroX must be set either in the template or with the --name command line option.")
	}

	cmUser := c.String(fl.FlCMUserOptional.Name)
	cmPassword := c.String(fl.FlCMPasswordOptional.Name)
	if len(cmUser) != 0 || len(cmPassword) != 0 {
		if req.Cluster != nil {
			if len(cmUser) != 0 {
				req.Cluster.UserName = &cmUser
			}
			if len(cmPassword) != 0 {
				req.Cluster.Password = &cmPassword
			}
		} else {
			commonutils.LogErrorMessageAndExit("Missing DistroX node in JSON")
		}
	}
	return &req
}

func ScaleDistroX(c *cli.Context) {
	desiredCount, err := strconv.Atoi(c.String(fl.FlDesiredNodeCount.Name))
	if err != nil {
		commonutils.LogErrorMessageAndExit("Unable to parse as number: " + c.String(fl.FlDesiredNodeCount.Name))
	}
	defer commonutils.TimeTrack(time.Now(), "scale DistroX")

	dxClient := DistroX(*oauth.NewCloudbreakHTTPClientFromContext(c))

	req := &model.DistroXScaleV1Request{
		DesiredCount: &(&types.I32{I: int32(desiredCount)}).I,
		Group:        &(&types.S{S: c.String(fl.FlGroupName.Name)}).S,
	}
	name := c.String(fl.FlName.Name)
	log.Infof("[ScaleDistroX] scaling DistroX, name: %s", name)
	err = dxClient.Cloudbreak.V1distrox.PutScalingDistroXV1ByName(v1distrox.NewPutScalingDistroXV1ByNameParams().WithName(name).WithBody(req))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[ScaleDistroX] DistroX scaled, name: %s", name)

	if c.Bool(fl.FlWaitOptional.Name) {
		dxClient.WaitForDistroXOperationToFinish(name, AVAILABLE, AVAILABLE)
	}
}

func StartDistroX(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "start DistroX")

	dxClient := DistroX(*oauth.NewCloudbreakHTTPClientFromContext(c))
	name := c.String(fl.FlName.Name)
	log.Infof("[StartDistroX] starting DistroX, name: %s", name)
	err := dxClient.Cloudbreak.V1distrox.StartDistroXV1ByName(v1distrox.NewStartDistroXV1ByNameParams().WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[StartDistroX] DistroX started, name: %s", name)

	if c.Bool(fl.FlWaitOptional.Name) {
		dxClient.WaitForDistroXOperationToFinish(name, AVAILABLE, AVAILABLE)
	}
}

func StopDistroX(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "stop DistroX")

	dxClient := DistroX(*oauth.NewCloudbreakHTTPClientFromContext(c))
	name := c.String(fl.FlName.Name)
	log.Infof("[StopDistroX] stopping DistroX, name: %s", name)
	err := dxClient.Cloudbreak.V1distrox.StopDistroXV1ByName(v1distrox.NewStopDistroXV1ByNameParams().WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[StopDistroX] DistroX stopted, name: %s", name)

	if c.Bool(fl.FlWaitOptional.Name) {
		dxClient.WaitForDistroXOperationToFinish(name, STOPPED, STOPPED)
	}
}

func CreateDistroX(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "create DistroX")

	req := assembleDistroXRequest(c)
	dxClient := DistroX(*oauth.NewCloudbreakHTTPClientFromContext(c))
	utils.CheckClientVersion(dxClient.Cloudbreak.V4utils, common.Version)
	dxClient.createDistroX(req)
	if c.Bool(fl.FlWaitOptional.Name) {
		dxClient.WaitForDistroXOperationToFinish(*req.Name, AVAILABLE, AVAILABLE)
	}
}

func ChangeImage(c *cli.Context) {
	//dxClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	//imageId := c.String(fl.FlImageId.Name)
	//imageCatalogName := c.String(fl.FlImageCatalog.Name)
	//dxName := c.String(fl.FlName.Name)
	//log.Infof("[ChangeImage] changing image for DistroX, name: %s, imageid: %s, imageCatalog: %s", dxName, imageId, imageCatalogName)
	//req := model.DistroXImageChangeV1Request{ImageCatalogName: imageCatalogName, ImageID: &imageId}
	//err := dxClient.Cloudbreak.V1distrox.ChangeImageDistroXV1(v1distrox.NewChangeImageDistroXV1Params().WithName(dxName).WithBody(&req))
	//if err != nil {
	//	commonutils.LogErrorAndExit(err)
	//}
}

func DescribeDistroX(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "describe DistroX")

	dxClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	resp, err := dxClient.Cloudbreak.V1distrox.GetDistroXV1ByName(v1distrox.NewGetDistroXV1ByNameParams().WithName(c.String(fl.FlName.Name)))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	s := resp.Payload

	envClient := oauth.Environment(*oauth.NewEnvironmentClientFromContext(c)).Environment
	envResp, err := envClient.V1env.GetEnvironmentV1ByCrn(v1env.NewGetEnvironmentV1ByCrnParams().WithCrn(s.EnvironmentCrn))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}

	output.Write(append(stackHeader, "STATUSREASON"), &stackOutDescribe{
		Stack:       *s,
		Environment: *envResp.Payload,
		Crn:         s.Crn,
	})
}

func ListDistroXs(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "list DistroXs")

	dxClient := DistroX(*oauth.NewCloudbreakHTTPClientFromContext(c))
	output := commonutils.Output{Format: c.String(fl.FlOutputOptional.Name)}

	resp, err := dxClient.Cloudbreak.V1distrox.ListDistroXV1(v1distrox.NewListDistroXV1Params())
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}

	var tableRows []commonutils.Row
	for _, stack := range resp.Payload.Responses {
		tableRows = append(tableRows, convertViewResponseToStack(stack))
	}

	output.WriteList(stackHeader, tableRows)
}

func GetListOfDistroXs(c *cli.Context) *v1distrox.ListDistroXV1OK {
	defer commonutils.TimeTrack(time.Now(), "list DistroXs")

	dxClient := DistroX(*oauth.NewCloudbreakHTTPClientFromContext(c))

	resp, err := dxClient.Cloudbreak.V1distrox.ListDistroXV1(v1distrox.NewListDistroXV1Params())
	if err != nil {
		return nil
	}
	return resp
}

func RepairDistroXHostGroups(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "repair DistroX")

	var request model.DistroXRepairV1Request
	hostGroups := strings.Split(c.String(fl.FlHostGroups.Name), ",")
	request.HostGroups = hostGroups
	log.Infof("[RepairDistroX] repairing DistroX hostgroups: %s", hostGroups)
	repairDistroX(c, request)
}

func RepairDistroXNodes(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "repair DistroX")

	var request model.DistroXRepairV1Request
	deleteVolumes := c.Bool(fl.FlDeleteVolumes.Name)
	nodes := strings.Split(c.String(fl.FlNodes.Name), ",")
	request.Nodes = &model.DistroXRepairNodesV1Request{DeleteVolumes: deleteVolumes, Ids: nodes}

	log.Infof("[RepairDistroXNodes] repairing DistroX nodes, deleteVolumes: %t, ids: %s", deleteVolumes, nodes)

	repairDistroX(c, request)
}

func repairDistroX(c *cli.Context, request model.DistroXRepairV1Request) {
	dxClient := DistroX(*oauth.NewCloudbreakHTTPClientFromContext(c))
	removeOnly := c.Bool(fl.FlRemoveOnly.Name)
	request.RemoveOnly = removeOnly
	name := c.String(fl.FlName.Name)
	log.Infof("[RepairDistroX] repairing DistroX, id: %s, removeOnly: %t", name, removeOnly)

	err := dxClient.Cloudbreak.V1distrox.RepairDistroXV1ByName(v1distrox.NewRepairDistroXV1ByNameParams().WithName(name).WithBody(&request))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[RepairDistroX] DistroX repaired, name: %s", name)

	if c.Bool(fl.FlWaitOptional.Name) {
		dxClient.WaitForDistroXOperationToFinish(name, AVAILABLE, AVAILABLE)
	}
}

func RetryDistroX(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "retry DistroX creation")

	dxClient := DistroX(*oauth.NewCloudbreakHTTPClientFromContext(c))
	name := c.String(fl.FlName.Name)
	log.Infof("[RetryDistroX retrying DistroX creation, name: %s", name)
	err := dxClient.Cloudbreak.V1distrox.RetryDistroXV1ByName(v1distrox.NewRetryDistroXV1ByNameParams().WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[RetryDistroX] DistroX creation retried, name: %s", name)

	if c.Bool(fl.FlWaitOptional.Name) {
		dxClient.WaitForDistroXOperationToFinish(name, AVAILABLE, AVAILABLE)
	}
}

func SyncDistroX(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "sync DistroX")

	dxClient := DistroX(*oauth.NewCloudbreakHTTPClientFromContext(c))
	name := c.String(fl.FlName.Name)
	log.Infof("[SyncDistroX] syncing DistroX, name: %s", name)
	err := dxClient.Cloudbreak.V1distrox.SyncDistroXV1ByName(v1distrox.NewSyncDistroXV1ByNameParams().WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	log.Infof("[SyncDistroX] DistroX synced, name: %s", name)
}

func GetRequestByName(c *cli.Context) {
	defer commonutils.TimeTrack(time.Now(), "getting the CDP CLI request")

	dxClient := DistroX(*oauth.NewCloudbreakHTTPClientFromContext(c))
	name := c.String(fl.FlName.Name)
	log.Infof("[GetRequestByName] getting the CDP CLI request, name: %s", name)
	result, err := dxClient.Cloudbreak.V1distrox.GetDistroXRequestV1ByName(v1distrox.NewGetDistroXRequestV1ByNameParams().WithName(name))
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		commonutils.LogErrorAndExit(err)
	}
	commonutils.Println(string(bytes))
	log.Infof("[GetRequestByName] getting the CDP CLI request, name: %s", name)
}
