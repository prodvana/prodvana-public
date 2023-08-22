package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var noTlsAllowlist = map[string]struct{}{
	"apiserver":         {}, // local testing
	"agent-interaction": {}, // local testing
	"localhost":         {}, // local testing
}

type ConnectionOptions struct {
	Address          string        // Prodvana server address
	AuthToken        string        // Prodvana api token
	GetAuthToken     func() string // same as AuthToken, but for tokens that may change/need to be computed
	AuthTokenPtr     *string       // save as AuthToken but passes a pointer, for tokens that need to be updated in process
	SkipTls          bool          // exposed for testing purposes only
	NoAuth           bool          // exposed for internal purposes only
	ExtraDialOptions []grpc.DialOption
}

func DefaultConnectionOptions() ConnectionOptions {
	return ConnectionOptions{
		Address:   os.Getenv("PVN_APISERVER_ADDR"),
		AuthToken: os.Getenv("PVN_TOKEN"),
	}
}

type authToken struct {
	MakeToken func() string
}

func (t authToken) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + t.MakeToken(),
	}, nil
}

func (authToken) RequireTransportSecurity() bool {
	return false
}

func MakeProdvanaConnection(options ConnectionOptions) (*grpc.ClientConn, error) {
	addr := options.Address
	if addr == "" {
		return nil, fmt.Errorf("either Address or PVN_APISERVER_ADDR must be set")
	}
	token := options.AuthToken
	if token == "" && options.GetAuthToken == nil && !options.NoAuth {
		return nil, fmt.Errorf("either AuthToken or PVN_TOKEN must be set")
	}
	parts := strings.Split(addr, ":")
	var cred credentials.TransportCredentials
	_, noTlsHost := noTlsAllowlist[parts[0]]
	if options.SkipTls || noTlsHost {
		cred = insecure.NewCredentials()
	} else {
		cred = credentials.NewTLS(&tls.Config{ServerName: parts[0]})
	}
	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(cred),
	}
	if !options.NoAuth {
		getToken := options.GetAuthToken
		if getToken == nil {
			getToken = func() string { return options.AuthToken }
		}
		grpcOpts = append(grpcOpts, grpc.WithPerRPCCredentials(authToken{getToken}))
	}
	return grpc.Dial(addr, grpcOpts...)
}
