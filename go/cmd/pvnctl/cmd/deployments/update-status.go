package deployments

import (
	"context"
	"log"
	"strings"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	deployment_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/deployment"
	deployment_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/deployment/model"
	"github.com/spf13/cobra"
)

var updateStatusFlags = struct {
	status string
}{}

var updateStatusCmd = &cobra.Command{
	Use:   "update-status",
	Short: "Updates status of a pending deployment.",
	Long: `Updates status of a pending deployment.

ID=$(pvnctl deployment record --deployment-system=circleci --pending)
pvnctl deployment update-status --status=success $ID
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		status := deployment_model_pb.DeploymentStatus(deployment_model_pb.DeploymentStatus_value[strings.ToUpper(updateStatusFlags.status)])
		_, err := cmdutil.GetDeploymentManagerClient().UpdateDeploymentStatus(ctx, &deployment_pb.UpdateDeploymentStatusReq{
			DeploymentId: args[0],
			Status:       status,
		})
		if err != nil {
			log.Fatal("failed to update deployment status", err)
		}
		log.Printf("Updated %s to %v\n", args[0], status)
	},
}

func init() {
	RootCmd.AddCommand(updateStatusCmd)
	updateStatusCmd.Flags().StringVarP(&updateStatusFlags.status, "status", "s", "", "Status. Allowed values: succeeded, failed")
	cmdutil.Must(updateStatusCmd.MarkFlagRequired("status"))
}
