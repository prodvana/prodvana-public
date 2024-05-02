package secrets

import (
	"prodvana/cmd/cmdutil"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "secrets",
	Aliases: []string{"secret"},
	Short:   "Operate on secrets.",
	Long: `Operate on secrets.

pvnctl secrets <subcommand>`,
	PersistentPreRun: cmdutil.ValidateAndRefreshAuthPreRun,
}
