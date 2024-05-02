package auth

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"prodvana/cmd/cmdutil"
	"strings"

	"github.com/prodvana/prodvana-public/go/pkg/client"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/session"

	auth_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/auth"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Login(ctx context.Context, admin bool) error {
	resp, err := client.GetAuthManagerClient().GetAuthUrl(ctx, &auth_pb.GetAuthUrlRequest{
		ForCli:               true,
		ProdvanaOnlyAdminOrg: admin,
	})
	if err != nil {
		return errors.Wrap(err, "Failed to talk to Prodvana")
	}
	fmt.Printf("Please login at %s\n\n", resp.AuthUrl)
	fmt.Printf("Paste the output of the screen after login: ")
	var redirectUrl string
	fmt.Scanln(&redirectUrl)
	redirectUrl = strings.TrimSpace(redirectUrl)
	u, err := url.Parse(redirectUrl)
	if err != nil {
		return errors.Wrap(err, "Invalid url")
	}
	query, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return errors.Wrap(err, "Invalid url query")
	}
	if len(query["error"]) != 0 {
		return fmt.Errorf("%+v: %+v", query["error"], query["error_description"])
	}
	code := query["code"]
	if len(code) != 1 {
		return errors.Errorf("Incorrect url, no code specified")
	}
	token, err := client.GetAuthManagerClient().GetAuthToken(ctx, &auth_pb.GetTokenRequest{
		Code:                 code[0],
		ForCli:               true,
		ProdvanaOnlyAdminOrg: admin,
	})
	if err != nil {
		return errors.Wrap(err, "Failed to get auth token")
	}

	authSession := session.GetSession()
	authSession.Contexts[session.InProcessContext].AuthToken = token.Token
	authSession.Contexts[session.InProcessContext].ApiToken = ""
	err = session.SaveSession(authSession)
	if err != nil {
		return err
	}
	return nil
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate via the web browser.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		admin, err := cmd.Flags().GetBool("admin")
		if err != nil {
			log.Fatal(err)
		}
		err = Login(ctx, admin)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().Bool("admin", false, "")
	cmdutil.Must(loginCmd.Flags().MarkHidden("admin"))
}
