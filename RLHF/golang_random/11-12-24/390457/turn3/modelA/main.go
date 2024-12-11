package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	rate      int64         // requests per second allowed
	burst     int           // maximum burst size
	mu        sync.Mutex    // to protect the last request time
	last      time.Time     // last request time
	tokenChan chan struct{} // channel to signal available slots
}

// NewRateLimiter initializes a RateLimiter.
func NewRateLimiter(rate int64, burst int) *RateLimiter {
	rl := RateLimiter{
		rate:      rate,
		burst:     burst,
		tokenChan: make(chan struct{}, burst),
	}

	// Fill the token channel initially to allow for burst requests
	for i := 0; i < burst; i++ {
		rl.tokenChan <- struct{}{}
	}

	go rl.startTokenRefill()
	return &rl
}

// startTokenRefill fills tokens periodically according to the rate
func (rl *RateLimiter) startTokenRefill() {
	ticker := time.NewTicker(time.Second / time.Duration(rl.rate))
	defer ticker.Stop()

	for {
		<-ticker.C
		rl.mu.Lock()
		select {
		case rl.tokenChan <- struct{}{}: // refill a token if burst capacity isn't exceeded
		default: // burst capacity is reached, do nothing
		}
		rl.mu.Unlock()
	}
}

// ServeHTTP is the HTTP handler that enforces rate limiting.
func (rl *RateLimiter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Wait for a token to become available
	<-rl.tokenChan

	// Process the request knowing we have permission
	rl.mu.Lock()
	now := time.Now()

	// Delay if the last request time was too close
	if now.Sub(rl.last) < (time.Second / time.Duration(rl.rate)) {
		time.Sleep((time.Second / time.Duration(rl.rate)) - now.Sub(rl.last))
	}
	rl.last = time.Now()
	rl.mu.Unlock()

	// Respond to the request
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Request processed at %s\n", time.Now())
}

func main() {
	rateLimit := int64(5) // allow 5 requests per second
	burst := 10           // allow a burst of 10 requests

	rateLimiter := NewRateLimiter(rateLimit, burst)

	http.Handle("/", rateLimiter)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
