package configs

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "configs <subcommand>",
	Aliases: []string{"config", "cfg", "cfgs"},
	Short:   "Apply config files",
	Long: `Apply config files.

pvnctl configs apply my-runtime/`,
}

func init() {
}
