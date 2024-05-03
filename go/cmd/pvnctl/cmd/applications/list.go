package applications

import (
	"context"
	"log"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/application"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List applications",
	Long: `List applications.

pvnctl applications list
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		resp, err := cmdutil.GetApplicationManagerClient().ListApplications(ctx, &application.ListApplicationsReq{})
		if err != nil {
			log.Fatal(err)
		}
		listing := cmdutil.NewOutputListing("APPLICATION")
		for _, app := range resp.Applications {
			listing.AddRow(app.Meta.Name)
		}
		err = cmdutil.WriteStdout(listing)
		if err != nil {
			cmd.PrintErrf("Failed to list apps: %+v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
