package protections

import (
	"prodvana/cmd/cmdutil"
	"prodvana/cmd/pvnctl/cmd/services/primitives"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "protections <subcommand>",
	Aliases: []string{"protection"},
	Short:   "Manage protections.",
	Long: `Manage protections.

pvnctl protections create
pvnctl protections list
pvnctl protections edit <name>
`,
	PersistentPreRun: cmdutil.ValidateAndRefreshAuthPreRun,
}

func init() {
	RootCmd.AddCommand(primitives.RootCmd)
}
