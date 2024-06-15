package recipes

import (
	"context"
	"log"
	"os"
	"sort"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	recipe_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/recipe"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all recipes",
	Long: `List all recipes"

pvnctl recipes list
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		service, err := cmd.Flags().GetString("service")
		if err != nil {
			log.Fatal(err)
		}

		app, err := cmd.Flags().GetString("app")
		if err != nil {
			log.Fatal(err)
		}

		if service == "" {
			log.Fatal("service is required")
		}

		if app == "" {
			log.Fatal("app is required")
		}

		recipesResp, err := cmdutil.GetRecipeManagerClient().ListRecipes(ctx, &recipe_pb.ListRecipesReq{
			Filter: &recipe_pb.ListRecipesReq_ServiceFilter{
				ServiceFilter: &recipe_pb.ListRecipesReq_ByService{
					Service:     service,
					Application: app,
				},
			},
		})
		if err != nil {
			log.Fatal(err)
		}
		recipes := recipesResp.Recipes
		sort.Slice(recipes, func(i, j int) bool {
			return recipes[i].Meta.Name < recipes[j].Meta.Name
		})
		listing := cmdutil.NewOutputListing("RECIPE ID", "SERVICE")
		for _, recipe := range recipes {
			listing.AddRow(recipe.Meta.Id, recipe.Meta.Name)
		}
		err = cmdutil.WriteStdout(listing)
		if err != nil {
			cmd.PrintErrf("Failed to list recipes: %+v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.Flags().String("service", "", "Service name to filter by")
	listCmd.Flags().String("app", "", "App name to filter by")
}
