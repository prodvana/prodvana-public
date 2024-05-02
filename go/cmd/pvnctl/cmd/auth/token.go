package auth

import (
	"fmt"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/session"

	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Authenticate via an api token.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		authSession := session.GetSession()
		authCtx, ok := authSession.Contexts[session.InProcessContext]
		if !ok {
			cmd.PrintErrf("error: auth context %s does not exist\n", session.InProcessContext)
			return
		}
		authCtx.AuthToken = nil
		authCtx.ApiToken = token
		err := session.SaveSession(authSession)
		if err != nil {
			fmt.Printf("Could not save session: %s", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(tokenCmd)
}
