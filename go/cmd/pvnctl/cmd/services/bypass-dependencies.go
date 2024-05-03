package services

import (
	"context"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"

	"github.com/spf13/cobra"
)

func bypassDeps(cmd *cobra.Command, args []string) {
	protectionDesiredStateId := args[0]

	ctx := context.Background()
	_, err := cmdutil.GetDesiredStateManagerClient().BypassDependencies(ctx, &ds_pb.BypassDependenciesReq{
		DesiredStateId: protectionDesiredStateId,
		Source:         "user",
	})
	if err != nil {
		cmd.PrintErrf("Failed to bypass dependencies: %+v\n", err)
		os.Exit(1)
	}
}

var bypassDepsCmd = &cobra.Command{
	Use:   "bypass-dependencies <desired-state-id>",
	Short: "Bypass dependencies (protections, release channel dependencies) for the provided desired state id.",
	Long: `Bypass dependencies (protections, release channel dependencies) for the provided desired state id.

pvnctl services --app=<app> bypass-dependencies <desired-state-id>`,
	Args: cobra.ExactArgs(1),
	Run:  bypassDeps,
}

func init() {
	RootCmd.AddCommand(bypassDepsCmd)
}
