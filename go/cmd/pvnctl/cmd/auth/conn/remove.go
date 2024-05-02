package conn

import (
	"os"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/session"

	"github.com/spf13/cobra"
)

var rmContextCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove an auth context.",
	Long: `Remove an auth context.

pvnctl auth context rm <context name>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		// no need to check resp of health check, all we care is that the connection succeeded
		authSession := session.GetSession()
		_, ok := authSession.Contexts[name]
		if !ok {
			cmd.PrintErrf("error: context %s does not exist\n", name)
			os.Exit(1)
			return
		}
		delete(authSession.Contexts, name)
		if authSession.CurrentContext == name && len(authSession.Contexts) > 0 {
			for k := range authSession.Contexts {
				authSession.CurrentContext = k
				session.InProcessContext = k
				break
			}
		} else {
			authSession.CurrentContext = ""
			session.InProcessContext = ""
		}
		err := session.SaveSession(authSession)
		if err != nil {
			cmd.PrintErrf("error: failed to remove context %s: %+v\n", name, err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(rmContextCmd)
}
