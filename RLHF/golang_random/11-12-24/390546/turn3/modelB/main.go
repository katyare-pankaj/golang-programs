package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	tokens     chan struct{}
	rate       int // Requests per second
	capacity   int // Maximum number of tokens to hold
	mu         sync.Mutex
	lastRefill time.Time
}

// NewRateLimiter initializes a new rate limiter.
func NewRateLimiter(rate, capacity int) *RateLimiter {
	return &RateLimiter{
		tokens:     make(chan struct{}, capacity),
		rate:       rate,
		capacity:   capacity,
		lastRefill: time.Now(),
	}
}

// refillTokens adds tokens to the channel at the specified rate.
func (rl *RateLimiter) refillTokens() {
	for {
		rl.mu.Lock()
		now := time.Now()
		interval := now.Sub(rl.lastRefill)
		tokensToAdd := int(interval.Seconds() * float64(rl.rate))
		if tokensToAdd > rl.capacity-len(rl.tokens) {
			tokensToAdd = rl.capacity - len(rl.tokens)
		}
		for i := 0; i < tokensToAdd; i++ {
			select {
			case rl.tokens <- struct{}{}:
			default:
			}
		}
		rl.lastRefill = now
		rl.mu.Unlock()
		time.Sleep(time.Second / time.Duration(rl.rate))
	}
}

// ServeHTTP is the HTTP handler that enforces rate limiting.
func (rl *RateLimiter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Acquire a token
	select {
	case <-rl.tokens:
	default:
		// Delay until a token is available
		<-rl.tokens
	}

	// Process the request
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Request processed at %s\n", time.Now().String())
}

func main() {
	rateLimit := 5 // requests per second
	capacity := 10 // maximum tokens in the bucket

	rateLimiter := NewRateLimiter(rateLimit, capacity)

	// Start the token replenishment goroutine
	go rateLimiter.refillTokens()

	http.Handle("/", rateLimiter)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
