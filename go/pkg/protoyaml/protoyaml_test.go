package protoyaml

import (
	config_pb "prodvana/proto/prodvana_internal/config"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestMissing(t *testing.T) {
	msg := `
org_id: foo
new_field_not_known_to_proto: bar
`
	var config config_pb.Config
	require.Error(t, Unmarshal([]byte(msg), &config))
	require.NoError(t, UnmarshalOptions{protojson.UnmarshalOptions{DiscardUnknown: true}}.Unmarshal([]byte(msg), &config))
	require.Equal(t, "foo", config.OrgId)
}
