package conn

import (
	"log"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/session"

	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:  "use",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		authSession := session.GetSession()
		if _, ok := authSession.Contexts[name]; !ok {
			log.Fatalf("Unknown auth context: %s", name)
		}
		authSession.CurrentContext = name
		err := session.SaveSession(authSession)
		if err != nil {
			log.Fatalf("Could not save session: %s", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(useCmd)
}
