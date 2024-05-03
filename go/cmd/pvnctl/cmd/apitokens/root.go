package apitokens

import (
	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "api-tokens <subcommand>",
	Short:   "Manage api tokens",
	Aliases: []string{"api-token"},
	Long: `Manage api tokens.

pvnctl auth api-tokens create
`,
	PersistentPreRun: cmdutil.ValidateAndRefreshAuthPreRun,
}
