package cmdutil

import (
	"context"
	"log"

	"github.com/prodvana/prodvana-public/go/pkg/perrors"

	"github.com/prodvana/prodvana-public/go/pkg/client"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/session"

	auth_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/auth"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	grpc_codes "google.golang.org/grpc/codes"
)

var (
	ApiServerAddrOverride string
)

func ValidateAndRefreshAuthPreRunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	_, err := client.GetAuthSessionManagerClient().Check(ctx, &auth_pb.Empty{})
	if err == nil {
		return nil
	}
	if perrors.GrpcErrCode(err) != grpc_codes.Unauthenticated {
		log.Printf("WARNING: Unable to check if user is authenticated.\n")
		return nil
	}
	authSession := session.GetSession()
	authContext := authSession.Contexts[session.InProcessContext]
	if authContext.ApiToken != "" {
		return errors.Errorf("Invalid or expired API token\n")
	}
	if authContext.AuthToken == nil {
		return errors.Errorf("Not logged in. Please run: pvnctl auth login\n")
	}
	newToken, err := client.GetAuthManagerClient().RefreshToken(ctx, &auth_pb.RefreshTokenReq{
		Token:  authContext.AuthToken,
		ForCli: true,
	})
	if err != nil {
		return errors.Wrap(err, "Unable to refresh token")
	}
	authContext.AuthToken = newToken.Token
	authContext.ApiToken = ""
	err = session.SaveSession(authSession)
	if err != nil {
		return err
	}
	if err := client.UpdateApiserverToken(newToken.Token.Token); err != nil {
		return err
	}
	return nil
}

// Using PreRunE causes the help text to be printed on error, which is not what we want.
func ValidateAndRefreshAuthPreRun(cmd *cobra.Command, args []string) {
	sess := session.GetSession()
	if session.InProcessContext == "" {
		log.Fatal("No auth context found. Please run: `pvnctl init` to get started.")
	}
	var addr string
	var skipTls bool
	if ApiServerAddrOverride != "" {
		addr = ApiServerAddrOverride
	} else {
		authCtx, ok := sess.Contexts[session.InProcessContext]
		if !ok {
			log.Fatalf("Unknown context %s\n", session.InProcessContext)
		}
		addr = authCtx.Addr
		skipTls = authCtx.SkipTls
	}
	err := client.SetDefaultApiServerAddr(addr, skipTls)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SetWaitForReady(false) // since pvnctl is (mostly) interactive, it is better to fail fast when failing to connect to Prodvana
	if err != nil {
		log.Fatal(err)
	}
	err = client.SetSource("pvnctl")
	if err != nil {
		log.Fatal(err)
	}

	err = ValidateAndRefreshAuthPreRunE(cmd, args)
	if err != nil {
		log.Fatal(err)
	}
}
