package client

import (
	"context"
	"os"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/session"
	"google.golang.org/grpc"
)

type AuthToken struct {
	Token string
}

func (t AuthToken) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + t.Token,
	}, nil
}

func (AuthToken) RequireTransportSecurity() bool {
	return false
}

func GetAuthToken() (*AuthToken, error) {
	apiToken := os.Getenv("PVN_AUTH_TOKEN")
	if apiToken != "" {
		return &AuthToken{Token: apiToken}, nil
	}

	var token string
	s := session.GetSession()

	currentContext := s.Contexts[session.InProcessContext]
	if currentContext != nil {
		if currentContext.ApiToken != "" {
			token = currentContext.ApiToken
		} else if currentContext.AuthToken != nil && currentContext.AuthToken.Token != "" {
			token = currentContext.AuthToken.Token
		}
	}

	return &AuthToken{Token: token}, nil
}

func GetUnaryClientInterceptor(errorHandler func(error) error) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			newErr := errorHandler(err)
			if newErr != nil {
				return err
			}
		}
		return err
	}
}
