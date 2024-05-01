package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRawJsonBytes(t *testing.T) {
	b := RawJsonBytes([]byte("{}"))
	marshalled, err := json.Marshal(b)
	require.NoError(t, err)
	require.Equal(t, "{}", string(marshalled))

	err = json.Unmarshal([]byte("\"hi\""), &b)
	require.NoError(t, err)
	marshalled, err = json.Marshal(b)
	require.NoError(t, err)
	require.Equal(t, `"hi"`, string(marshalled))
}
