package applications

import (
	"prodvana/cmd/cmdutil"
	"prodvana/cmd/pvnctl/cmd/services/primitives"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "applications <subcommand>",
	Aliases: []string{"application", "app", "apps"},
	Short:   "Operate on applications.",
	Long: `Operate on applications.

pvnctl applications create
pvnctl applications list
pvnctl applications edit <name>
`,
	PersistentPreRun: cmdutil.ValidateAndRefreshAuthPreRun,
}

func init() {
	RootCmd.AddCommand(primitives.RootCmd)
}
