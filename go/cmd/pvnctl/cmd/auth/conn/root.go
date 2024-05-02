package conn

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "context <subcommand>",
	Short: "Manage api server connection",
	Long: `Manage api server connections.

pvnctl auth context add name address
pvnctl auth context use name
`,
}
