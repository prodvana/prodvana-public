package blobs

import (
	"prodvana/cmd/cmdutil"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "blobs <subcommand>",
	Aliases: []string{"blobs"},
	Short:   "Operate on blobs stored on Prodvana.",
	Long: `Operate on blobs stored on Prodvana.

pvnctl blobs get`,
	PersistentPreRun: cmdutil.ValidateAndRefreshAuthPreRun,
}

func init() {
}
