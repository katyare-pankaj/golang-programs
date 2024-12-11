package main

import (
	"fmt"
	"net/http"
	"time"
)

type RateLimiter struct {
	tokens    chan struct{}
	resetTime time.Duration
}

// NewRateLimiter initializes a rate limiter with a specified rate limit and reset duration.
func NewRateLimiter(rateLimit int, resetTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:    make(chan struct{}, rateLimit),
		resetTime: resetTime,
	}

	go func() {
		ticker := time.NewTicker(resetTime)
		defer ticker.Stop()

		for range ticker.C {
			for i := 0; i < rateLimit; i++ {
				rl.tokens <- struct{}{}
			}
		}
	}()

	return rl
}

// ServeHTTP is the HTTP handler that enforces rate limiting.
func (rl *RateLimiter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	select {
	case <-rl.tokens:
		// Process the request
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Request processed at %s\n", time.Now().String())
	default:
		// Rate limit exceeded
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
	}
}

func main() {
	rateLimit := 5               // requests per interval
	resetTime := 1 * time.Second // time interval

	rateLimiter := NewRateLimiter(rateLimit, resetTime)

	http.Handle("/", rateLimiter)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
