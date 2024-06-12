package recipes

import (
	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "recipes <subcommand>",
	Short:   "Operate on recipes.",
	Aliases: []string{"recipe"},
	Long: `Operate on recipes.

pvnctl recipes <subcommand>`,
	PersistentPreRun: cmdutil.ValidateAndRefreshAuthPreRun,
}
