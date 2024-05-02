package registry

import (
	"prodvana/cmd/pvnctl/cmd/integration/registry/add"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "registry <subcommand>",
	Short: "Manage Container Registries",
	Long: `Manage Container Registries

pvnctl integration registry add
pvnctl integration registry connect
`,
}

func init() {
	RootCmd.AddCommand(add.RootCmd)
}
