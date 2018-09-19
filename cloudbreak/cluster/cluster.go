package cluster

import (
	"github.com/hortonworks/cb-cli/cloudbreak/oauth"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/cloudbreak/api/client/v3_workspace_id_stacks"
	"github.com/hortonworks/cb-cli/cloudbreak/api/model"
	fl "github.com/hortonworks/cb-cli/cloudbreak/flags"
	"github.com/hortonworks/cb-cli/cloudbreak/types"
	"github.com/hortonworks/cb-cli/utils"
	"github.com/urfave/cli"
)

func ChangeAmbariPassword(c *cli.Context) {
	fl.CheckRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "update ambari password")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	name := c.String(fl.FlName.Name)
	log.Infof("[ChangeAmbariPassword] updating ambari password, name: %s", name)
	req := &model.UserNamePassword{
		OldPassword: &(&types.S{S: c.String(fl.FlOldPassword.Name)}).S,
		Password:    &(&types.S{S: c.String(fl.FlNewPassword.Name)}).S,
		UserName:    &(&types.S{S: c.String(fl.FlAmbariUser.Name)}).S,
	}
	err := cbClient.Cloudbreak.V3WorkspaceIDStacks.PutpasswordStackV3(v3_workspace_id_stacks.NewPutpasswordStackV3Params().WithWorkspaceID(workspaceID).WithName(name).WithBody(req))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[RepairStack] ambari password updated, name: %s", name)
}