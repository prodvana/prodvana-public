package secrets

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	secrets_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/secrets"
	"golang.org/x/term"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a secret value.",
	Long: `Set a secret value.

pvnctl secrets set <key> <value>
cat secret | pvnctl secrets set <key>

`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		key := args[0]

		var value string
		if len(args) == 2 {
			value = args[1]
		} else if !term.IsTerminal(int(os.Stdin.Fd())) {
			b, err := io.ReadAll(os.Stdin)
			if err != nil {
				cmd.PrintErrf("Failed reading value from stdin: %+v\n", err)
				os.Exit(1)
			}
			value = strings.TrimSuffix(string(b), "\n")
		} else {
			cmd.PrintErrf("No value passed, please pass a value as an argument or from stdin\n")
			os.Exit(1)
		}

		resp, err := cmdutil.GetSecretsManagerClient().SetSecret(ctx, &secrets_pb.SetSecretReq{
			Key:   key,
			Value: value,
		})
		if err != nil {
			cmd.PrintErrf("Failed to set secret: %+v\n", err)
			os.Exit(1)
		}

		cmd.Printf("Created secret version: %s\n", resp.Version)
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
