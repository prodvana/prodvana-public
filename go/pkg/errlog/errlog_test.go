package errlog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	orig := fmt.Errorf("orig")
	other := fmt.Errorf("other")
	skipOrig := Skip(orig)
	require.True(t, IsSkipped(skipOrig))
	require.ErrorIs(t, skipOrig, orig)
	require.False(t, IsSkipped(other))
	require.False(t, IsSkipped(orig))
	require.NotErrorIs(t, skipOrig, other)
}
