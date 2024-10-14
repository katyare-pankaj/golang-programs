package main

import (
	"fmt"
	"sync"
	"time"
)

// Bucket is a struct to represent a rate limiter bucket
type Bucket struct {
	tokens     int
	rate       int // tokens added per second
	lastUpdate time.Time
}

// NewBucket creates a new Bucket with the specified rate
func NewBucket(rate int) *Bucket {
	return &Bucket{
		rate:       rate,
		lastUpdate: time.Now(),
	}
}

// Allow checks if the bucket allows an action based on the current token count
func (b *Bucket) Allow() bool {
	now := time.Now()
	timeElapsed := now.Sub(b.lastUpdate)

	// Add tokens based on the elapsed time
	b.tokens += int(timeElapsed.Seconds()) * b.rate
	b.lastUpdate = now

	if b.tokens > 0 {
		b.tokens--
		return true
	}
	return false
}

// ApiGateway is our API gateway struct
type ApiGateway struct {
	mu        sync.Mutex
	limiter   map[string]*Bucket
	limitRate int
}

// NewApiGateway creates a new ApiGateway with the specified limit rate per second
func NewApiGateway(limitRate int) *ApiGateway {
	return &ApiGateway{
		limiter:   make(map[string]*Bucket),
		limitRate: limitRate,
	}
}

// HandleRequest handles an API request and checks for rate limiting
func (ag *ApiGateway) HandleRequest(userID string) bool {
	ag.mu.Lock()
	defer ag.mu.Unlock()

	// Create a bucket for the user if it doesn't exist
	if _, ok := ag.limiter[userID]; !ok {
		ag.limiter[userID] = NewBucket(ag.limitRate)
	}

	bucket := ag.limiter[userID]
	return bucket.Allow()
}

func main() {
	// Create an API gateway with a limit rate of 5 requests per second per user
	apiGateway := NewApiGateway(5)

	userID := "user123"

	for i := 0; i < 10; i++ {
		if apiGateway.HandleRequest(userID) {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request rate limited")
		}
		time.Sleep(time.Second)
	}
}
