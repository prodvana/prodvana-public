package runtimes

import (
	"context"
	"fmt"
	"log"

	"prodvana/cmd/cmdutil"

	"github.com/prodvana/prodvana-public/go/pkg/protoyaml"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environment"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: "Detect configuration of a runtime.",
	Long: `Detect configuration of a Runtime.

pvnctl runtimes detect [FLAGS]
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		runtime := args[0]
		resp, err := cmdutil.GetEnvironmentManagerClient().DetectClusterConfig(ctx, &environment.DetectClusterConfigReq{
			RuntimeName: runtime,
		})
		if err != nil {
			log.Fatal(err)
		}

		var outBytes []byte
		switch cmdutil.OutputFormat {
		case "json":
			outBytes, err = protojson.Marshal(resp.Config)
		default:
			outBytes, err = protoyaml.Marshal(resp.Config)
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", outBytes)
	},
}

func init() {
	RootCmd.AddCommand(detectCmd)
}
