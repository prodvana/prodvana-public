package services

import (
	"context"
	"fmt"
	"log"
	"prodvana/cmd/cmdutil"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/prototext"

	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"
)

var getMaterializedConfigsCmd = &cobra.Command{
	Use:   "get-materialized-configs <service>",
	Short: "Get full materialized configs for a service",
	Long: `Get full materialized configs for a service (optionally for a specific version).

pvnctl services --app <app> get-materialized-configs <service>
`,
	Args:   cobra.ExactArgs(1),
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ctx := context.Background()
		version, err := cmd.Flags().GetString("version")
		if err != nil {
			log.Fatal(err)
		}

		resp, err := cmdutil.GetServiceManagerClient().GetMaterializedConfig(ctx, &service_pb.GetMaterializedConfigReq{
			Service:     name,
			Application: cmdutil.App,
			Version:     version,
		})
		if err != nil {
			log.Fatalf("error getting materialized configs: %+v", err)
		}

		fmt.Println(prototext.Format(resp))
	},
}

func init() {
	RootCmd.AddCommand(getMaterializedConfigsCmd)
	getMaterializedConfigsCmd.Flags().String("version", "", "Version to get materialized configs for")
}
