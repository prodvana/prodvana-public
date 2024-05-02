package secrets

import (
	"context"
	"os"

	"prodvana/cmd/cmdutil"

	secrets_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/secrets"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete all versions of a secret.",
	Long: `Delete all versions of a secret.

pvnctl secrets delete <key>

`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		key := args[0]

		_, err := cmdutil.GetSecretsManagerClient().DeleteSecret(ctx, &secrets_pb.DeleteSecretReq{
			Key: key,
		})
		if err != nil {
			cmd.PrintErrf("Failed to delete secret: %+v\n", err)
			os.Exit(1)
		}

		cmd.Println("Secret deleted.")
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
