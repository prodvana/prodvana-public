package timeutils

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestJitter(t *testing.T) {
	const base = 1 * time.Minute
	const iterations = 1000
	for _, tc := range []struct {
		percent    int
		lowerBound time.Duration
		upperBound time.Duration
	}{
		{
			percent:    50,
			lowerBound: 30 * time.Second,
			upperBound: 90 * time.Second,
		},
		{
			percent:    10,
			lowerBound: 54 * time.Second,
			upperBound: 66 * time.Second,
		},
	} {
		t.Run(fmt.Sprintf("%d_%v_%v", tc.percent, tc.lowerBound, tc.upperBound), func(t *testing.T) {
			lowerCount := 0
			greaterCount := 0
			for i := 0; i < iterations; i++ {
				result := JitterPercent(base, tc.percent)
				require.GreaterOrEqual(t, result, tc.lowerBound)
				require.LessOrEqual(t, result, tc.upperBound)
				if result > base {
					greaterCount++
				} else {
					lowerCount++
				}
			}
			// assert that we get values that are reasonable spread between +/- 50%
			require.Greater(t, float64(lowerCount), iterations*0.25)
			require.Greater(t, float64(greaterCount), iterations*0.25)
		})
	}
}

func TestJitterPositive(t *testing.T) {
	const base = 1 * time.Minute
	const iterations = 1000
	for _, tc := range []struct {
		percent    int
		lowerBound time.Duration
		upperBound time.Duration
	}{
		{
			percent:    50,
			lowerBound: 60 * time.Second,
			upperBound: 90 * time.Second,
		},
		{
			percent:    10,
			lowerBound: 60 * time.Second,
			upperBound: 66 * time.Second,
		},
	} {
		t.Run(fmt.Sprintf("%d_%v_%v", tc.percent, tc.lowerBound, tc.upperBound), func(t *testing.T) {
			midPoint := (tc.lowerBound + tc.upperBound) / 2.0
			lowerCount := 0
			greaterCount := 0
			for i := 0; i < iterations; i++ {
				result := JitterPercentPositive(base, tc.percent)
				require.GreaterOrEqual(t, result, tc.lowerBound)
				require.LessOrEqual(t, result, tc.upperBound)
				if result > midPoint {
					greaterCount++
				} else {
					lowerCount++
				}
			}
			// assert that we get values that are reasonable spread between +/- 50% of midpoint
			require.Greater(t, float64(lowerCount), iterations*0.25)
			require.Greater(t, float64(greaterCount), iterations*0.25)
		})
	}
}
