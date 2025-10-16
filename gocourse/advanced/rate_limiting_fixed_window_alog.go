package advancced

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{limit: limit, window: window}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}

	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false

}

func main() {

	rateLimter := NewRateLimiter(3, time.Second)

	for range 10 {
		if rateLimter.Allow() {
			fmt.Println("Request Allowed.")
		} else {
			fmt.Println("Request Denied.")
		}
		time.Sleep(200 * time.Millisecond)
	}

}
