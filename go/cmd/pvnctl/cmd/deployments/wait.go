package deployments

import (
	"context"
	go_errors "errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	maestro_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/maestro"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"
	"github.com/spf13/cobra"
)

var (
	errDesiredStateConvergenceFailed = fmt.Errorf("Failed to converge")
)

type releaseStatus struct {
	releaseId          string
	releaseChannelName string
	exitOnExactStatus  ds_model_pb.Status
}

func findEntity(graph *ds_model_pb.EntityGraph, dsId string) (*ds_model_pb.Entity, error) {
	for _, entity := range graph.Entities {
		if entity.DesiredStateId == dsId {
			return entity, nil
		}
	}
	return nil, errors.Errorf("%s not found", dsId)
}

func getReleaseChannelId(maestroConfig *maestro_pb.MaestroReleaseConfig, releaseChannelName string) (string, error) {
	if maestroConfig.CompiledDesiredState.GetService() == nil {
		return "", errors.Errorf("Release is not a service release")
	}

	for _, rc := range maestroConfig.CompiledDesiredState.GetService().ReleaseChannels {
		if rc.ReleaseChannel == releaseChannelName {
			return rc.ReleaseChannelId, nil
		}
	}
	return "", errors.Errorf("Release channel %s not found", releaseChannelName)
}

func (s *releaseStatus) renderDesiredState(ctx context.Context, writer io.Writer) (bool, error) {
	release, err := cmdutil.GetDesiredStateManagerClient().GetMaestroRelease(ctx, &ds_pb.GetMaestroReleaseReq{
		ReleaseId: s.releaseId,
	})
	if err != nil {
		return false, errors.Wrap(err, "Failed to get release")
	}

	rcId, err := getReleaseChannelId(release.MaestroRelease.Config, s.releaseChannelName)
	if err != nil {
		return false, errors.Wrapf(err, "Release channel %s not present in release", s.releaseChannelName)
	}

	rcState, ok := release.MaestroRelease.State.ReleaseChannelStates[rcId]
	if !ok {
		color.New(color.FgYellow).Fprintf(writer, "Status: PENDING\n")
		return false, nil
	}
	dsId := rcState.DesiredStateId
	if len(dsId) == 0 {
		color.New(color.FgYellow).Fprintf(writer, "Status: PENDING\n")
		return false, nil
	}
	fmt.Fprintf(writer, "Release status: %s\n", release.MaestroRelease.State.Status)
	fmt.Fprintf(writer, "Entity status: %s\n", rcState.Status)

	// If the release was skipped, check if the release channel had a terminal state.
	// If yes, return that. Otherwise, abort and indicate that release was skipped.
	// This can happen when a release gets skipped and we don't explicitly set
	// release channel state as SKIPPED.
	if release.MaestroRelease.State.Status == maestro_pb.Status_SKIPPED {
		switch rcState.Status {
		case maestro_pb.Status_FAILED, maestro_pb.Status_SUCCEEDED:
		default:
			return false, nil
		}
	}

	resp, err := cmdutil.GetDesiredStateManagerClient().GetDesiredStateGraph(ctx, &ds_pb.GetDesiredStateGraphReq{
		Query: &ds_pb.GetDesiredStateGraphReq_DesiredStateId{
			DesiredStateId: dsId,
		},
		Depth: -1,
	})
	if err != nil {
		return false, errors.Wrap(err, "Failed to get desired state for service")
	}

	if resp.PendingSetDesiredState != nil {
		fmt.Fprintf(writer, "Desired state ID: %s\n", dsId)
		color.New(color.FgYellow).Fprintf(writer, "Status: PENDING\n")
		return false, nil
	}

	entity, err := findEntity(resp.EntityGraph, dsId)
	if err != nil {
		return false, err
	}
	fmt.Fprintf(writer, "Desired state ID: %s\n", dsId)
	fmt.Fprintf(writer, "Entity ID: %s\n", entity.Id)
	cmdutil.ColorForStatus(entity.Status).Fprintf(writer, "Status: %s (%s)\n", entity.Status, rcState.Status)
	if len(entity.StatusExplanations) > 0 {
		fmt.Fprintf(writer, "Status explanations: %s\n", entity.StatusExplanations)
	}

	if s.exitOnExactStatus != 0 {
		return entity.Status == s.exitOnExactStatus, nil
	}

	var explanationMsgs []string
	for _, expl := range entity.StatusExplanations {
		explanationMsgs = append(explanationMsgs, fmt.Sprintf("%s on %s/%s: %s", expl.Reason, expl.Subject.GetType(), expl.Subject.GetName(), expl.Message))
	}
	switch entity.Status {
	case ds_model_pb.Status_CONVERGED:
		return true, nil
	case ds_model_pb.Status_FAILED:
		return false, errors.Wrapf(errDesiredStateConvergenceFailed, "Desired state failed to converge.\n%s", strings.Join(explanationMsgs, "\n"))
	case ds_model_pb.Status_ROLLED_BACK:
		return false, errors.Wrapf(errDesiredStateConvergenceFailed, "Desired state rolled back.\n%s", strings.Join(explanationMsgs, "\n"))
	}

	return false, nil
}

var waitCmd = &cobra.Command{
	Use:   "wait",
	Short: "Wait for release to converge or fail.",
	Long: `Wait for a release to either converge or fail to converge. Exit 0 on convergence success, 3 if timed out, 2 if failed, 1 otherwise.

pvnctl releases wait <releaseid> <releasechannelname>
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		releaseId := args[0]
		releaseChannelId := args[1]
		timeout, err := cmd.Flags().GetDuration("timeout")
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

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		s := &releaseStatus{releaseId: releaseId, releaseChannelName: releaseChannelId, exitOnExactStatus: requiredStatus}
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
	RootCmd.AddCommand(waitCmd)
	waitCmd.Flags().Duration("timeout", 7*time.Minute, "Timeout to wait for")
	waitCmd.Flags().String("wait-for-status", "", "If set, wait for entity status to reach this exact value.")
}
