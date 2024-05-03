package secrets

import (
	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

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
