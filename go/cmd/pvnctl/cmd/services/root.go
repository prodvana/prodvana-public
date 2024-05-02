package services

import (
	"prodvana/cmd/cmdutil"
	"prodvana/cmd/pvnctl/cmd/services/primitives"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "services <subcommand>",
	Short:   "Operate on services.",
	Aliases: []string{"svc", "svcs", "service"},
	Long: `Operate on services.

pvnctl services --app=<app> create
pvnctl services --app=<app> list
pvnctl services --app=<app> push <name>
`,
	PersistentPreRun: cmdutil.ValidateAndRefreshAuthPreRun,
}

func init() {
	RootCmd.AddCommand(primitives.RootCmd)
	RootCmd.PersistentFlags().StringVar(&cmdutil.App, "app", "", "Application to manage services for.")
	cmdutil.Must(RootCmd.MarkPersistentFlagRequired("app"))
}
