package primitives

import (
	"context"
	"fmt"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/pkg/protohelpers"

	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"
	version_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"

	"github.com/spf13/cobra"
)

var configureProtoCmd = &cobra.Command{
	Use:   "configure-proto-experimental <name>",
	Short: "Configure a service change sourced from a prototext file.",
	Long: `Configure a service change to be pushed later.

pvnctl services --app=<app> primitives configure-proto-experimental <config-path> [FLAGS] # will output version string
pvnctl services --app=<app> primitives push <name> --version=<version from configure command>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configPath := args[0]

		protoBytes, err := os.ReadFile(configPath)
		if err != nil {
			cmd.PrintErrf("Failed to read service proto file: %+v\n", err)
			os.Exit(1)
		}
		serviceConfig := &service_pb.ServiceConfig{}
		if err := protohelpers.Unmarshal(protohelpers.FileTypeInfer, configPath, protoBytes, serviceConfig, false); err != nil {
			cmd.PrintErrf("Failed to parse service proto file: %+v\n", err)
			os.Exit(1)
		}
		ctx := context.Background()
		resp, err := cmdutil.GetServiceManagerClient().ApplyParameters(
			ctx,
			&service_pb.ApplyParametersReq{
				Oneof: &service_pb.ApplyParametersReq_ServiceConfig{
					ServiceConfig: serviceConfig,
				},
				Application: cmdutil.App,
				Source:      version_pb.Source_INTERACTIVE_PVNCTL,
			},
		)
		if err != nil {
			cmd.PrintErrf("Failed to configure service: %+v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", resp.Version)
	},
}

func init() {
	RootCmd.AddCommand(configureProtoCmd)
}
