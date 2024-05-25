package services

import (
	"context"
	"encoding/json"
	go_errors "errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/async_task"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	errDesiredStateConvergenceFailed = fmt.Errorf("Failed to converge")
)

type desiredStateStatus struct {
	service              string
	exitOnManualApproval bool
	exitOnExactStatus    ds_model_pb.Status
}

func (s *desiredStateStatus) handleReturn(summary *ds_pb.DesiredStateSummary) (bool, error) {
	if summary.GetPendingSetDesiredState() != nil {
		switch summary.PendingSetDesiredState.TaskStatus {
		case async_task.TaskStatus_FAILED:
			return false, errors.Wrapf(errDesiredStateConvergenceFailed, "Pending desired state failed to apply. %s", string(summary.PendingSetDesiredState.TaskResult.GetLog()))
		default: // wait for pending states to go away and become latest
			return false, nil
		}
	}
	switch summary.GetStatus() {
	case ds_model_pb.Status_FAILED, ds_model_pb.Status_FAILED_ROLLBACK:
		return false, errors.Wrap(errDesiredStateConvergenceFailed, "Desired state failed to converge")
	case ds_model_pb.Status_ROLLED_BACK:
		return false, errors.Wrap(errDesiredStateConvergenceFailed, "Desired state rolled back")
	}

	var waitingForManualApproval bool
	// check status of each release channel, as partial roll backs show up as converged
	for _, node := range summary.EntityGraph.GetEntities() {
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
		return summary.Status == s.exitOnExactStatus, nil
	}

	if waitingForManualApproval && s.exitOnManualApproval {
		return true, nil
	}
	return summary.Status == ds_model_pb.Status_CONVERGED, nil
}

func getServiceInstanceNodesSorted(graph *ds_model_pb.EntityGraph) []*ds_model_pb.Entity {
	var siNodes []*ds_model_pb.Entity
	for _, node := range graph.GetEntities() {
		switch node.GetId().GetType() {
		case ds_model_pb.Type_SERVICE_INSTANCE:
			siNodes = append(siNodes, node)
		}
	}
	sort.Slice(siNodes, func(i, j int) bool {
		return siNodes[i].DesiredState.GetServiceInstance().GetReleaseChannel() < siNodes[j].DesiredState.GetServiceInstance().GetReleaseChannel()
	})
	return siNodes
}

func (s *desiredStateStatus) renderDesiredState(ctx context.Context, writer io.Writer) (bool, error) {
	resp, err := cmdutil.GetDesiredStateManagerClient().GetServiceDesiredStateConvergenceSummary(ctx, &ds_pb.GetServiceDesiredStateConvergenceSummaryReq{
		Service:                s.service,
		Application:            cmdutil.App,
		FastNoDeprecatedFields: true,
	})
	if err != nil {
		return false, errors.Wrap(err, "Failed to get desired state for service")
	}
	if resp.Summary.PendingSetDesiredState != nil {
		fmt.Fprintf(writer, "Pending desired state: %s\n", resp.Summary.PendingSetDesiredState.DesiredStateId)
		fmt.Fprintf(writer, "Pending desired state queue status: %s\n", resp.Summary.PendingSetDesiredState.TaskStatus)
	} else {
		fmt.Fprintf(writer, "Desired state: %s\n", resp.Summary.DesiredState.GetService().Meta.RootDesiredStateId)
		fmt.Fprintf(writer, "Status: %s\n", resp.Summary.Status)
		fmt.Fprintf(writer, "Release channels:\n")

		siNodes := getServiceInstanceNodesSorted(resp.Summary.EntityGraph)

		for _, node := range siNodes {
			ds := node.DesiredState.GetServiceInstance()
			fmt.Fprintf(writer, "%s: %s\n", ds.GetReleaseChannel(), node.Status)
		}
	}

	return s.handleReturn(resp.Summary)
}

func (s *desiredStateStatus) renderDesiredStateJSON(ctx context.Context, writer io.Writer) (bool, error) {
	resp, err := cmdutil.GetDesiredStateManagerClient().GetServiceDesiredStateConvergenceSummary(ctx, &ds_pb.GetServiceDesiredStateConvergenceSummaryReq{
		Service:                s.service,
		Application:            cmdutil.App,
		FastNoDeprecatedFields: true,
	})
	if err != nil {
		return false, errors.Wrap(err, "Failed to get desired state for service")
	}

	status := resp.Summary.Status

	type resultT struct {
		Name           string `json:"name"`
		DesiredStateId string `json:"desired_state_id"`
		Status         string `json:"status"`
		DesiredVersion string `json:"desired_version"`
	}

	var result struct {
		DesiredState       string    `json:"desired_state"`
		Status             string    `json:"status"`
		PendingStatus      string    `json:"pending_status"`
		ReleaseChannels    []resultT `json:"release_channels"`
		CustomTasks        []resultT `json:"custom_tasks"`
		DeliveryExtensions []resultT `json:"delivery_extensions"`
	}
	if resp.Summary.PendingSetDesiredState != nil {
		result.DesiredState = resp.Summary.PendingSetDesiredState.DesiredStateId
		result.PendingStatus = resp.Summary.PendingSetDesiredState.TaskStatus.String()
	} else {
		result.DesiredState = resp.Summary.DesiredState.GetService().Meta.RootDesiredStateId
		result.Status = status.String()
		for _, node := range resp.Summary.EntityGraph.GetEntities() {
			switch node.GetId().GetType() {
			case ds_model_pb.Type_SERVICE_INSTANCE:
				ds := node.DesiredState.GetServiceInstance()
				result.ReleaseChannels = append(result.ReleaseChannels, resultT{
					Name:           ds.ReleaseChannel,
					DesiredStateId: ds.Meta.DesiredStateId,
					Status:         node.Status.String(),
					DesiredVersion: ds.Versions[0].Version,
				})
			case ds_model_pb.Type_DELIVERY_EXTENSION:
				ds := node.DesiredState.GetDeliveryExtension()
				result.DeliveryExtensions = append(result.DeliveryExtensions, resultT{
					Name:           ds.ExtensionId,
					DesiredStateId: ds.Meta.DesiredStateId,
					Status:         node.Status.String(),
				})
			}
		}
	}

	err = json.NewEncoder(writer).Encode(&result)
	if err != nil {
		return false, errors.Wrap(err, "failed to encode")
	}

	return s.handleReturn(resp.Summary)
}

var getDesiredState = &cobra.Command{
	Use:   "get-desired-state",
	Short: "Get the desired state convergence summary for a service",
	Long: `Get the desired state convergence summary for a service.

pvnctl services --app <app> get-desired-state <service>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ctx := context.Background()
		s := &desiredStateStatus{service: name}

		outputFormat, err := cmd.Flags().GetString("format")
		if err != nil {
			log.Fatal(err)
		}

		switch outputFormat {
		case "json":
			_, err = s.renderDesiredStateJSON(ctx, os.Stdout)
		default:
			_, err = s.renderDesiredState(ctx, os.Stdout)
		}

		if err != nil {
			if !go_errors.Is(err, errDesiredStateConvergenceFailed) {
				cmd.PrintErrf("%s\n", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(getDesiredState)
}
