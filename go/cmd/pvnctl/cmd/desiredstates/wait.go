package desiredstates

import (
	"context"
	_ "embed"
	"encoding/json"
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
	"github.com/prodvana/prodvana-public/go/prodvana-sdk/desiredstates"
	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

var (
	errDesiredStateConvergenceFailed = fmt.Errorf("Failed to converge")
)

type desiredStateStatus struct {
	dsId              string
	exitOnExactStatus ds_model_pb.Status
}

func mustGetMeta(state *ds_model_pb.State) *ds_model_pb.Metadata {
	meta, err := desiredstates.GetMeta(state)
	if err != nil {
		panic(err)
	}
	return meta
}

func findEntity(graph *ds_model_pb.EntityGraph, dsId string) (*ds_model_pb.Entity, error) {
	for _, entity := range graph.Entities {
		if entity.DesiredStateId == dsId {
			return entity, nil
		}
	}
	return nil, errors.Errorf("%s not found", dsId)
}

func findEntityByEntityId(graph *ds_model_pb.EntityGraph, id *ds_model_pb.Identifier) (*ds_model_pb.Entity, error) {
	for _, entity := range graph.Entities {
		if proto.Equal(entity.Id, id) {
			return entity, nil
		}
	}
	return nil, errors.Errorf("%s not found", id)
}

func (s *desiredStateStatus) renderDesiredState(ctx context.Context, writer io.Writer) (bool, error) {
	resp, err := cmdutil.GetDesiredStateManagerClient().GetDesiredStateGraph(ctx, &ds_pb.GetDesiredStateGraphReq{
		Query: &ds_pb.GetDesiredStateGraphReq_DesiredStateId{
			DesiredStateId: s.dsId,
		},
		Depth: -1,
	})
	if err != nil {
		return false, errors.Wrap(err, "Failed to get desired state for service")
	}

	trackingDesiredStateId := s.dsId

	if resp.PendingSetDesiredState != nil {
		fmt.Fprintf(writer, "Desired state ID: %s\n", trackingDesiredStateId)
		color.New(color.FgYellow).Fprintf(writer, "Status: PENDING\n")
		return false, nil
	}

	entity, err := findEntity(resp.EntityGraph, trackingDesiredStateId)
	if err != nil {
		return false, err
	}
	fmt.Fprintf(writer, "Desired state ID: %s\n", trackingDesiredStateId)
	fmt.Fprintf(writer, "Entity ID: %s\n", entity.Id)
	cmdutil.ColorForStatus(entity.Status).Fprintf(writer, "Status: %s\n", entity.Status)
	if len(entity.StatusExplanations) > 0 {
		fmt.Fprintf(writer, "Status explanations: %s\n", entity.StatusExplanations)
	}

	if len(entity.Dependencies) > 0 {
		fmt.Fprintf(writer, "Child entities:\n")

		for _, child := range entity.Dependencies {
			childEntity, err := findEntityByEntityId(resp.EntityGraph, child)
			if err != nil {
				return false, err
			}
			fmt.Fprintf(writer, "• Entity ID: %s\n", childEntity.Id)
			cmdutil.ColorForStatus(childEntity.Status).Fprintf(writer, "  Status: %s\n", childEntity.Status)
			if len(childEntity.StatusExplanations) > 0 {
				fmt.Fprintf(writer, "  Status explanations: %s\n", childEntity.StatusExplanations)
			}
		}
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

func (s *desiredStateStatus) renderDesiredStateJSON(ctx context.Context, writer io.Writer) (bool, error) {
	resp, err := cmdutil.GetDesiredStateManagerClient().GetDesiredStateConvergenceSummary(ctx, &ds_pb.GetDesiredStateConvergenceReq{
		DesiredStateId:         s.dsId,
		FastNoDeprecatedFields: true,
	})
	if err != nil {
		return false, errors.Wrap(err, "Failed to get desired state for service")
	}

	trackingDesiredStateId := s.dsId
	if trackingDesiredStateId == "" {
		trackingDesiredStateId = mustGetMeta(resp.Summary.DesiredState).RootDesiredStateId
	}
	entity, err := findEntity(resp.Summary.EntityGraph, trackingDesiredStateId)
	if err != nil {
		return false, err
	}

	type entityId struct {
		Type string `json:"type"`
		Id   string `json:"id"`
	}

	type missingApproval struct {
		Topic          string `json:"topic"`
		DesiredStateId string `json:"desired_state_id"`
	}

	type entityResult struct {
		EntityId        entityId          `json:"entity_id"`
		DesiredState    string            `json:"desired_state"`
		MissingApproval *missingApproval  `json:"missing_approval"`
		Status          string            `json:"status"`
		Annotations     map[string]string `json:"annotations"`
	}

	var result struct {
		DesiredState string         `json:"desired_state"`
		Status       string         `json:"status"`
		Children     []entityResult `json:"children"`
	}

	status := entity.Status

	result.DesiredState = entity.RootDesiredStateId
	result.Status = status.String()

	for _, child := range entity.Dependencies {
		childEntity, err := findEntityByEntityId(resp.Summary.EntityGraph, child)
		if err != nil {
			return false, err
		}
		annotations := make(map[string]string)
		for _, annotation := range childEntity.Annotations.Annotations {
			annotations[annotation.Key] = annotation.Value
		}
		res := entityResult{
			DesiredState: childEntity.DesiredStateId,
			Status:       childEntity.Status.String(),
			EntityId: entityId{
				Type: child.Type.String(),
				Id:   child.Name,
			},
			Annotations: annotations,
		}
		if childEntity.MissingApproval != nil {
			res.MissingApproval = &missingApproval{
				Topic:          childEntity.MissingApproval.Topic,
				DesiredStateId: childEntity.MissingApproval.DesiredStateId,
			}
		}
		result.Children = append(result.Children, res)
	}

	_ = json.NewEncoder(writer).Encode(&result)

	switch status {
	case ds_model_pb.Status_CONVERGED:
		return true, nil
	case ds_model_pb.Status_FAILED:
		return false, errors.Wrap(errDesiredStateConvergenceFailed, "Desired state failed to converge")
	case ds_model_pb.Status_ROLLED_BACK:
		return false, errors.Wrap(errDesiredStateConvergenceFailed, "Desired state rolled back")
	default:
		return false, nil
	}
}

var waitCmd = &cobra.Command{
	Use:   "wait",
	Short: "Wait for desired state to converge or fail.",
	Long: `Wait for a desired state to either converge or fail to converge. Exit 0 on convergence success, 3 if timed out, 2 if failed, 1 otherwise.

pvnctl desired-states wait <desired state id>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dsId := args[0]
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
		s := &desiredStateStatus{dsId: dsId, exitOnExactStatus: requiredStatus}
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
