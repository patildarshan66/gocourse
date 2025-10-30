package main

import (
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity int
	leakRate time.Duration
	tokens   int
	lastLeak time.Time
	mu       sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens:   capacity,
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()
	elapsedTime := now.Sub(lb.lastLeak)
	tokensToAdd := int(elapsedTime / lb.leakRate)
	lb.tokens += tokensToAdd

	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}

	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)

	if lb.tokens > 0 {
		lb.tokens--
		return true
	}

	return false
}

func main() {
	leakyBucketInstance := NewLeakyBucket(6, 500*time.Millisecond)
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if leakyBucketInstance.Allow() {
				println("Request allowed")
			} else {
				println("Request denied")
			}
			time.Sleep(200 * time.Millisecond)
		}()
	}
	wg.Wait()
}
