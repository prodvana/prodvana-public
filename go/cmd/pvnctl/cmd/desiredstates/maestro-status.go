package desiredstates

import (
	"context"
	"fmt"
	"io"
	"slices"
	"sort"
	"strings"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/pkg/sets"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	app_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/application"
	ds "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	maestro_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/maestro"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"
	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release_channel"
	"github.com/spf13/cobra"
)

func sortReleaseChannelsbyDependencies(ctx context.Context, app string, releaseChannels []string) ([]string, error) {
	resp, err := cmdutil.GetApplicationManagerClient().GetApplicationConfig(ctx, &app_pb.GetApplicationConfigReq{
		Application: app,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get application config for %s", app)
	}
	rcSet := sets.FromSlice(releaseChannels)
	rcDeps := map[string][]string{}
	for _, rc := range resp.CompiledConfig.ReleaseChannels {
		if !rcSet.Contains(rc.Name) {
			continue
		}
		for _, dep := range rc.Preconditions {
			switch inner := dep.Precondition.(type) {
			case *release_channel.Precondition_ReleaseChannelStable_:
				if rcSet.Contains(inner.ReleaseChannelStable.GetReleaseChannel()) {
					rcDeps[rc.Name] = append(rcDeps[rc.Name], inner.ReleaseChannelStable.GetReleaseChannel())
				}
			}
		}
	}
	sorted := make([]string, 0, len(releaseChannels))
	toProcess := slices.Clone(releaseChannels)
	for len(toProcess) > 0 {
		toAdd := sets.NewSet[string]()
		for _, rc := range toProcess {
			if len(rcDeps[rc]) == 0 {
				toAdd.Add(rc)
			}
		}
		toAddSorted := toAdd.AsSlice()
		if len(toAddSorted) == 0 {
			return nil, errors.Errorf("internal error: no progress made trying to sort release channels %+v", releaseChannels)
		}
		sort.Strings(toAddSorted)
		sorted = append(sorted, toAddSorted...)
		toProcess = slices.DeleteFunc(toProcess, toAdd.Contains)
		for rc, deps := range rcDeps {
			rcDeps[rc] = slices.DeleteFunc(deps, toAdd.Contains)
		}
	}
	return sorted, nil
}

var maestroStatusCmd = &cobra.Command{
	Use:    "maestro-status <entity type> <entity id>",
	Short:  "Show status of pending Maestro releases.",
	Hidden: true,
	Args:   cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		entityTypeStr := args[0]
		entityTypeInt, ok := ds_model_pb.Type_value[strings.ToUpper(entityTypeStr)]
		if !ok {
			return fmt.Errorf("unknown entity type %s", entityTypeStr)
		}
		entityType := ds_model_pb.Type(entityTypeInt)
		switch entityType {
		case ds_model_pb.Type_SERVICE:
		default:
			return fmt.Errorf("unsupported entity type %s", entityTypeStr)
		}
		entityId := args[1]
		ctx := context.Background()
		var firstReleaseId string
		err := cmdutil.RenderWatch(ctx, func(ctx context.Context, writer io.Writer) (bool, error) {
			var releases []*maestro_pb.MaestroRelease
			var nextToken string
			for {
				resp, err := cmdutil.GetDesiredStateManagerClient().ListMaestroReleases(ctx, &ds.ListMaestroReleasesReq{
					EntityId: &ds_model_pb.Identifier{
						Type: entityType,
						Name: entityId,
					},
					StartingReleaseId: firstReleaseId,
					PageToken:         nextToken,
					PageSize:          100,
				})
				if err != nil {
					return false, err
				}
				nextToken = resp.NextPageToken
				releases = append(releases, resp.MaestroReleases...)
				if resp.NextPageToken == "" {
					break
				}
			}
			if len(releases) == 0 {
				fmt.Fprintf(writer, "No pending releases found for %s %s\n", entityType, entityId)
				return false, nil
			}
			firstReleaseId = releases[0].Meta.Id // avoid the table shifting around after first release is rendered
			var releaseChannels []string
			var app string
			{
				releaseChannelsSet := sets.NewSet[string]()
				for _, release := range releases {
					for _, inst := range release.Config.CompiledDesiredState.GetService().ReleaseChannels {
						releaseChannelsSet.Add(inst.ReleaseChannel)
						if app == "" {
							app = inst.Application
						}
						if app != inst.Application {
							return false, fmt.Errorf("multiple applications found: %s and %s", app, inst.Application)
						}
					}
				}
				releaseChannels = releaseChannelsSet.AsSlice()
				var err error
				releaseChannels, err = sortReleaseChannelsbyDependencies(ctx, app, releaseChannels)
				if err != nil {
					return false, err
				}
			}
			releaseChannelIdx := map[string]int{}
			for idx, rc := range releaseChannels {
				releaseChannelIdx[rc] = idx
			}
			outputListing := cmdutil.NewOutputListing(append([]string{"RELEASE"}, releaseChannels...)...)
			for _, release := range releases {
				row := make([]interface{}, len(releaseChannels)+1)
				row[0] = fmt.Sprintf("%s at %s", release.Meta.Id, release.Config.CreationTimestamp.AsTime())
				for _, inst := range release.Config.CompiledDesiredState.GetService().ReleaseChannels {
					idx := releaseChannelIdx[inst.ReleaseChannel]
					version := inst.Versions[0].Version
					instStatus := release.State.ReleaseChannelStatuses[inst.ReleaseChannelId]
					var statusString string
					switch instStatus {
					case maestro_pb.Status_SUCCEEDED:
						statusString = fmt.Sprintf("%s ✓", version)
					case maestro_pb.Status_FAILED:
						statusString = fmt.Sprintf("%s ✖", version)
					case maestro_pb.Status_IN_PROGRESS:
						statusString = fmt.Sprintf("%s →", version)
					default:
						statusString = version
					}
					switch instStatus {
					case maestro_pb.Status_SUCCEEDED:
						statusString = color.New(color.FgGreen).Sprint(statusString)
					case maestro_pb.Status_FAILED:
						statusString = color.New(color.FgRed).Sprint(statusString)
					case maestro_pb.Status_IN_PROGRESS:
						statusString = color.New(color.FgBlue).Sprint(statusString)
					}
					row[idx+1] = statusString
				}
				outputListing.AddRow(row...)
			}
			return false, outputListing.WriteTable(writer)
		})
		return err
	},
}

func init() {
	RootCmd.AddCommand(maestroStatusCmd)
}
