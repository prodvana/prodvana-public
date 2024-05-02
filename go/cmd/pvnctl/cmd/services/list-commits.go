package services

import (
	"context"
	"log"
	"os"
	"prodvana/cmd/cmdutil"

	"github.com/spf13/cobra"

	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"
)

var listCommits = &cobra.Command{
	Use:   "list-commits <service>",
	Short: "List commits for a service",
	Long: `List commits based on image annotations for a service.

pvnctl services --app <app> list-commits <service>
`,
	Args:   cobra.ExactArgs(1),
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ctx := context.Background()

		resp, err := cmdutil.GetServiceManagerClient().ListCommits(ctx, &service_pb.ListCommitsReq{
			Service:     name,
			Application: cmdutil.App,
		})
		if err != nil {
			log.Fatalf("error listing commits: %+v", err)
		}

		listing := cmdutil.NewOutputListing("Commit", "URL", "Author", "Message")
		for _, commit := range resp.Commits {
			msg := commit.Message
			if len(msg) > 50 {
				msg = msg[:46] + " ..."
			}

			listing.AddRow(commit.CommitId, commit.Url, commit.Author.Name, msg)
		}
		err = cmdutil.WriteStdout(listing)
		if err != nil {
			cmd.PrintErrf("Failed to list commits: %+v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCommits)
}
