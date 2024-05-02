package protections

import (
	"context"
	"log"
	"os"

	"prodvana/cmd/cmdutil"

	protection_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/protection"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List protections",
	Long: `List protections.

pvnctl protections list
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		var nextPageToken string
		protections := []*protection_pb.Protection{}
		for {
			resp, err := cmdutil.GetProtectionManagerClient().ListProtections(ctx, &protection_pb.ListProtectionsReq{
				PageSize:  100,
				PageToken: nextPageToken,
			})
			if err != nil {
				log.Fatal(err)
			}
			protections = append(protections, resp.Protections...)
			if resp.NextPageToken == "" {
				break
			}
			nextPageToken = resp.NextPageToken
		}

		listing := cmdutil.NewOutputListing("PROTECTION", "SOURCE")

		for _, protection := range protections {
			source := "custom"
			if protection.Builtin {
				source = "builtin"
			}
			listing.AddRow(protection.Meta.Name, source)
		}
		err := cmdutil.WriteStdout(listing)
		if err != nil {
			cmd.PrintErrf("Failed to list protections: %+v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
