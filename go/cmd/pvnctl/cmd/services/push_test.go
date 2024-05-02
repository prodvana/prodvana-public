package services

import (
	"testing"

	"github.com/prodvana/prodvana-public/go/pkg/protoyaml"

	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"

	"github.com/stretchr/testify/require"
)

func TestServiceConfigTemplate(t *testing.T) {
	var serviceConfig service_pb.ServiceConfig
	require.NoError(t, protoyaml.Unmarshal([]byte(defaultServiceConfigYaml), &serviceConfig))
	require.Equal(t, "<service name>", serviceConfig.Name)
}
