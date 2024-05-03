package preflight

import (
	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "preflight <subcommand>",
	Short: "Commands to validate your cloud infrastructure.",
	Long: `Commands to validate your cloud infrastructure.

pvnctl preflight k8s
pvnctl preflight ecs
`,
	PersistentPreRun: cmdutil.ValidateAndRefreshAuthPreRun,
}
