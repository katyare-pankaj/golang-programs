package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// RateLimiter struct represents the token bucket rate limiter
type RateLimiter struct {
	mu     sync.Mutex
	tokens int
	cap    int
	rate   time.Duration
	last   time.Time
}

// NewRateLimiter creates a new RateLimiter with the given capacity and rate.
func NewRateLimiter(cap int, rate time.Duration) *RateLimiter {
	return &RateLimiter{
		cap:  cap,
		rate: rate,
		last: time.Now(),
	}
}

// Allow checks if the request can be allowed based on the rate limit.
func (l *RateLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	l.tokens += int(now.Sub(l.last) / l.rate)
	if l.tokens > l.cap {
		l.tokens = l.cap
	}

	if l.tokens > 0 {
		l.tokens--
		l.last = now
		return true
	}
	return false
}

func main() {
	// Create a new rate limiter with a capacity of 10 requests per second
	rateLimiter := NewRateLimiter(10, time.Second)

	// Simulate data and serve requests
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		if !rateLimiter.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		// Simulate data processing time
		data := simulateData(rand.Intn(100))

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Simulated Data: %s\n", data)
	})

	fmt.Println("Server listening on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// simulateData function simulates some processing and generates a random data string.
func simulateData(duration int) string {
	// Simulate processing time for 'duration' milliseconds
	time.Sleep(time.Duration(duration) * time.Millisecond)

	// Generate a random string
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
