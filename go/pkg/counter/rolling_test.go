package counter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRollingWindowCounter(t *testing.T) {
	r := NewRollingWindowCounter(100*time.Millisecond, 10*time.Millisecond)
	now := time.Now()
	for i := 0; i < 10; i++ {
		r.AddWithNow(now)
	}
	require.Equal(t, 10, r.CountWithNow(now))
	require.Equal(t, 0, r.CountWithNow(now.Add(100*time.Millisecond)))

	r = NewRollingWindowCounter(200*time.Millisecond, 10*time.Millisecond)
	r.AddWithNow(now.Add(50 * time.Millisecond))
	r.AddWithNow(now.Add(100 * time.Millisecond))
	r.AddWithNow(now.Add(110 * time.Millisecond))
	r.AddWithNow(now.Add(250 * time.Millisecond))
	r.AddWithNow(now.Add(300 * time.Millisecond))
	require.Equal(t, 3, r.Count())

	r = NewRollingWindowCounter(200*time.Millisecond, 10*time.Millisecond)
	r.AddWithNow(now.Add(50 * time.Millisecond))
	r.AddWithNow(now.Add(100 * time.Millisecond))
	require.Equal(t, 2, r.Count())
	r.Reset()
	require.Equal(t, 0, r.Count())
}
