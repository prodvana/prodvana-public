package conn

import (
	"fmt"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/session"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:  "list",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		authSession := session.GetSession()
		if len(authSession.Contexts) == 0 {
			cmd.Print("No auth contexts found. Run `pvnctl init` to get started.\n")
			return
		}
		for name, conn := range authSession.Contexts {
			str := fmt.Sprintf("%s: %s", name, conn.Addr)
			if name == session.InProcessContext {
				str = color.GreenString(str)
			}
			fmt.Println(str)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
