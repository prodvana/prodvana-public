package recipes

import (
	"context"
	"fmt"
	"log"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"
	"google.golang.org/protobuf/encoding/protojson"

	recipe_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/recipe"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get recipe config",
	Long: `Get a recipe config

pvnctl recipes get recipe-name
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		recipeName := args[0]
		recipeCfgResp, err := cmdutil.GetRecipeManagerClient().GetRecipeConfig(ctx, &recipe_pb.GetRecipeConfigReq{
			Recipe: recipeName,
		})
		if err != nil {
			log.Fatal(err)
		}

		switch cmdutil.OutputFormat {
		case "json":
			fmt.Printf("%s\n", protojson.Format(recipeCfgResp))
		default:
			fmt.Printf("Version: %s\n%s\n", recipeCfgResp.Version, recipeCfgResp.InputConfig)
		}
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
