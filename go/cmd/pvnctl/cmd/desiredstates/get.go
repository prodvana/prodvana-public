package desiredstates

import (
	"context"
	_ "embed"
	go_errors "errors"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the desired state convergence summary.",
	Long: `Get the desired state convergence summary.

pvnctl desired-states get <desired state id>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dsId := args[0]
		ctx := context.Background()
		s := &desiredStateStatus{dsId: dsId}

		outputFormat, err := cmd.Flags().GetString("format")
		if err != nil {
			log.Fatal(err)
		}

		switch outputFormat {
		case "json":
			_, err = s.renderDesiredStateJSON(ctx, os.Stdout)
		default:
			_, err = s.renderDesiredState(ctx, os.Stdout)
		}

		if err != nil {
			if !go_errors.Is(err, errDesiredStateConvergenceFailed) {
				cmd.PrintErrf("%s\n", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
