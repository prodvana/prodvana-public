package slack

import (
	"context"
	"log"

	"prodvana/cmd/cmdutil"

	workflow_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/workflow"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use: "install",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := cmdutil.GetWorkflowManagerClient().GetInstallSlackUrl(context.Background(), &workflow_pb.GetInstallSlackUrlReq{})
		if err != nil {
			log.Fatal(err)
		}

		cmd.Printf("To install the slack app, please visit %s\n\n", resp.Url)
	},
}

func init() {
	RootCmd.AddCommand(installCmd)
}
