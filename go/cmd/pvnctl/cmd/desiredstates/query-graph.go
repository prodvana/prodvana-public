package desiredstates

import (
	"context"
	_ "embed"
	"log"
	"prodvana/cmd/cmdutil"
	"sort"

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
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dsId := args[0]
		ctx := context.Background()
		req := &ds_pb.GetDesiredStateGraphReq{
			Query: &ds_pb.GetDesiredStateGraphReq_DesiredStateId{DesiredStateId: dsId},
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
	queryGraphCmd.Flags().StringSlice("types", nil, "Types to query for")
	queryGraphCmd.Flags().StringSlice("required-entity-names", nil, "Required entity names")
	queryGraphCmd.Flags().Bool("include-parents", false, "Should include parents in the graph")
	queryGraphCmd.Flags().Int32("depth", 0, "Depth of the graph")
}
