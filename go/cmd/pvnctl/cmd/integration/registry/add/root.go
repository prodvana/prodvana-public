package add

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "add <subcommand>",
	Short: "Add Container Registry.",
	Long: `Add Container Registry.

pvnctl integration registry add docker <name> --url <url> --username <username> --password <password>
pvnctl integration registry add ecr <name> --access-key <key> --secret-key <secret> --region <region>
`,
}
