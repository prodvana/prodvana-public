package services

import (
	"context"
	go_errors "errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"time"

	"github.com/bradenaw/juniper/xslices"
	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/maestro"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"
	"golang.org/x/exp/maps"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (s *releaseStatus) renderDesiredStateCompact(ctx context.Context, writer io.Writer) (bool, error) {
	releasesList, err := cmdutil.GetDesiredStateManagerClient().ListServiceCombinedReleases(ctx, &ds_pb.ListServiceCombinedReleasesReq{
		Service:     s.service,
		Application: cmdutil.App,
		Descending:  true,
		PageSize:    1,
	})
	if err != nil {
		return false, errors.Wrap(err, "Failed to get desired state for service")
	}

	if len(releasesList.Releases) == 0 {
		return false, errors.New("No releases found")
	}

	latestRelease := releasesList.Releases[0]
	var dsId string

	switch latestRelease.ReleaseOneof.(type) {
	case *ds_pb.ListCombinedReleasesResp_Release_ReleaseId:
		latestReleaseId := releasesList.Releases[0].GetReleaseId()

		release, err := cmdutil.GetDesiredStateManagerClient().GetMaestroRelease(ctx, &ds_pb.GetMaestroReleaseReq{
			ReleaseId: latestReleaseId,
		})
		if err != nil {
			return false, errors.Wrap(err, "Failed to get release")
		}

		if release.MaestroRelease.Config.MaestroConfig.Strategy != maestro.Strategy_IMMEDIATE {
			return false, errors.New("Only immediate strategy is supported")
		}

		switch release.MaestroRelease.State.Status {
		case maestro.Status_FAILED:
			fmt.Fprintf(writer, "Release id: %s\n", latestReleaseId)
			return false, errors.Wrap(errDesiredStateConvergenceFailed, "Release failed")
		}

		if len(release.MaestroRelease.State.ReleaseChannelStates) == 0 {
			fmt.Fprintf(writer, "Release pending: %s\n", latestReleaseId)
			return false, nil
		}

		for _, rcState := range release.MaestroRelease.State.ReleaseChannelStates {
			// Pick any release channel state to get the desired state id
			dsId = rcState.DesiredStateId
			if dsId != "" {
				break
			}
		}

		if dsId == "" {
			fmt.Fprintf(writer, "Release pending: %s\n", latestReleaseId)
			return false, nil
		}

		fmt.Fprintf(writer, "Release id: %s\n", latestReleaseId)

	case *ds_pb.ListCombinedReleasesResp_Release_DesiredStateId:
		dsId = latestRelease.GetDesiredStateId()
	}

	resp, err := cmdutil.GetDesiredStateManagerClient().GetDesiredStateGraph(ctx, &ds_pb.GetDesiredStateGraphReq{
		Query: &ds_pb.GetDesiredStateGraphReq_DesiredStateId{
			DesiredStateId: dsId,
		},
		Types:                         []ds_model_pb.Type{ds_model_pb.Type_SERVICE, ds_model_pb.Type_SERVICE_INSTANCE},
		Depth:                         0,
		IncludeDesiredStateTimestamps: true,
	})
	if err != nil {
		return false, errors.Wrap(err, "Failed to get desired state for service")
	}

	status := ds_model_pb.Status_UNKNOWN_STATUS
	var serviceInstanceStates []*ds_model_pb.Entity

	if resp.PendingSetDesiredState != nil {
		fmt.Fprintf(writer, "Pending desired state: %s\n", resp.PendingSetDesiredState.DesiredStateId)
		fmt.Fprintf(writer, "Pending desired state queue status: %s\n", resp.PendingSetDesiredState.TaskStatus)
	} else {
		//nolint:staticcheck // xslices.Filter(s, func) is way simpler to read than slices.DeleteFunc(slices.Clone(s), inverseFilterFunc)
		serviceStatuses := xslices.Filter(resp.EntityGraph.Entities, func(e *ds_model_pb.Entity) bool {
			return e.GetId().GetType() == ds_model_pb.Type_SERVICE
		})
		if len(serviceStatuses) != 1 {
			return false, fmt.Errorf("Expected exactly 1 service status, found %d", len(serviceStatuses))
		}
		serviceStatus := serviceStatuses[0]
		status = serviceStatus.Status

		//nolint:staticcheck // xslices.Filter(s, func) is way simpler to read than slices.DeleteFunc(slices.Clone(s), inverseFilterFunc)
		serviceInstanceStates = xslices.Filter(resp.EntityGraph.Entities, func(e *ds_model_pb.Entity) bool {
			return e.GetId().GetType() == ds_model_pb.Type_SERVICE_INSTANCE
		})

		cmdutil.ColorForStatus(serviceStatus.Status).Fprintf(writer, "Status: %s\n", serviceStatus.Status)
		dsStarted := resp.Timestamps.CreationTimestamp
		dsUpdated := resp.Timestamps.LastUpdateTimestamp
		if dsStarted != nil && dsUpdated != nil {
			fmt.Fprintf(
				writer,
				"Desired state created %s ago (last updated %s ago)\n",
				time.Since(dsStarted.AsTime()).Round(time.Second),
				time.Since(dsUpdated.AsTime()).Round(time.Second),
			)
		}
		fmt.Fprintf(writer, "Release channels by status:\n")

		rcsByStatus := make(map[ds_model_pb.Status]int)
		for _, siState := range serviceInstanceStates {
			rcsByStatus[siState.Status]++
		}

		keys := maps.Keys(rcsByStatus)
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})

		for _, status := range keys {
			cmdutil.ColorForStatus(status).Fprintf(writer, "  %s: %d\n", status, rcsByStatus[status])
		}
	}

	return s.handleReturn(resp.PendingSetDesiredState, status, serviceInstanceStates)
}

var waitReleaseCompact = &cobra.Command{
	Use:   "wait-release-compact",
	Short: "Wait for a release state to either converge or fail to converge",
	Long: `Wait for a release state to either converge or fail to converge. Exit 0 on convergence success, 3 if timed out, 2 if failed, 1 otherwise.

pvnctl services --app <app> wait-release-compact <service>
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
		s := &releaseStatus{service: name, exitOnManualApproval: exitOnManualApproval, exitOnExactStatus: requiredStatus}
		err = cmdutil.RenderWatch(ctx, s.renderDesiredStateCompact)
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
	RootCmd.AddCommand(waitReleaseCompact)
	waitReleaseCompact.Flags().Bool("wait-for-manual-approval", false, "Also exit if any entities are waiting for manual approval.")
	waitReleaseCompact.Flags().Duration("timeout", 7*time.Minute, "Timeout to wait for")
	waitReleaseCompact.Flags().String("wait-for-status", "", "If set, wait for entity status to reach this exact value.")
}
