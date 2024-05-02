package organizations

import (
	"prodvana/cmd/cmdutil"
	"prodvana/cmd/pvnctl/cmd/services/primitives"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "organizations <subcommand>",
	Aliases: []string{"organization", "org", "orgs"},
	Short:   "Operate on organizations.",
	Long: `Operate on organizations.

pvnctl organizations edit-metadata
`,
	PersistentPreRun: cmdutil.ValidateAndRefreshAuthPreRun,
}

func init() {
	RootCmd.AddCommand(primitives.RootCmd)
}
