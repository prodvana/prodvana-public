package secrets

import (
	"context"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	secrets_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/secrets"

	"github.com/spf13/cobra"
)

var listVersionsCmd = &cobra.Command{
	Use:   "list-versions",
	Short: "List all versions of this secret.",
	Long: `List all versions of this secret.

pvnctl secrets list-version <secret name>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		secretKey := args[0]
		resp, err := cmdutil.GetSecretsManagerClient().ListSecretVersions(ctx, &secrets_pb.ListSecretVersionsReq{
			Key: secretKey,
		})
		if err != nil {
			cmd.PrintErrf("Failed to list secret versions: %+v\n", err)
			os.Exit(1)
		}
		listing := cmdutil.NewOutputListing("Version")
		for _, version := range resp.Versions {
			listing.AddRow(version)
		}
		err = cmdutil.WriteStdout(listing)
		if err != nil {
			cmd.PrintErrf("Failed to list secret versions: %+v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(listVersionsCmd)
}
