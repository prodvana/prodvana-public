package main

import (
	// TODO(naphat) use public library
	"prodvana/auth/cli/session"
	"sync"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/client"
	application_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/application"
	organization_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/organization"
	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func must[T any](f func() (T, error)) func() T {
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

	getApplicationManagerClient = must(sync.OnceValues(func() (application_pb.ApplicationManagerClient, error) {
		conn, err := getProdvanaConnection()
		if err != nil {
			return nil, err
		}
		return application_pb.NewApplicationManagerClient(conn), nil
	}))

	getServiceManagerClient = must(sync.OnceValues(func() (service_pb.ServiceManagerClient, error) {
		conn, err := getProdvanaConnection()
		if err != nil {
			return nil, err
		}
		return service_pb.NewServiceManagerClient(conn), nil
	}))

	getOrganizationManagerClient = must(sync.OnceValues(func() (organization_pb.OrganizationManagerClient, error) {
		conn, err := getProdvanaConnection()
		if err != nil {
			return nil, err
		}
		return organization_pb.NewOrganizationManagerClient(conn), nil
	}))
)
