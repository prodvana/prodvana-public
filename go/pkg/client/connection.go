package client

import (
	"context"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"

	client_auth "github.com/prodvana/prodvana-public/go/pkg/auth/client"
	"github.com/prodvana/prodvana-public/go/prodvana-sdk/client"
	auth_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/auth"

	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	ApiServerAddr         string
	defaultApiServerNoTls bool
	apiserverConn         *grpc.ClientConn
	apiserverConnOnce     sync.Once
	apiserverToken        *client_auth.AuthToken

	authManagerClient     auth_pb.AuthManagerClient
	authManagerClientOnce sync.Once

	authSessionManagerClient     auth_pb.AuthSessionManagerClient
	authSessionManagerClientOnce sync.Once

	AdminServerAddr     string
	adminServerConn     *grpc.ClientConn
	adminServerConnOnce sync.Once

	GlobalApiAddr     string
	globalApiConn     *grpc.ClientConn
	globalApiConnOnce sync.Once

	internalTrafficApiserverRegex = regexp.MustCompile(`^apiserver\.[a-z0-9\-_]+$`) // internal traffic within k8s cluster.

	traceAPIRequests       bool
	grpcClientWaitForReady = true
	grpcSource             string
)

func SetWaitForReady(waitForReady bool) error {
	if apiserverConn != nil || globalApiConn != nil || adminServerConn != nil {
		return errors.Errorf("Cannot change wait for ready after connection has been initialized")
	}
	grpcClientWaitForReady = waitForReady
	return nil
}

func SetDefaultApiServerAddr(addr string, noTls bool) error {
	if apiserverConn != nil || globalApiConn != nil || adminServerConn != nil {
		return errors.Errorf("Cannot change apiserver address after connection has been initialized")
	}
	ApiServerAddr = addr
	defaultApiServerNoTls = noTls
	return nil
}

func EnableTraceAPIRequests() {
	traceAPIRequests = true
}

func SetSource(source string) error {
	if apiserverConn != nil || globalApiConn != nil || adminServerConn != nil {
		return errors.Errorf("Cannot change source after connection has been initialized")
	}
	grpcSource = source
	return nil
}

func TraceUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		callOpts ...grpc.CallOption,
	) error {
		startTime := time.Now()
		err := invoker(ctx, method, req, reply, cc, callOpts...)
		if traceAPIRequests {
			log.Printf("Method %s completed in %s", method, time.Since(startTime))
		}
		return err
	}
}

func sourceUnaryClientInterceptor(source string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		callOpts ...grpc.CallOption,
	) error {
		ctx = metadata.AppendToOutgoingContext(ctx, "pvn-source", source)
		return invoker(ctx, method, req, reply, cc, callOpts...)
	}
}

func sourceStreamClientInterceptor(source string) grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		callOpts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		ctx = metadata.AppendToOutgoingContext(ctx, "pvn-source", source)
		return streamer(ctx, desc, cc, method, callOpts...)
	}
}

func MakeApiServerConnWithAuth(addr string, token *client_auth.AuthToken, noTls bool) (*grpc.ClientConn, error) {
	parts := strings.Split(addr, ":")
	serverName := parts[0]
	internalTraffic := internalTrafficApiserverRegex.MatchString(serverName)
	conn, err := client.MakeProdvanaConnection(client.ConnectionOptions{
		Address: addr,
		GetAuthToken: func() string {
			if token != nil {
				return token.Token
			}
			return ""
		},
		SkipTls: noTls || internalTraffic,
		NoAuth:  token == nil || token.Token == "", // we have places that pass empty token string as no auth instead of
		ExtraDialOptions: []grpc.DialOption{
			grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
			grpc.WithChainUnaryInterceptor(TraceUnaryClientInterceptor(), sourceUnaryClientInterceptor(grpcSource)),
			grpc.WithStreamInterceptor(sourceStreamClientInterceptor(grpcSource)),
			grpc.WithDefaultCallOptions(
				grpc.WaitForReady(grpcClientWaitForReady),
				grpc.MaxCallRecvMsgSize(200*1024*1024), // 200MB. Allow client to read large responses from APIServer
			),
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create connection")
	}
	return conn, nil
}

func MakeUnauthApiServerConn(addr string, noTls bool) (*grpc.ClientConn, error) {
	return MakeApiServerConnWithAuth(addr, nil, noTls)
}

func apiServerConn() (*grpc.ClientConn, *client_auth.AuthToken) {
	token, err := client_auth.GetAuthToken()
	if err != nil {
		log.Fatal(err)
	}
	conn, err := MakeApiServerConnWithAuth(ApiServerAddr, token, defaultApiServerNoTls)
	if err != nil {
		log.Fatal(err)
	}
	return conn, token
}

func GetApiserverConnWithAuth(token *client_auth.AuthToken) *grpc.ClientConn {
	apiserverConnOnce.Do(func() {
		var err error
		apiserverConn, err = MakeApiServerConnWithAuth(ApiServerAddr, token, defaultApiServerNoTls)
		if err != nil {
			panic(err)
		}
		apiserverToken = token
	})

	return apiserverConn
}

func GetApiserverConn() *grpc.ClientConn {
	apiserverConnOnce.Do(func() {
		apiserverConn, apiserverToken = apiServerConn()
	})

	return apiserverConn
}

func GetAdminServerConn() *grpc.ClientConn {
	adminServerConnOnce.Do(func() {
		conn, err := grpc.NewClient(AdminServerAddr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
			grpc.WithDefaultCallOptions(grpc.WaitForReady(grpcClientWaitForReady)),
		)
		if err != nil {
			log.Fatalf("Failed to connect to admin server: %s", err)
		}
		adminServerConn = conn
	})
	return adminServerConn
}

func GetAuthManagerClient() auth_pb.AuthManagerClient {
	authManagerClientOnce.Do(func() {
		authManagerClient = auth_pb.NewAuthManagerClient(GetApiserverConn())
	})
	return authManagerClient
}

func GetAuthSessionManagerClient() auth_pb.AuthSessionManagerClient {
	authSessionManagerClientOnce.Do(func() {
		authSessionManagerClient = auth_pb.NewAuthSessionManagerClient(GetApiserverConn())
	})
	return authSessionManagerClient
}

// HACK(naphat) this is not threadsafe
func UpdateApiserverToken(token string) error {
	if apiserverToken == nil {
		return errors.Errorf("No token set")
	}
	apiserverToken.Token = token
	return nil
}

func SetDefaultGlobalApiAddr(addr string) error {
	if globalApiConn != nil {
		return errors.Errorf("Cannot change global api address after connection has been initialized")
	}
	GlobalApiAddr = addr
	return nil
}

func GetGlobalApiConn() *grpc.ClientConn {
	globalApiConnOnce.Do(func() {
		conn, err := grpc.NewClient(GlobalApiAddr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
			grpc.WithDefaultCallOptions(grpc.WaitForReady(grpcClientWaitForReady)),
		)
		if err != nil {
			log.Fatalf("Failed to connect to global api: %s", err)
		}
		globalApiConn = conn
	})
	return globalApiConn
}
