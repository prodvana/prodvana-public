package integration

import (
	"prodvana/cmd/cmdutil"
	"prodvana/cmd/pvnctl/cmd/integration/grafana"
	"prodvana/cmd/pvnctl/cmd/integration/pagerduty"
	"prodvana/cmd/pvnctl/cmd/integration/registry"
	"prodvana/cmd/pvnctl/cmd/integration/slack"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "integration <subcommand>",
	Aliases: []string{"integrations"},
	Short:   "Manage Integrations",
	Long: `Manage Integrations

pvnctl integration registry add
pvnctl integration slack install
pvnctl integration pagerduty install
pvnctl integration grafana install
`,
	PersistentPreRun: cmdutil.ValidateAndRefreshAuthPreRun,
}

func init() {
	RootCmd.AddCommand(registry.RootCmd)
	RootCmd.AddCommand(slack.RootCmd)
	RootCmd.AddCommand(pagerduty.RootCmd)
	RootCmd.AddCommand(grafana.RootCmd)
}
