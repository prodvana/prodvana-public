package applications

import (
	"context"
	"log"
	"os"

	"prodvana/cmd/cmdutil"

	application_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/application"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an application",
	Long: `Delete an application and resources (services, release channels, etc.) associated with it.

pvnctl applications delete <app>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		dryRun, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			log.Fatal(err)
		}
		if dryRun {
			cmd.PrintErrf("DRY-RUN: Would delete %s\n", name)
			return
		}
		ctx := context.Background()
		_, err = cmdutil.GetApplicationManagerClient().DeleteApplication(ctx, &application_pb.DeleteApplicationReq{
			Application: name,
		})
		if err != nil {
			cmd.PrintErrf("Failed to delete application: %+v\n", err)
			os.Exit(1)
		}
		cmd.Printf("Deleted %s\n", name)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Bool("dry-run", true, "Do a dry run instead of actual deletion.")
}
