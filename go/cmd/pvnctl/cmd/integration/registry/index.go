package registry

import (
	"context"
	"log"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	workflow_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/workflow"

	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:     "index-repository",
	Aliases: []string{"index-repo", "index"},
	Short:   "Queue this repository to be indexed in Prodvana.",
	Long: `Queue this repository to be indexed in Prodvana.
pvnctl integration registry index-repository --integration-id <integration id> <repository>
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		integrationId, err := cmd.Flags().GetString("integration-id")
		if err != nil {
			log.Fatal(err)
		}
		repository := args[0]

		_, err = cmdutil.GetWorkflowManagerClient().TrackImageRepositories(context.Background(), &workflow_pb.TrackImageRepositoriesReq{
			IntegrationId: integrationId,
			Repositories: []string{
				repository,
			},
		})
		if err != nil {
			log.Fatal(err)
		}

		cmd.Printf("Queued %s for indexing\n", repository)
	},
}

func init() {
	RootCmd.AddCommand(indexCmd)
	indexCmd.Flags().String("integration-id", "", "Integration ID of registry to use")
	cmdutil.Must(indexCmd.MarkFlagRequired("integration-id"))
}
