package apitokens

import (
	"context"
	"log"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	auth_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/auth"

	"github.com/spf13/cobra"
)

const (
	DateTimeFormat = "2006-01-02 15:04:05"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List information about all organization api tokens.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		resp, err := cmdutil.GetApiTokenManagerClient().ListOrgApiTokens(ctx, &auth_pb.ListOrgApiTokensReq{})
		if err != nil {
			log.Fatal(err)
		}

		listing := cmdutil.NewOutputListing("NAME", "DESC", "EXPIRES", "CREATED")
		for _, token := range resp.Tokens {
			created := "N/A"
			if token.CreationTimestamp.IsValid() {
				created = token.CreationTimestamp.AsTime().Format(DateTimeFormat)
			}
			expires := "Never"
			if token.ExpiresTimestamp.IsValid() {
				expires = token.ExpiresTimestamp.AsTime().Format(DateTimeFormat)
			}
			listing.AddRow(token.Name, token.Description, expires, created)
		}
		err = cmdutil.WriteStdout(listing)
		if err != nil {
			cmd.PrintErrf("Failed to list api tokens: %+v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
