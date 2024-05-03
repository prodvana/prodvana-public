package services

import (
	"context"
	"log"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a service",
	Long: `Delete a service everywhere it is deployed to.

pvnctl services --app=<app> delete <service>
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
		_, err = cmdutil.GetServiceManagerClient().DeleteService(ctx, &service_pb.DeleteServiceReq{
			Service:     name,
			Application: cmdutil.App,
		})
		if err != nil {
			cmd.PrintErrf("Failed to delete services: %+v\n", err)
			os.Exit(1)
		}
		cmd.Printf("Deleted %s\n", name)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Bool("dry-run", true, "Do a dry run instead of actual deletion.")
}
