package services

import (
	"context"
	go_errors "errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/bradenaw/juniper/xslices"
	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/async_task"
	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/maestro"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type releaseStatus struct {
	service              string
	exitOnManualApproval bool
	exitOnExactStatus    ds_model_pb.Status
}

func (s *releaseStatus) handleReturn(pending *ds_pb.PendingSetDesiredState, status ds_model_pb.Status, serviceInstanceStates []*ds_model_pb.Entity) (bool, error) {
	if pending != nil {
		switch pending.TaskStatus {
		case async_task.TaskStatus_FAILED:
			return false, errors.Wrapf(errDesiredStateConvergenceFailed, "Pending desired state failed to apply. %s", string(pending.TaskResult.GetLog()))
		default: // wait for pending states to go away and become latest
			return false, nil
		}
	}
	switch status {
	case ds_model_pb.Status_FAILED, ds_model_pb.Status_FAILED_ROLLBACK:
		return false, errors.Wrap(errDesiredStateConvergenceFailed, "Desired state failed to converge")
	case ds_model_pb.Status_ROLLED_BACK:
		return false, errors.Wrap(errDesiredStateConvergenceFailed, "Desired state rolled back")
	}

	var waitingForManualApproval bool
	// check status of each release channel, as partial roll backs show up as converged
	for _, node := range serviceInstanceStates {
		switch node.GetId().GetType() {
		case ds_model_pb.Type_SERVICE_INSTANCE:
			var explanationMsgs []string
			for _, expl := range node.StatusExplanations {
				explanationMsgs = append(explanationMsgs, fmt.Sprintf("%s on %s/%s: %s", expl.Reason, expl.Subject.GetType(), expl.Subject.GetName(), expl.Message))
			}
			switch node.Status {
			case ds_model_pb.Status_WAITING_MANUAL_APPROVAL:
				waitingForManualApproval = true
			case ds_model_pb.Status_ROLLED_BACK:
				return false, errors.Wrapf(errDesiredStateConvergenceFailed, "Desired state rolled back.\n%s", strings.Join(explanationMsgs, "\n"))
			case ds_model_pb.Status_FAILED_ROLLBACK, ds_model_pb.Status_FAILED:
				return false, errors.Wrapf(errDesiredStateConvergenceFailed, "Desired state failed to converge.\n%s", strings.Join(explanationMsgs, "\n"))
			}
		}
	}

	if s.exitOnExactStatus != 0 {
		return status == s.exitOnExactStatus, nil
	}

	if waitingForManualApproval && s.exitOnManualApproval {
		return true, nil
	}
	return status == ds_model_pb.Status_CONVERGED, nil
}

func (s *releaseStatus) renderDesiredState(ctx context.Context, writer io.Writer) (bool, error) {
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

	var dsId string
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

	rcIdMap := make(map[string]string)
	for _, rcConfig := range release.MaestroRelease.Config.CompiledDesiredState.GetService().GetReleaseChannels() {
		rcIdMap[rcConfig.ReleaseChannelId] = rcConfig.ReleaseChannel
	}

	resp, err := cmdutil.GetDesiredStateManagerClient().GetDesiredStateGraph(ctx, &ds_pb.GetDesiredStateGraphReq{
		Query: &ds_pb.GetDesiredStateGraphReq_DesiredStateId{
			DesiredStateId: dsId,
		},
		Types: []ds_model_pb.Type{ds_model_pb.Type_SERVICE, ds_model_pb.Type_SERVICE_INSTANCE},
		Depth: 0,
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
		fmt.Fprintf(writer, "Release channels:\n")

		getReleaseChannelName := func(siState *ds_model_pb.Entity) string {
			for _, annotation := range siState.Annotations.Annotations {
				if annotation.Key == "release_channel_id" {
					rcName, ok := rcIdMap[annotation.Value]
					if ok {
						return rcName
					}
				}
			}

			return "unknown release channel"
		}

		rcStatuses := make(map[string]ds_model_pb.Status)
		rcNames := make([]string, 0, len(serviceInstanceStates))
		for _, siState := range serviceInstanceStates {
			rcName := getReleaseChannelName(siState)
			rcStatuses[rcName] = siState.Status
			rcNames = append(rcNames, rcName)
		}

		sort.Strings(rcNames)

		for _, rcName := range rcNames {
			cmdutil.ColorForStatus(rcStatuses[rcName]).Fprintf(writer, "  %s: %s\n", rcName, rcStatuses[rcName])
		}
	}

	return s.handleReturn(resp.PendingSetDesiredState, status, serviceInstanceStates)
}

var getRelease = &cobra.Command{
	Use:   "get-release-state",
	Short: "Get the release summary for a service",
	Long: `Get the release summary for a service.

pvnctl services --app <app> get-release-state <service>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ctx := context.Background()
		s := &releaseStatus{service: name}

		_, err := s.renderDesiredState(ctx, os.Stdout)

		if err != nil {
			if !go_errors.Is(err, errDesiredStateConvergenceFailed) {
				cmd.PrintErrf("%s\n", err)
				os.Exit(1)
			}
		}
	},
}

var waitReleaseState = &cobra.Command{
	Use:   "wait-release-state",
	Short: "Wait for a release state to either converge or fail to converge",
	Long: `Wait for a release state to either converge or fail to converge. Exit 0 on convergence success, 3 if timed out, 2 if failed, 1 otherwise.

pvnctl services --app <app> wait-release-state <service>
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
	RootCmd.AddCommand(getRelease)
	RootCmd.AddCommand(waitReleaseState)
	waitReleaseState.Flags().Bool("wait-for-manual-approval", false, "Also exit if any entities are waiting for manual approval.")
	waitReleaseState.Flags().Duration("timeout", 7*time.Minute, "Timeout to wait for")
	waitReleaseState.Flags().String("wait-for-status", "", "If set, wait for entity status to reach this exact value.")
}
