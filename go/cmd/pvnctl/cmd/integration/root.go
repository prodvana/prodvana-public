package integration

import (
	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/integration/grafana"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/integration/pagerduty"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/integration/registry"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/integration/slack"

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
