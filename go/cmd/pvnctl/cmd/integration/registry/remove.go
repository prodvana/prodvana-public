package registry

import (
	"context"
	"log"

	"prodvana/cmd/cmdutil"

	workflow_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/workflow"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove this registry integration from Prodvana.",
	Long: `Remove this registry integration from Prodvana.
pvnctl integration registry remove <registry name>
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		registryName := args[0]
		_, err := cmdutil.GetWorkflowManagerClient().DeleteContainerRegistryIntegration(context.Background(), &workflow_pb.DeleteContainerRegistryIntegrationReq{
			RegistryName: registryName,
		})
		if err != nil {
			log.Fatal(err)
		}

		cmd.Printf("Removed registry integration %s\n", registryName)
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
