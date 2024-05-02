package runtimes

import (
	"prodvana/cmd/cmdutil"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "runtimes <subcommand>",
	Aliases: []string{"clusters", "runtime", "cluster"},
	Short:   "Operate on runtimes.",
	Long: `Operate on runtimes.

pvnctl runtimes add-k8s <name>`,
	PersistentPreRun: cmdutil.ValidateAndRefreshAuthPreRun,
}
