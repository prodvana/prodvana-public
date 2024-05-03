package releasechannels

import (
	"context"
	"log"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	release_channel_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release_channel"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "Delete a release channel",
	Long: `Delete a release channel.


pvnctl applications delete --app=<app> <name>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dryRun, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			log.Fatal(err)
		}
		rcName := args[0]
		if dryRun {
			cmd.PrintErrf("DRY-RUN: Would delete %s\n", rcName)
			return
		}
		ctx := context.Background()
		resp, err := cmdutil.GetReleaseChannelManagerClient().DeleteReleaseChannel(ctx, &release_channel_pb.DeleteReleaseChannelReq{
			Application:    cmdutil.App,
			ReleaseChannel: rcName,
		})
		if err != nil {
			cmd.PrintErrf("Failed to delete release channel: %+v\n", err)
			os.Exit(1)
		}
		cmd.Printf("Deleted release channel %s\n", rcName)
		cmd.Printf("%s\n", resp.Version)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Bool("dry-run", true, "Do a dry run instead of actual deletion.")
}
