package conn

import (
	"context"
	"log"

	"github.com/prodvana/prodvana-public/go/pkg/client"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/session"

	auth_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/auth"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/health/grpc_health_v1"
)

var addFlags = struct {
	skipTls bool
}{}

func ValidateAndSaveAuthContext(ctx context.Context, name, addr string, skipTls bool) error {
	conn, err := client.MakeUnauthApiServerConn(addr, skipTls)
	if err != nil {
		return err
	}
	defer conn.Close()
	healthClient := grpc_health_v1.NewHealthClient(conn)
	_, err = healthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
	if err != nil {
		return errors.Wrap(err, "failed to connect to Prodvana")
	}

	// no need to check resp of health check, all we care is that the connection succeeded
	authSession := session.GetSession()
	if _, ok := authSession.Contexts[name]; !ok {
		if authSession.Contexts == nil {
			authSession.Contexts = map[string]*auth_pb.AuthContext{}
		}
		authSession.Contexts[name] = &auth_pb.AuthContext{}
	}
	authSession.Contexts[name].Addr = addr
	authSession.CurrentContext = name
	err = session.SaveSession(authSession)
	if err != nil {
		return err
	}
	session.InProcessContext = name
	return nil
}

var contextCmd = &cobra.Command{
	Use:  "add",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		addr := args[1]
		ctx := context.Background()
		err := ValidateAndSaveAuthContext(ctx, name, addr, addFlags.skipTls)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	contextCmd.Flags().BoolVar(&addFlags.skipTls, "skip-tls", false, "Skip TLS verification")
	RootCmd.AddCommand(contextCmd)
}
