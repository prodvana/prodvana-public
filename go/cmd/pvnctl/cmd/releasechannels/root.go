package releasechannels

import (
	"prodvana/cmd/cmdutil"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "release-channels <subcommand>",
	Short: "Operate on release channels.",
	Long: `Operate on release channels.

pvnctl release-channels configure <name>`,
	PersistentPreRun: cmdutil.ValidateAndRefreshAuthPreRun,
}

func init() {
	RootCmd.PersistentFlags().StringVar(&cmdutil.App, "app", "", "Application to manage services for.")
	cmdutil.Must(RootCmd.MarkPersistentFlagRequired("app"))
}
