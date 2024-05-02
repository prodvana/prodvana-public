package slack

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "slack <subcommand>",
	Short: "Manage Slack Integration",
	Long: `Manage Slack Integration

pvnctl integration slack install
`,
}
