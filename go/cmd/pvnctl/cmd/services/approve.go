package services

import (
	"context"
	"log"
	"os"

	"prodvana/cmd/cmdutil"

	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"

	"github.com/spf13/cobra"
)

func approveReject(cmd *cobra.Command, args []string) {
	releaseChannel := args[1]
	desiredStateId := args[2]
	ctx := context.Background()

	var reject bool
	var completedMsg string
	switch cmd.Name() {
	case "approve":
		completedMsg = "approved"
	case "reject":
		reject = true
		completedMsg = "rejected"
	default:
		log.Fatalf("Unexpected command %s", cmd.Name())
	}

	signalType, err := cmd.Flags().GetString("signal-type")
	if err != nil {
		log.Fatal(err)
	}

	_, err = cmdutil.GetDesiredStateManagerClient().SetManualApproval(ctx, &ds_pb.SetManualApprovalReq{
		DesiredStateId: desiredStateId,
		Topic:          releaseChannel,
		Reject:         reject,
		SignalType:     signalType,
	})
	if err != nil {
		cmd.PrintErrf("Failed to %s service instance: %+v\n", cmd.Name(), err)
		os.Exit(1)
	}
	cmd.Println(completedMsg)
}

var approveCmd = &cobra.Command{
	Use:   "approve <service-name> <release-channel> <desired-state>",
	Short: "Approve pending push of a service instance.",
	Long: `Approve pending push of a service instance.

pvnctl services --app=<app> approve <service-name> <release-channel> <root-desired-state>
`,
	Args: cobra.ExactArgs(3),
	Run:  approveReject,
}

var rejectCmd = &cobra.Command{
	Use:   "reject <service-name> <release-channel> <desired-state>",
	Short: "Reject pending push of a service instance.",
	Long: `Reject pending push of a service instance.

pvnctl services --app=<app> reject <service-name> <release-channel> <root-desired-state>
`,
	Args: cobra.ExactArgs(3),
	Run:  approveReject,
}

func init() {
	RootCmd.AddCommand(approveCmd)
	RootCmd.AddCommand(rejectCmd)
	for _, cmd := range []*cobra.Command{approveCmd, rejectCmd} {
		cmd.Flags().String("signal-type", "", "What kind of approval/reject is this?")
	}
}
