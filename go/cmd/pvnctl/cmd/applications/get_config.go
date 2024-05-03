package applications

import (
	"context"
	"log"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/pkg/protoyaml"

	application_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/application"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

var getConfigCmd = &cobra.Command{
	Use:   "get-config",
	Short: "Get app config",
	Long: `Get app config.

pvnctl applications get-config <app name> [--version=<version>]
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		originalConfig, err := cmd.Flags().GetBool("original-config")
		if err != nil {
			log.Fatal(err)
		}
		version, err := cmd.Flags().GetString("version")
		if err != nil {
			log.Fatal(err)
		}
		outputFormat, err := cmd.Flags().GetString("output-format")
		if err != nil {
			log.Fatal(err)
		}

		app := args[0]

		resp, err := cmdutil.GetApplicationManagerClient().GetApplicationConfig(ctx, &application_pb.GetApplicationConfigReq{
			Application: app,
			Version:     version,
		})
		if err != nil {
			if status.Code(err) == codes.NotFound {
				log.Fatalf("Application %s not found\n", app)
			}
			log.Fatal(err)
		}
		config := resp.CompiledConfig
		if originalConfig {
			config = resp.Config
		}
		var outputBytes []byte
		switch outputFormat {
		case "yaml":
			outputBytes, err = protoyaml.Marshal(config)
		case "json":
			outputBytes, err = protojson.Marshal(config)
		case "pbtxt":
			outputBytes, err = prototext.Marshal(config)
		case "pb":
			outputBytes, err = proto.Marshal(config)
		default:
			log.Fatalf("Unknown output format %s\n", outputFormat)
		}
		if err != nil {
			log.Fatal(err)
		}
		_, err = cmd.OutOrStdout().Write(outputBytes)
		if err != nil {
			log.Fatal(err)
		}
		cmd.Println()
	},
}

func init() {
	getConfigCmd.Flags().Bool("original-config", false, "Output the original config supplied when creating/editing the app, without any defaults/post-processing.")
	getConfigCmd.Flags().String("version", "", "Config version to get, defaults to latest.")
	getConfigCmd.Flags().String("output-format", "yaml", "Output format. Supported values: yaml, json, pbtxt, pb")
	RootCmd.AddCommand(getConfigCmd)
}
