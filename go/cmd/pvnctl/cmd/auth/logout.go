package auth

import (
	"log"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/session"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		authSession := session.GetSession()
		authSession.Contexts[session.InProcessContext].AuthToken = nil
		authSession.Contexts[session.InProcessContext].ApiToken = ""
		err := session.SaveSession(authSession)
		if err != nil {
			log.Fatalf("Could not save session: %s", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(logoutCmd)
}
