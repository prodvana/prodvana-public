package apitokens

import (
	"context"
	"fmt"
	"log"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	auth_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/auth"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: "delete",
	Short: `Delete an API Token.

	pvnctl api-tokens delete <name>
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ctx := context.Background()
		_, err := cmdutil.GetApiTokenManagerClient().DeleteOrgApiToken(ctx, &auth_pb.DeleteOrgApiTokenReq{
			Name: name,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Deleted API token %s\n", name)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
