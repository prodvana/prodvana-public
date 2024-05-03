package services

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get-release-channel",
	Short: "Get a service instance.",
	Long: `Get a service instance.

pvnctl services --app=<app> get-release-channel -f json <service> <release-channel>
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		serviceName := args[0]
		releaseChannelName := args[1]
		outputFormat, err := cmd.Flags().GetString("format")
		if err != nil {
			log.Fatal(err)
		}
		resp, err := cmdutil.GetServiceManagerClient().GetServiceInstance(ctx, &service.GetServiceInstanceReq{
			Application:    cmdutil.App,
			ReleaseChannel: releaseChannelName,
			Service:        serviceName,
		})
		if err != nil {
			cmd.PrintErrf("Failed to get service instance: %+v\n", err)
			os.Exit(1)
		}
		switch outputFormat {
		case "json":
			jsonBytes, err := protojson.Marshal(resp.ServiceInstance)
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
