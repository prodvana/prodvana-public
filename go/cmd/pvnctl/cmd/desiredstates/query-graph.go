package desiredstates

import (
	"context"
	_ "embed"
	"log"
	"sort"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"

	"github.com/spf13/cobra"
)

var queryGraphCmd = &cobra.Command{
	Use:    "query-graph",
	Hidden: true,
	Short:  "Query desired state graph",
	Long: `Query desired state graph.

pvnctl desired-states query-graph <desired state id>
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		req := &ds_pb.GetDesiredStateGraphReq{}

		dsId, err := cmd.Flags().GetString("dsid")
		cmdutil.Must(err)

		service, err := cmd.Flags().GetString("service")
		cmdutil.Must(err)

		app, err := cmd.Flags().GetString("app")
		cmdutil.Must(err)

		if dsId != "" {
			if service != "" {
				log.Fatal("Cannot set both -dsid and -service")
			}

			if app != "" {
				log.Fatal("Cannot set both -dsid and -app")
			}
		} else {
			if service == "" {
				log.Fatal("Must set one of (-dsid, -service)")
			}

			if app == "" {
				log.Fatal("Must set -app when setting -service")
			}
		}

		if dsId != "" {
			req.Query = &ds_pb.GetDesiredStateGraphReq_DesiredStateId{DesiredStateId: dsId}
		} else {
			req.Query = &ds_pb.GetDesiredStateGraphReq_QueryByService_{
				QueryByService: &ds_pb.GetDesiredStateGraphReq_QueryByService{
					Application: app,
					Service:     service,
				},
			}
		}

		types, err := cmd.Flags().GetStringSlice("types")
		if err != nil {
			log.Fatal(err)
		}
		for _, type_ := range types {
			req.Types = append(req.Types, ds_model_pb.Type(ds_model_pb.Type_value[type_]))
		}

		requiredEntityNames, err := cmd.Flags().GetStringSlice("required-entity-names")
		if err != nil {
			log.Fatal(err)
		}
		req.RequiredEntityNames = requiredEntityNames

		includeParents, err := cmd.Flags().GetBool("include-parents")
		if err != nil {
			log.Fatal(err)
		}
		req.IncludeParents = includeParents

		depth, err := cmd.Flags().GetInt32("depth")
		if err != nil {
			log.Fatal(err)
		}
		req.Depth = depth

		resp, err := cmdutil.GetDesiredStateManagerClient().GetDesiredStateGraph(ctx, req)
		if err != nil {
			log.Fatal(err)
		}

		if resp.PendingSetDesiredState != nil {
			log.Printf("Pending set desired state: %s", resp.PendingSetDesiredState.DesiredStateId)
			return
		}

		sort.Slice(resp.EntityGraph.Entities, func(i, j int) bool {
			entI := resp.EntityGraph.Entities[i]
			entJ := resp.EntityGraph.Entities[j]
			if entI.Id.Type != entJ.Id.Type {
				return entI.Id.Type < entJ.Id.Type
			}

			return entI.Id.Name < entJ.Id.Name
		})

		log.Printf("Found %d entities", len(resp.EntityGraph.Entities))
		for _, ent := range resp.EntityGraph.Entities {
			log.Printf("DesiredStateId: %s Entity: %s", ent.DesiredStateId, ent.Id)
		}
	},
}

func init() {
	RootCmd.AddCommand(queryGraphCmd)
	queryGraphCmd.Flags().String("dsid", "", "Desired state id to query for")
	queryGraphCmd.Flags().String("app", "", "Application to query for (must set service)")
	queryGraphCmd.Flags().String("service", "", "Service to query for (must set app)")
	queryGraphCmd.Flags().StringSlice("types", nil, "Types to query for")
	queryGraphCmd.Flags().StringSlice("required-entity-names", nil, "Required entity names")
	queryGraphCmd.Flags().Bool("include-parents", false, "Should include parents in the graph")
	queryGraphCmd.Flags().Int32("depth", 0, "Depth of the graph")
}
