package services

import (
	"context"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"

	"github.com/spf13/cobra"
)

func bypass(cmd *cobra.Command, args []string) {
	protectionDesiredStateId := args[0]

	ctx := context.Background()
	_, err := cmdutil.GetDesiredStateManagerClient().BypassProtection(ctx, &ds_pb.BypassProtectionReq{
		DesiredStateId: protectionDesiredStateId,
		Source:         "user",
	})
	if err != nil {
		cmd.PrintErrf("Failed to bypass protection: %+v\n", err)
		os.Exit(1)
	}
}

var bypassCmd = &cobra.Command{
	Use:   "bypass-protection <proection-desired-state-id>",
	Short: "Bypass the protection with the provided desired state id.",
	Long: `Bypass the protection with the provided desired state id.

pvnctl services --app=<app> bypass-protection <protection-desired-state-id>`,
	Args: cobra.ExactArgs(1),
	Run:  bypass,
}

func init() {
	RootCmd.AddCommand(bypassCmd)
}
