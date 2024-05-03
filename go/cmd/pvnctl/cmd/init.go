package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/auth"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/auth/conn"

	"github.com/prodvana/prodvana-public/go/pkg/client"

	"github.com/spf13/cobra"
)

var (
	initDomainSuffix string
	initPort         int
	initAuthContext  string
	initOrg          string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize pvnctl.",
	Long: `Initialize pvnctl and login to your account. This command is a shortcut for:

pvnctl auth context add default ...
pvnctl auth login
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		admin, err := cmd.Flags().GetBool("admin")
		if err != nil {
			log.Fatal(err)
		}

		org := initOrg
		if org == "" {
			org = cmdutil.Prompt("Organization Slug (the part of your Prodvana subdomain before .runprodvana.com)")
		}
		authContext := initAuthContext
		if authContext == "" {
			authContext = org
		}

		ctx := context.Background()

		addr := fmt.Sprintf("%s%s:%d", org, initDomainSuffix, initPort)
		err = conn.ValidateAndSaveAuthContext(ctx, authContext, addr, false)
		if err != nil {
			log.Fatal(err)
		}
		err = client.SetDefaultApiServerAddr(addr, false)
		if err != nil {
			log.Fatal(err)
		}
		err = auth.Login(ctx, admin)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	// customizations exposed for testing purposes
	initCmd.Flags().StringVar(&initDomainSuffix, "domain-suffix", ".grpc.runprodvana.com", "Prodvana domain suffix.")
	cmdutil.Must(initCmd.Flags().MarkHidden("domain-suffix"))
	initCmd.Flags().IntVar(&initPort, "port", 443, "Prodvana api port.")
	cmdutil.Must(initCmd.Flags().MarkHidden("port"))
	initCmd.Flags().StringVar(&initAuthContext, "auth-context", "", "Auth context to use to refer to this organization locally. Default to organization slug.")

	initCmd.Flags().StringVar(&initOrg, "org-slug", "", "The slug for Prodvana organization to connect to, i.e. the org-slug part of <org-slug>.runprodvana.com. If not specified, the command will prompt.")

	initCmd.Flags().Bool("admin", false, "")
	cmdutil.Must(initCmd.Flags().MarkHidden("admin"))
}
