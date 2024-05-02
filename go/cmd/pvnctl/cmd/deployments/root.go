package deployments

import (
	"prodvana/cmd/cmdutil"
	"prodvana/cmd/pvnctl/cmd/services/primitives"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "deployments <subcommand>",
	Aliases: []string{"release", "rel", "releases", "deployment", "deploy", "deploys"},
	Short:   "Operate on deployments.",
	Long: `Operate on deployments.

pvnctl deployments record --deployment-system=manual
`,
	PersistentPreRun: cmdutil.ValidateAndRefreshAuthPreRun,
}

func init() {
	RootCmd.AddCommand(primitives.RootCmd)
}
