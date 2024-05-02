package timeutils

import (
	"math"
	"math/rand"
	"time"
)

func Jitter50(d time.Duration) time.Duration {
	return JitterPercent(d, 50)
}

func Jitter10(d time.Duration) time.Duration {
	return JitterPercent(d, 10)
}

// return jitter +/-jitterPercent
func JitterPercent(d time.Duration, jitterPercent int) time.Duration {
	if jitterPercent == 0 {
		return d
	}
	// mostly copied from retry.WithJitterPercent
	top := rand.Int63n(int64(jitterPercent)*2) - int64(jitterPercent)
	pct := 1 - float64(top)/100.0
	return time.Duration(float64(d) * pct)
}

// return jitter +jitterPercent
func JitterPercentPositive(d time.Duration, jitterPercent int) time.Duration {
	if jitterPercent == 0 {
		return d
	}
	top := rand.Int63n(int64(jitterPercent))
	pct := 1 + float64(top)/100.0
	return time.Duration(float64(d) * pct)
}

func ExpBackoff(min, max time.Duration, iteration int) time.Duration {
	const decay = 2
	scale := math.Pow(decay, float64(iteration))
	result := min * time.Duration(scale)
	if result > max {
		return max
	}
	return result
}
