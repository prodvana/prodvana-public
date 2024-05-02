package grafana

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "grafana <subcommand>",
	Short: "Manage Grafana Integration",
	Long: `Manage Grafana Integration

pvnctl integration grafana install
`,
}
