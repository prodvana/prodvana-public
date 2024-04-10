package main

import (
	"sync"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/client"
	application_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/application"
	environment_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environment"
	organization_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/organization"
	secrets_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/secrets"
	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"
	"github.com/prodvana/prodvana-public/go/prodvana-sdk/session"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func mustValueConstructor[T any](f func() (T, error)) func() T {
	return func() T {
		v, err := f()
		if err != nil {
			panic(err)
		}
		return v
	}
}

var (
	getProdvanaConnection = sync.OnceValues(func() (*grpc.ClientConn, error) {
		s := session.GetSession()
		authCtx, ok := s.Contexts[session.InProcessContext]
		if !ok {
			return nil, errors.Errorf("pvnctl not authenticated")
		}
		options := client.DefaultConnectionOptions()
		options.Address = authCtx.Addr
		if authCtx.AuthToken != nil {
			options.AuthToken = authCtx.AuthToken.Token
		} else {
			options.AuthToken = authCtx.ApiToken
		}
		return client.MakeProdvanaConnection(options)
	})

	getApplicationManagerClient = mustValueConstructor(sync.OnceValues(func() (application_pb.ApplicationManagerClient, error) {
		conn, err := getProdvanaConnection()
		if err != nil {
			return nil, err
		}
		return application_pb.NewApplicationManagerClient(conn), nil
	}))

	getServiceManagerClient = mustValueConstructor(sync.OnceValues(func() (service_pb.ServiceManagerClient, error) {
		conn, err := getProdvanaConnection()
		if err != nil {
			return nil, err
		}
		return service_pb.NewServiceManagerClient(conn), nil
	}))

	getOrganizationManagerClient = mustValueConstructor(sync.OnceValues(func() (organization_pb.OrganizationManagerClient, error) {
		conn, err := getProdvanaConnection()
		if err != nil {
			return nil, err
		}
		return organization_pb.NewOrganizationManagerClient(conn), nil
	}))

	getEnvironmentManagerClient = mustValueConstructor(sync.OnceValues(func() (environment_pb.EnvironmentManagerClient, error) {
		conn, err := getProdvanaConnection()
		if err != nil {
			return nil, err
		}
		return environment_pb.NewEnvironmentManagerClient(conn), nil
	}))

	getSecretsManagerClient = mustValueConstructor(sync.OnceValues(func() (secrets_pb.SecretsManagerClient, error) {
		conn, err := getProdvanaConnection()
		if err != nil {
			return nil, err
		}
		return secrets_pb.NewSecretsManagerClient(conn), nil
	}))
)
