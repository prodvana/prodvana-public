package primitives

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "primitives <subcommand>",
	Short: "Primitive commands for operating on a service. Can be useful for scripting, but most will not need this.",
	Long: `Primitive commands for operating on a service. Can be useful for scripting, but most will not need this.

pvnctl services --app=<app> primitives configure <name>`,
	Hidden: true,
}

func init() {
}
