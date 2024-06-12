package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/pkg/sets"

	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"
	events_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/events"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/spf13/cobra"
)

var eventLogsCmd = &cobra.Command{
	Use:    "event-logs",
	Short:  "Get event logs for a desired state",
	Hidden: true,
	Long: `Get event logs for a desired state.

pvnctl services --app <app> event-logs <desired-state-id>
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		seenEvents := sets.NewSet[string]()

		follow, err := cmd.Flags().GetBool("follow")
		if err != nil {
			log.Fatal(err)
		}

		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			log.Fatal(err)
		}

		var lookups []*events_pb.Lookup

		rootDesiredStateId, err := cmd.Flags().GetString("root-desired-state-id")
		if err != nil {
			log.Fatal(err)
		}
		desiredStateId, err := cmd.Flags().GetString("desired-state-id")
		if err != nil {
			log.Fatal(err)
		}
		releaseChannel, err := cmd.Flags().GetString("release-channel")
		if err != nil {
			log.Fatal(err)
		}
		limit, err := cmd.Flags().GetInt32("limit")
		if err != nil {
			log.Fatal(err)
		}
		includeDsIdDeps, err := cmd.Flags().GetBool("include-desired-state-dependencies")
		if err != nil {
			log.Fatal(err)
		}

		if rootDesiredStateId != "" {
			lookups = append(lookups, &events_pb.Lookup{
				LookupOneof: &events_pb.Lookup_RootDesiredStateId{
					RootDesiredStateId: rootDesiredStateId,
				},
			})
		} else if desiredStateId != "" {
			if !includeDsIdDeps {
				lookups = append(lookups, &events_pb.Lookup{
					LookupOneof: &events_pb.Lookup_DesiredStateId{
						DesiredStateId: desiredStateId,
					},
				})
			} else {
				// 2-step process to get useful desired states to get events for.
				// First find the entity name for the desired state ID.
				graph, err := cmdutil.GetDesiredStateManagerClient().GetDesiredStateGraph(ctx, &ds_pb.GetDesiredStateGraphReq{
					Query: &ds_pb.GetDesiredStateGraphReq_DesiredStateId{
						DesiredStateId: desiredStateId,
					},
				})
				if err != nil {
					log.Fatalf("Failed to get desired state graph: %s", err)
				}

				var entity *ds_model_pb.Entity
				for _, ent := range graph.EntityGraph.Entities {
					if ent.DesiredStateId == desiredStateId {
						entity = ent
						break
					}
				}

				// Then find everything that entity depends on, except other release channels.
				graph, err = cmdutil.GetDesiredStateManagerClient().GetDesiredStateGraph(ctx, &ds_pb.GetDesiredStateGraphReq{
					Query: &ds_pb.GetDesiredStateGraphReq_DesiredStateId{
						DesiredStateId: desiredStateId,
					},
					Depth:                 -1,
					ExcludeDependencyType: []ds_model_pb.DependencyType{ds_model_pb.DependencyType_DEPENDENCY_RELEASE_CHANNEL},
					RequiredEntityNames:   []string{entity.Id.Name},
				})
				if err != nil {
					log.Fatalf("Failed to get desired state graph: %s", err)
				}

				for _, ent := range graph.EntityGraph.Entities {
					lookups = append(lookups, &events_pb.Lookup{
						LookupOneof: &events_pb.Lookup_DesiredStateId{
							DesiredStateId: ent.DesiredStateId,
						},
					})
				}
			}
		} else if releaseChannel != "" {
			lookups = append(lookups, &events_pb.Lookup{
				LookupOneof: &events_pb.Lookup_ReleaseChannel{
					ReleaseChannel: &events_pb.ReleaseChannelLookup{
						Application:    cmdutil.App,
						ReleaseChannel: releaseChannel,
					},
				},
			})
		} else {
			log.Fatal("One of Root DSID or Desired state ID or release channel is required")
		}

		var afterTimestamp time.Time

		printEvent := func(event *events_pb.Event) {
			if seenEvents.Contains(event.Id) {
				return
			}

			seenEvents.Add(event.Id)
			if event.Timestamp.AsTime().After(afterTimestamp) {
				afterTimestamp = event.Timestamp.AsTime()
			}

			fmt.Printf("%s %s %+v\n", event.Timestamp.AsTime().Format(time.RFC3339), event.RelatedObjects.DesiredStateId, event.Title)
			if debug {
				fmt.Printf("%s\n", prototext.Format(event))
			}
		}

		// Find events
		for {
			toFetch := limit
			if toFetch > 100 || toFetch <= 0 {
				toFetch = 100
			}
			eventsResp, err := cmdutil.GetEventsManagerClient().GetEvents(ctx, &events_pb.GetEventsReq{
				Lookups:             lookups,
				OrderByAscTimestamp: true,
				UseOr:               true,
				PageSize:            toFetch,
				AfterTimestamp:      timestamppb.New(afterTimestamp),
			})
			if err != nil {
				log.Fatalf("Failed to get convergence details: %s", err)
			}
			for _, event := range eventsResp.Events {
				printEvent(event)
			}

			if limit <= 0 || follow {
				// No new events
				if len(eventsResp.Events) != 0 {
					continue
				}
			} else {
				limit -= int32(len(eventsResp.Events))
				if limit > 0 {
					continue
				}
			}

			if !follow {
				break
			}

			time.Sleep(5 * time.Second)
		}
	},
}

func init() {
	RootCmd.AddCommand(eventLogsCmd)
	eventLogsCmd.Flags().Bool("follow", false, "Stream event logs")
	eventLogsCmd.Flags().Int32("limit", -1, "Limit number of events to show (default: all)")
	eventLogsCmd.Flags().String("root-desired-state-id", "", "Root Desired state ID to get event logs for")
	eventLogsCmd.Flags().String("desired-state-id", "", "Desired state ID to get event logs for")
	eventLogsCmd.Flags().Bool("include-desired-state-dependencies", false, "Include dependencies of desired state id in event logs")
	eventLogsCmd.Flags().String("release-channel", "", "Release channel to get event logs for")
	eventLogsCmd.Flags().Bool("debug", false, "Print the entire event log proto (*very* verbose)")
}
