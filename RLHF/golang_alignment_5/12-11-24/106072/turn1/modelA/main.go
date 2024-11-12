package main

import (
	"fmt"
	"sync"
	"time"
)

// RateLimiter is a struct to manage rate limiting
type RateLimiter struct {
	maxRequestsPerSecond int
	currentRequests      int
	lastRequestTime      time.Time
	mutex                sync.Mutex
}

// NewRateLimiter creates a new RateLimiter
func NewRateLimiter(maxRequestsPerSecond int) *RateLimiter {
	return &RateLimiter{
		maxRequestsPerSecond: maxRequestsPerSecond,
	}
}

// AllowRequest checks if a new request is allowed based on the rate limit
func (r *RateLimiter) AllowRequest() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	currentTime := time.Now()
	timeElapsed := currentTime.Sub(r.lastRequestTime)

	// If enough time has elapsed, reset the counter
	if timeElapsed >= time.Second {
		r.currentRequests = 0
		r.lastRequestTime = currentTime
	}

	// Check if the rate limit is exceeded
	if r.currentRequests >= r.maxRequestsPerSecond {
		return false
	}

	r.currentRequests++
	return true
}

func main() {
	// Create a rate limiter with a maximum of 10 requests per second
	limiter := NewRateLimiter(10)

	apiCallCounter := 0

	for {
		if limiter.AllowRequest() {
			// Simulate API call
			apiCallCounter++
			fmt.Println("Making API call...", apiCallCounter)
			// Simulate the time taken for the API call
			time.Sleep(50 * time.Millisecond)
		} else {
			fmt.Println("Rate limit exceeded. Waiting...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
