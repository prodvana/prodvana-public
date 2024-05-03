package services

import (
	"context"
	"log"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"

	"github.com/spf13/cobra"
)

func promote(cmd *cobra.Command, args []string) {
	service := args[0]
	releaseChannel := args[1]
	full, err := cmd.Flags().GetBool("full")
	if err != nil {
		log.Fatal(err)
	}
	stage, err := cmd.Flags().GetInt("stage")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	summaryResp, err := cmdutil.GetDesiredStateManagerClient().GetServiceDesiredStateConvergenceSummary(ctx, &ds_pb.GetServiceDesiredStateConvergenceSummaryReq{
		Service:                service,
		Application:            cmdutil.App,
		FastNoDeprecatedFields: true,
	})
	if err != nil {
		cmd.PrintErrf("Failed to get current state for service: %v", err)
		os.Exit(1)
	}

	var latestState *ds_model_pb.ServiceInstanceState
	for _, ds := range summaryResp.Summary.LastSeenState.GetService().ReleaseChannels {
		if ds.ReleaseChannel == releaseChannel {
			latestState = ds
			break
		}
	}
	if latestState == nil {
		cmd.PrintErrf("Could not find state for Release Channel '%s'", releaseChannel)
		os.Exit(1)
	}

	if latestState.Delivery == nil {
		cmd.PrintErr("No delivery state found")
		os.Exit(1)
	}

	if !full && len(latestState.Delivery.CanaryProgress) <= stage {
		cmd.PrintErrf("Invalid stage %d, delivery has %d stages", stage, len(latestState.Delivery.CanaryProgress))
		os.Exit(1)
	}

	_, err = cmdutil.GetDesiredStateManagerClient().PromoteDelivery(ctx, &ds_pb.PromoteDeliveryReq{
		DesiredStateId: latestState.Delivery.DesiredStateId,
		Stage:          int64(stage),
		Full:           full,
		Source:         "user",
	})
	if err != nil {
		cmd.PrintErrf("Failed to %s service instance: %+v\n", cmd.Name(), err)
		os.Exit(1)
	}
}

var promoteCmd = &cobra.Command{
	Use:   "promote <service-name> <release-channel>",
	Short: "Promote Progressive Delivery state.",
	Long: `Promote Progressive Delivery state.

pvnctl services --app=<app> promote (--full | --stage=<stage>) <service-name> <release-channel>`,
	Args: cobra.ExactArgs(2),
	Run:  promote,
}

func init() {
	RootCmd.AddCommand(promoteCmd)
	promoteCmd.Flags().Int("stage", 0, "Delivery stage to promote (e.g. skip).")
	promoteCmd.Flags().Bool("full", false, "Fully promote the Delivery to 100%. Takes precedent over --stage if both are provided.")
}
