package counter

import (
	"sync"
	"time"
)

type RollingWindowCounter struct {
	mu                sync.Mutex
	window            time.Duration
	bucketGranularity time.Duration
	max               int
	lastBucket        int
	lastTime          time.Time
	buckets           []int
}

func NewRollingWindowCounter(window time.Duration, bucketGranularity time.Duration) *RollingWindowCounter {
	max := int(window / bucketGranularity)
	return &RollingWindowCounter{
		window:            window,
		buckets:           make([]int, max),
		max:               max,
		bucketGranularity: bucketGranularity,
	}
}

func (r *RollingWindowCounter) prune(now time.Time) {
	if now.Sub(r.lastTime) < r.bucketGranularity {
		return
	}
	diff := int(now.Sub(r.lastTime) / r.bucketGranularity)
	if diff > r.max {
		diff = r.max
	}
	for i := 0; i < diff; i++ {
		r.lastBucket = (r.lastBucket + 1) % r.max
		r.buckets[r.lastBucket] = 0
	}
	r.lastTime = now
}

func (r *RollingWindowCounter) Add() {
	r.AddWithNow(time.Now())
}

func (r *RollingWindowCounter) AddWithNow(now time.Time) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.prune(now)
	r.buckets[r.lastBucket]++
}

func (r *RollingWindowCounter) Count() int {
	return r.CountWithNow(time.Now())
}

func (r *RollingWindowCounter) CountWithNow(now time.Time) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.prune(now)
	count := 0
	for _, b := range r.buckets {
		count += b
	}
	return count
}

func (r *RollingWindowCounter) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.buckets = make([]int, r.max)
}
