package registry

import (
	"context"
	"log"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	workflow_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/workflow"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List container registry integrations.",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := cmdutil.GetWorkflowManagerClient().ListContainerRegistryIntegrations(context.Background(), &workflow_pb.ListContainerRegistryIntegrationsReq{
			FetchStatus: true,
		})
		if err != nil {
			log.Fatal(err)
		}

		listing := cmdutil.NewOutputListing("ID", "NAME", "URL", "TYPE", "STATUS")
		for _, cr := range resp.ContainerRegistries {
			listing.AddRow(cr.IntegrationId, cr.Name, cr.Url, cr.Type, cr.Status)
		}
		err = cmdutil.WriteStdout(listing)
		if err != nil {
			cmd.PrintErrf("Failed to list container registries: %+v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
