package auth

import (
	"prodvana/cmd/pvnctl/cmd/auth/conn"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "auth <subcommand>",
	Short: "Log in and verify creds",
	Long: `Log in and verify creds.

pvnctl auth login
pvnctl auth verify
`,
}

func init() {
	RootCmd.AddCommand(conn.RootCmd)
}
