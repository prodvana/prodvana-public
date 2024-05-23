package services

import (
	"context"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"

	"github.com/spf13/cobra"
)

func forceExecuteTask(cmd *cobra.Command, args []string) {
	runtimeExtensionDsId := args[0]

	ctx := context.Background()
	_, err := cmdutil.GetDesiredStateManagerClient().ForceExecuteTask(ctx, &ds_pb.ForceExecuteTaskReq{
		DesiredStateId: runtimeExtensionDsId,
		Source:         "user",
	})
	if err != nil {
		cmd.PrintErrf("Failed to force execute task: %+v\n", err)
		os.Exit(1)
	}
}

var forceExecuteTaskCmd = &cobra.Command{
	Use:   "force-execute-task <desired-state-id>",
	Short: "Force execution of a runtime extension now.",
	Long: `Force execution of a runtime extension task now.

pvnctl services --app=<app> force-execute-task <desired-state-id>`,
	Args: cobra.ExactArgs(1),
	Run:  forceExecuteTask,
}

func init() {
	RootCmd.AddCommand(forceExecuteTaskCmd)
}
