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
		recipesResp, err := cmdutil.GetRecipeManagerClient().ListRecipes(ctx, &recipe_pb.ListRecipesReq{})
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
}
