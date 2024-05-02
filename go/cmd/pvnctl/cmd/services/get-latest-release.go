package services

import (
	"context"
	"fmt"
	"log"
	"prodvana/cmd/cmdutil"

	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	"github.com/spf13/cobra"
)

var getLatestRelease = &cobra.Command{
	Use:   "get-latest-release <service>",
	Short: "Get the latest release id for a service.",
	Long: `Get the latest release id for a service.

pvnctl services --app <app> get-latest-release <service>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ctx := context.Background()
		latest, err := cmdutil.GetDesiredStateManagerClient().GetServiceLatestCombinedReleaseDesiredState(ctx, &ds_pb.GetServiceLatestCombinedReleaseDesiredStateReq{
			Service:     name,
			Application: cmdutil.App,
		})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(latest.CompiledDesiredState.GetService().Meta.ReleaseId)
	},
}

func init() {
	RootCmd.AddCommand(getLatestRelease)
}
