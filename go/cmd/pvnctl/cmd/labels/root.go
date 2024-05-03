package labels

import (
	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "labels <subcommand>",
	Short:   "Query objects in Prodvana using labels",
	Aliases: []string{"label"},
	Long: `Query objects in prodvana using labels.

pvnctl labels query "@type=runtime"
pvnctl labels get runtime my-runtime
`,
	PersistentPreRun: cmdutil.ValidateAndRefreshAuthPreRun,
}

func init() {
}
