package secrets

import (
	"context"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	secrets_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/secrets"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list <name>",
	Short: "List secrets defined in Prodvana.",
	Long: `List secrets defined in Prodvana.

pvnctl secrets list
`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		resp, err := cmdutil.GetSecretsManagerClient().ListSecrets(ctx, &secrets_pb.ListSecretsReq{})
		if err != nil {
			cmd.PrintErrf("Failed to list secrets: %+v\n", err)
			os.Exit(1)
		}
		listing := cmdutil.NewOutputListing("KEY", "LATEST VERSION")
		for _, secret := range resp.Secrets {
			listing.AddRow(secret.Key, secret.Version)
		}
		err = cmdutil.WriteStdout(listing)
		if err != nil {
			cmd.PrintErrf("Failed to list secrets: %+v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
