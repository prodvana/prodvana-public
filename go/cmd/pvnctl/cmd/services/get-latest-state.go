package services

import (
	"context"
	"encoding/json"
	go_errors "errors"
	"fmt"
	"io"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func getProtections(ds *ds_model_pb.ServiceInstanceState, entityGraph *ds_model_pb.EntityGraph) ([]*ds_model_pb.Entity, error) {
	if ds == nil {
		return nil, nil
	}
	var rootEntity *ds_model_pb.Entity
	for _, entity := range entityGraph.Entities {
		if entity.Id.Type == ds_model_pb.Type_SERVICE_INSTANCE && entity.DesiredState.GetServiceInstance().ReleaseChannelId == ds.ReleaseChannelId {
			rootEntity = entity
			break
		}
	}
	if rootEntity == nil {
		return nil, errors.Errorf("Could not find root entity for %s in entity graph", ds.ReleaseChannel)
	}
	depsSet := map[string]struct{}{}
	for _, dep := range rootEntity.Dependencies {
		if dep.Type == ds_model_pb.Type_PROTECTION_LINK {
			depsSet[dep.Name] = struct{}{}
		}
	}
	var protections []*ds_model_pb.Entity
	for _, entity := range entityGraph.Entities {
		if _, ok := depsSet[entity.Id.Name]; ok {
			protections = append(protections, entity)
		}
	}
	return protections, nil
}

func (s *desiredStateStatus) renderLatestState(ctx context.Context, writer io.Writer, includeProtections bool) error {
	resp, err := cmdutil.GetDesiredStateManagerClient().GetServiceDesiredStateConvergenceSummary(ctx, &ds_pb.GetServiceDesiredStateConvergenceSummaryReq{
		Service:                s.service,
		Application:            cmdutil.App,
		FastNoDeprecatedFields: true,
	})
	if err != nil {
		return errors.Wrap(err, "Failed to get desired state for service")
	}

	// LastSeenState is not guaranteed to have metadata
	status := resp.Summary.Statuses[resp.Summary.DesiredState.GetService().Meta.DesiredStateId]
	fmt.Fprintf(writer, "Desired state: %s\n", resp.Summary.DesiredState.GetService().Meta.RootDesiredStateId)
	fmt.Fprintf(writer, "Status: %s\n", status)
	fmt.Fprintf(writer, "Release channels:\n")

	siNodes := getServiceInstanceNodesSorted(resp.Summary.EntityGraph)

	for _, node := range siNodes {
		ds := node.DesiredState.GetServiceInstance()
		fmt.Fprintln(writer, "")
		w := tabwriter.NewWriter(writer, 0, 2, 2, ' ', 0)
		fmt.Fprintf(w, "Name:\t%s\n", ds.ReleaseChannel)
		fmt.Fprintf(w, "Status:\t%s\n", node.Status.String())
		lastSeen := node.LastSeenState.GetServiceInstance()
		if lastSeen.GetRollbackVersion() != nil {
			fmt.Fprintf(w, "Rollback Version:\t%s\n", lastSeen.GetRollbackVersion().Version)
		}
		fmt.Fprintf(w, "Rollback:\t%t\n", lastSeen.Rollback)
		w.Flush()
		if len(lastSeen.GetVersions()) > 0 {
			fmt.Fprintln(writer, "Versions:")
			for _, version := range lastSeen.Versions {
				fmt.Fprintf(w, "  %s:\t%d\n", version.Version, version.Replicas)
			}
			w.Flush()
		}
		if lastSeen.GetDelivery() != nil {
			fmt.Fprintln(writer, "Delivery:")
			fmt.Fprintf(w, "  Status:\t%s\n", lastSeen.Delivery.Status.String())
			fmt.Fprintf(w, "  Desired State:\t%s\n", lastSeen.Delivery.DesiredStateId)
			w.Flush()

			fmt.Fprintln(writer, "  Progress:")
			for idx, prog := range lastSeen.Delivery.CanaryProgress {
				fmt.Fprintf(w, "    %d:\t%d%%\t%s\t%s\n", idx, prog.CanaryWeight, prog.Duration.AsDuration(), prog.Status.String())
			}
			w.Flush()
		}
		if includeProtections {
			protections, err := getProtections(lastSeen, resp.Summary.EntityGraph)
			if err != nil {
				return err
			}
			if len(protections) > 0 {
				fmt.Fprintln(w, "Protections:")
				for _, protection := range protections {
					fmt.Fprintf(w, "  Name:\t%s\n", protection.Id.Name)
					fmt.Fprintf(w, "  Id:\t%s\n", protection.DesiredStateId)
					fmt.Fprintf(w, "  Status:\t%s\n", protection.LastSeenState.GetProtectionLink().Status.String())
					fmt.Fprintln(w, "")
				}
				w.Flush()
			}
		}
	}

	return nil
}

func (s *desiredStateStatus) renderLatestStateJSON(ctx context.Context, writer io.Writer, includeProtections bool) error {
	resp, err := cmdutil.GetDesiredStateManagerClient().GetServiceDesiredStateConvergenceSummary(ctx, &ds_pb.GetServiceDesiredStateConvergenceSummaryReq{
		Service:                s.service,
		Application:            cmdutil.App,
		FastNoDeprecatedFields: true,
	})
	if err != nil {
		return errors.Wrap(err, "Failed to get desired state for service")
	}

	status := resp.Summary.Status

	type canaryProgress struct {
		Weight          int    `json:"weight"`
		DurationSeconds int    `json:"duration"`
		Status          string `json:"status"`
	}

	type delivery struct {
		Status         string           `json:"status"`
		DesiredState   string           `json:"desired_state"`
		CanaryProgress []canaryProgress `json:"canary_progress"`
	}

	type version struct {
		Version  string `json:"version"`
		Replicas int    `json:"replicas"`
	}

	type protection struct {
		Name   string `json:"name"`
		Id     string `json:"id"`
		Status string `json:"status"`
	}

	type rcResult struct {
		Name        string       `json:"name"`
		Status      string       `json:"status"`
		Rollback    bool         `json:"rollback"`
		Versions    []version    `json:"versions"`
		Delivery    delivery     `json:"delivery"`
		Protections []protection `json:"protections"`
	}

	var result struct {
		DesiredState    string              `json:"desired_state"`
		Status          string              `json:"status"`
		ReleaseChannels map[string]rcResult `json:"release_channels"`
	}

	result.DesiredState = resp.Summary.DesiredState.GetService().Meta.RootDesiredStateId
	result.Status = status.String()
	result.ReleaseChannels = map[string]rcResult{}

	// LastSeenState is not guaranteed to have metadata
	for _, node := range resp.Summary.EntityGraph.GetEntities() {
		switch node.GetId().GetType() {
		case ds_model_pb.Type_SERVICE_INSTANCE:
			lastSeen := node.LastSeenState.GetServiceInstance()
			ds := node.DesiredState.GetServiceInstance()
			rc := &rcResult{
				Name:     ds.ReleaseChannel,
				Status:   node.Status.String(),
				Rollback: lastSeen.GetRollback(),
			}
			if len(lastSeen.GetVersions()) > 0 {
				versions := make([]version, len(lastSeen.Versions))
				for idx, v := range lastSeen.Versions {
					versions[idx] = version{
						Version:  v.Version,
						Replicas: int(v.Replicas),
					}
				}
				rc.Versions = versions
			}
			if lastSeen.GetDelivery() != nil {
				delivery := delivery{
					Status:       lastSeen.Delivery.Status.String(),
					DesiredState: lastSeen.Delivery.DesiredStateId,
				}
				if len(lastSeen.Delivery.CanaryProgress) > 0 {
					progress := make([]canaryProgress, len(lastSeen.Delivery.CanaryProgress))
					for idx, prog := range lastSeen.Delivery.CanaryProgress {
						progress[idx] = canaryProgress{
							Weight:          int(prog.CanaryWeight),
							DurationSeconds: int(prog.Duration.AsDuration() * time.Second),
							Status:          prog.Status.String(),
						}
					}
					delivery.CanaryProgress = progress
				}
				rc.Delivery = delivery
			}
			if includeProtections {
				rawProtections, err := getProtections(lastSeen, resp.Summary.EntityGraph)
				if err != nil {
					return err
				}
				protections := make([]protection, len(rawProtections))
				for idx, prot := range rawProtections {
					protections[idx] = protection{
						Name:   prot.Id.Name,
						Id:     prot.DesiredStateId,
						Status: prot.LastSeenState.GetProtectionLink().Status.String(),
					}
				}
				rc.Protections = protections
			}
			result.ReleaseChannels[ds.ReleaseChannel] = *rc
		}
	}

	_ = json.NewEncoder(writer).Encode(&result)

	return nil
}

var getLatestState = &cobra.Command{
	Use:   "get-latest-state <service>",
	Short: "Get the latest state from the convergence summary for a service.",
	Long: `Get the latest state from the convergence summary for a service.

pvnctl services --app <app> get-latest-state <service>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		includeProtections, err := cmd.Flags().GetBool("protections")
		if err != nil {
			log.Fatal(err)
		}

		name := args[0]
		ctx := context.Background()
		s := &desiredStateStatus{service: name}

		outputFormat, err := cmd.Flags().GetString("format")
		if err != nil {
			log.Fatal(err)
		}

		switch outputFormat {
		case "json":
			err = s.renderLatestStateJSON(ctx, os.Stdout, includeProtections)
		default:
			err = s.renderLatestState(ctx, os.Stdout, includeProtections)
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
	RootCmd.AddCommand(getLatestState)
	getLatestState.Flags().Bool("protections", false, "Show protection details")
}
