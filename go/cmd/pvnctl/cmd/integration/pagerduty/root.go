package pagerduty

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "pagerduty <subcommand>",
	Short: "Manage PagerDuty Integration",
	Long: `Manage PagerDuty Integration

pvnctl integration pagerduty install
`,
}
