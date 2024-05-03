package desiredstates

import (
	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "desired-states <subcommand>",
	Hidden:  true,
	Aliases: []string{"ds", "desired-state"},
	Short:   "Operate on desired states.",
	Long: `Operate on desired states.

pvnctl desired-states configure`,
	PersistentPreRun: cmdutil.ValidateAndRefreshAuthPreRun,
}

func init() {
}
