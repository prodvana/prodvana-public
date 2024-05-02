package registry

import (
	"context"
	"log"
	"os"

	"prodvana/cmd/cmdutil"

	workflow_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/workflow"

	"github.com/spf13/cobra"
)

var listReposCmd = &cobra.Command{
	Use:     "list-repositories",
	Aliases: []string{"list-repos"},
	Short:   "List repositories for a given registry.",
	Run: func(cmd *cobra.Command, args []string) {
		integrationID, err := cmd.Flags().GetString("integration-id")
		if err != nil {
			log.Fatal(err)
		}
		resp, err := cmdutil.GetWorkflowManagerClient().ListTrackedImageRepositories(context.Background(), &workflow_pb.ListTrackedImageRepositoriesReq{
			IntegrationId: integrationID,
		})
		if err != nil {
			log.Fatal(err)
		}

		listing := cmdutil.NewOutputListing("REPOSITORY", "INDEX STATUS", "LAST UPDATE")
		for _, repo := range resp.Repositories {
			updateTime := "Never"
			if repo.LastIndex.IsValid() {
				updateTime = repo.LastIndex.AsTime().String()
			}
			listing.AddRow(repo.Repository, repo.IndexStatus, updateTime)
		}
		err = cmdutil.WriteStdout(listing)
		if err != nil {
			cmd.PrintErrf("Failed to list repositoriees: %+v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(listReposCmd)
	listReposCmd.Flags().String("integration-id", "", "Integration ID of registry to use")
	cmdutil.Must(listReposCmd.MarkFlagRequired("integration-id"))
}
