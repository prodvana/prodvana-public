package services

import (
	"context"
	"os"

	"prodvana/cmd/cmdutil"

	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"

	"github.com/spf13/cobra"
)

func bypassConcurrencyLimit(cmd *cobra.Command, args []string) {
	desiredStateId := args[0]

	ctx := context.Background()
	_, err := cmdutil.GetDesiredStateManagerClient().BypassConcurrencyLimit(ctx, &ds_pb.BypassConcurrencyLimitReq{
		DesiredStateId: desiredStateId,
		Source:         "user",
	})
	if err != nil {
		cmd.PrintErrf("Failed to bypass concurrency limit: %+v\n", err)
		os.Exit(1)
	}
}

var bypassConcurrencyLimitCmd = &cobra.Command{
	Use:   "bypass-concurrency-limit <desired-state-id>",
	Short: "Bypass concurrency limits for the provided desired state id",
	Long: `Bypass concurrency limits for the provided desired state id.

pvnctl services --app=<app> bypass-concurrency-limit <desired-state-id>`,
	Args: cobra.ExactArgs(1),
	Run:  bypassConcurrencyLimit,
}

func init() {
	RootCmd.AddCommand(bypassConcurrencyLimitCmd)
}
