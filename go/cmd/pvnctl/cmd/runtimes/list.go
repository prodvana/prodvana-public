package runtimes

import (
	"context"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environment"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list <name>",
	Short: "List runtimes connected to Prodvana.",
	Long: `List runtimes connected to Prodvana.

pvnctl runtimes list
`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		resp, err := cmdutil.GetEnvironmentManagerClient().ListClusters(ctx, &environment.ListClustersReq{})
		if err != nil {
			cmd.PrintErrf("Failed to list clusters: %+v\n", err)
			os.Exit(1)
		}
		listing := cmdutil.NewOutputListing("ID", "NAME", "HEARTBEAT")
		for _, cluster := range resp.Clusters {
			hbeat := "N/A"
			if cluster.Type == environment.ClusterType_K8S {
				if cluster.LastHeartbeatTimestamp.IsValid() {
					hbeat = cluster.LastHeartbeatTimestamp.AsTime().String()
				} else {
					hbeat = "Never"
				}
			}
			listing.AddRow(cluster.Id, cluster.Name, hbeat)
		}
		err = cmdutil.WriteStdout(listing)
		if err != nil {
			cmd.PrintErrf("Failed to list clusters: %+v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
