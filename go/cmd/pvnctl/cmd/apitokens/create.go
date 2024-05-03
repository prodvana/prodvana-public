package apitokens

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	auth_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/auth"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use: "create",
	Short: `Create a new api token tied to the organization. This token never expires and is not tied to the user in any way.

	pvnctl api-tokens create <name>
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		desc, err := cmd.Flags().GetString("desc")
		if err != nil {
			log.Fatal(err)
		}
		expires, err := cmd.Flags().GetInt("expires")
		if err != nil {
			log.Fatal(err)
		}
		name := args[0]

		var expireTime *timestamppb.Timestamp
		if expires > 0 {
			t := time.Now().Add(time.Duration(expires*24) * time.Hour)
			expireTime = timestamppb.New(t)
		}

		ctx := context.Background()
		resp, err := cmdutil.GetApiTokenManagerClient().CreateOrgApiToken2(ctx, &auth_pb.CreateOrgApiTokenReq{
			Name:             name,
			Description:      desc,
			ExpiresTimestamp: expireTime,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", resp.ApiToken)
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("desc", "d", "", "Optional, description for the API Token")
	createCmd.Flags().IntP("expires", "e", 0, "Optional, number of days the token will be valid for before expiring. The default (0) is non-expiring.")
}
