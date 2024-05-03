package runtimes

import (
	"context"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environment"
	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:    "pause <name>",
	Hidden: true,
	Short:  "Pause a runtime in Prodvana.",
	Long: `Pause a runtime in Prodvana.

pvnctl runtimes pause <name>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ctx := context.Background()
		_, err := cmdutil.GetEnvironmentManagerClient().PauseCluster(ctx, &environment.PauseClusterReq{
			Runtime: name,
		})
		if err != nil {
			cmd.PrintErrf("Failed to pause cluster: %+v\n", err)
			os.Exit(1)
		}
		cmd.Printf("Successfully paused cluster %s\n", name)
	},
}

var resumeCmd = &cobra.Command{
	Use:    "resume <name>",
	Hidden: true,
	Short:  "Resume a runtime in Prodvana.",
	Long: `Resume a runtime in Prodvana.

pvnctl runtimes resume <name>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ctx := context.Background()
		_, err := cmdutil.GetEnvironmentManagerClient().ResumeCluster(ctx, &environment.ResumeClusterReq{
			Runtime: name,
		})
		if err != nil {
			cmd.PrintErrf("Failed to resume cluster: %+v\n", err)
			os.Exit(1)
		}
		cmd.Printf("Successfully resume cluster %s\n", name)
	},
}

func init() {
	RootCmd.AddCommand(pauseCmd)
	RootCmd.AddCommand(resumeCmd)
}
