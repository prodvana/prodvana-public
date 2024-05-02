package releasechannels

import (
	"context"
	"log"
	"os"

	"prodvana/cmd/cmdutil"

	release_channel_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release_channel"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all release channels.",
	Long: `List all release channels.

pvnctl release-channels --app=<app> list
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		resp, err := cmdutil.GetReleaseChannelManagerClient().ListReleaseChannelsV2(ctx, &release_channel_pb.ListReleaseChannelsReq{
			Application: cmdutil.App,
		})
		if err != nil {
			cmd.PrintErrf("Failed to list release channels: %+v\n", err)
			os.Exit(1)
		}
		listing := cmdutil.NewOutputListing("ID", "RELEASE CHANNEL")

		for _, element := range resp.ReleaseChannels {
			listing.AddRow(element.Meta.Id, element.Meta.Name)
		}
		err = cmdutil.WriteStdout(listing)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
