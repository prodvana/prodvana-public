package registry

import (
	"context"
	"log"

	"prodvana/cmd/cmdutil"

	workflow_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/workflow"

	"github.com/spf13/cobra"
)

var rmRepoCmd = &cobra.Command{
	Use:     "rm-repository",
	Aliases: []string{"rm-repo"},
	Short:   "Stop indexing this repository in Prodvana.",
	Long: `Stop indexing this repository in Prodvana.
pvnctl integration registry rm-repository --integration-id <integration id> <repository>
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		integrationId, err := cmd.Flags().GetString("integration-id")
		if err != nil {
			log.Fatal(err)
		}
		repository := args[0]

		_, err = cmdutil.GetWorkflowManagerClient().StopTrackingImageRepository(context.Background(), &workflow_pb.StopTrackingImageRepositoryReq{
			IntegrationId: integrationId,
			Repository:    repository,
		})
		if err != nil {
			log.Fatal(err)
		}

		cmd.Printf("Stopped indexing %s\n", repository)
	},
}

func init() {
	RootCmd.AddCommand(rmRepoCmd)
	rmRepoCmd.Flags().String("integration-id", "", "Integration ID of registry to use")
	cmdutil.Must(rmRepoCmd.MarkFlagRequired("integration-id"))
}
