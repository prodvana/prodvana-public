package grafana

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
		url, err := cmd.Flags().GetString("instance-url")
		if err != nil {
			log.Fatal(err)
		}
		token, err := cmd.Flags().GetString("api-token")
		if err != nil {
			log.Fatal(err)
		}
		resp, err := cmdutil.GetWorkflowManagerClient().InstallGrafana(context.Background(), &workflow_pb.InstallGrafanaReq{
			Url:      url,
			ApiToken: token,
		})
		if err != nil {
			log.Fatal(err)
		}
		cmd.Printf("Successfuly installed the Grafana integration %s\n", resp.IntegrationId)
	},
}

func init() {
	RootCmd.AddCommand(installCmd)
	installCmd.Flags().String("instance-url", "", "Your grafana instance URL")
	cmdutil.Must(installCmd.MarkFlagRequired("instance-url"))
	installCmd.Flags().String("api-token", "", "API token to use to authenticate against your instance")
	cmdutil.Must(installCmd.MarkFlagRequired("api-token"))
}
