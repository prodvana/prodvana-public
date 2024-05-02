package services

import (
	"context"
	go_errors "errors"
	"log"
	"os"
	"prodvana/cmd/cmdutil"
	"time"

	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"
	"github.com/spf13/cobra"
)

var waitDesiredState = &cobra.Command{
	Use:   "wait-desired-state",
	Short: "Wait for a desired state to either converge or fail to converge",
	Long: `Wait for a desired state to either converge or fail to converge. Exit 0 on convergence success, 3 if timed out, 2 if failed, 1 otherwise.

pvnctl services --app <app> wait-desired-state <service>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		exitOnManualApproval, err := cmd.Flags().GetBool("wait-for-manual-approval")
		if err != nil {
			log.Fatal(err)
		}

		waitForStatus, err := cmd.Flags().GetString("wait-for-status")
		if err != nil {
			log.Fatal(err)
		}
		var requiredStatus ds_model_pb.Status
		if waitForStatus != "" {
			requiredStatus = ds_model_pb.Status(ds_model_pb.Status_value[waitForStatus])
		}

		if exitOnManualApproval && waitForStatus != "" {
			log.Fatal("Cannot pass both --wait-for-manual-approval and --wait-for-status")
		}

		timeout, err := cmd.Flags().GetDuration("timeout")
		if err != nil {
			log.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		s := &desiredStateStatus{service: name, exitOnManualApproval: exitOnManualApproval, exitOnExactStatus: requiredStatus}
		err = cmdutil.RenderWatch(ctx, s.renderDesiredState)
		if err != nil {
			cmd.PrintErrf("%s\n", err)
			if ctx.Err() == context.DeadlineExceeded {
				os.Exit(3)
			}
			if go_errors.Is(err, errDesiredStateConvergenceFailed) {
				os.Exit(2)
			}
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(waitDesiredState)
	waitDesiredState.Flags().Bool("wait-for-manual-approval", false, "Also exit if any entities are waiting for manual approval.")
	waitDesiredState.Flags().Duration("timeout", 7*time.Minute, "Timeout to wait for")
	waitDesiredState.Flags().String("wait-for-status", "", "If set, wait for entity status to reach this exact value.")
}
