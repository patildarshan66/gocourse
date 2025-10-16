package advanced

import (
	"fmt"
	"time"
)

///================== Rate Limiting - Token Bucket Algorithm =======================

type RateLimiter struct {
	tokens     chan struct{}
	refillTIme time.Duration
}

func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	r1 := &RateLimiter{
		tokens:     make(chan struct{}, rateLimit),
		refillTIme: refillTime,
	}
	for range rateLimit {
		r1.tokens <- struct{}{}
	}
	go r1.startRefill()
	return r1
}

func (rl *RateLimiter) startRefill() {

	ticker := time.NewTicker(rl.refillTIme)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			select {
			case rl.tokens <- struct{}{}:
			default:
			}
		}
	}
}

func (rl *RateLimiter) allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func main() {
	rateLimiter := NewRateLimiter(5, time.Second)

	for range 10 {
		if rateLimiter.allow() {
			fmt.Println("Request Allowed!")
		} else {
			fmt.Println("Request Denined!")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
