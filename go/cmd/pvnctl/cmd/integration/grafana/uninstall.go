package grafana

import (
	"context"
	"log"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	workflow_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/workflow"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use: "uninstall",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := cmdutil.GetWorkflowManagerClient().UninstallGrafana(context.Background(), &workflow_pb.UninstallGrafanaReq{})
		if err != nil {
			log.Fatal(err)
		}

		cmd.Printf("Successfully uninstalled the Grafana integration %s\n", resp.IntegrationId)
	},
}

func init() {
	RootCmd.AddCommand(uninstallCmd)
}
