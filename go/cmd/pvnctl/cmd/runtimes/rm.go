package runtimes

import (
	"context"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environment"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm <name>",
	Short: "Remove a runtime from Prodvana.",
	Long: `Remove a runtime from Prodvana.

pvnctl runtimes rm <name>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ctx := context.Background()
		_, err := cmdutil.GetEnvironmentManagerClient().RemoveCluster(ctx,
			&environment.RemoveClusterReq{Name: name})
		if err != nil {
			cmd.PrintErrf("Failed to remove cluster: %+v\n", err)
			os.Exit(1)
		}
		cmd.Printf("Successfully removed cluster %s\n", name)
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
