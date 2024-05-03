package releasechannels

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	release_channel_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release_channel"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a release channel.",
	Long: `Get a release channels.

pvnctl release-channels --app=<app> get -f json <release-channel>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		releaseChannelName := args[0]
		outputFormat, err := cmd.Flags().GetString("format")
		if err != nil {
			log.Fatal(err)
		}
		resp, err := cmdutil.GetReleaseChannelManagerClient().GetReleaseChannel(ctx, &release_channel_pb.GetReleaseChannelReq{
			Application:    cmdutil.App,
			ReleaseChannel: releaseChannelName,
		})
		if err != nil {
			cmd.PrintErrf("Failed to get release channels: %+v\n", err)
			os.Exit(1)
		}
		switch outputFormat {
		case "json":
			jsonBytes, err := protojson.Marshal(resp.ReleaseChannel)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(os.Stdout, "%s\n", string(jsonBytes))
		default:
			log.Fatal("Only json format supported, pass --format json")
		}
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
