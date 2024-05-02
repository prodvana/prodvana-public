package secrets

import (
	"context"
	"os"

	"prodvana/cmd/cmdutil"

	secrets_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/secrets"

	"github.com/spf13/cobra"
)

var deleteVersionCmd = &cobra.Command{
	Use:   "delete-version",
	Short: "Delete a secret version.",
	Long: `Delete a secret version.

pvnctl secrets delete-version <key> <version>
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		key := args[0]
		version := args[1]
		_, err := cmdutil.GetSecretsManagerClient().DeleteSecretVersion(ctx, &secrets_pb.DeleteSecretVersionReq{
			Key:     key,
			Version: version,
		})
		if err != nil {
			cmd.PrintErrf("Failed to delete secret version: %+v\n", err)
			os.Exit(1)
		}

		cmd.Println("Secret version deleted.")
	},
}

func init() {
	RootCmd.AddCommand(deleteVersionCmd)
}
