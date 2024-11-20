package main

import (
	"context"
	"log"
	"sync"
	"time"
)

type tokenBucket struct {
	cap     int           // Capacity of the bucket
	tokens  int           // Current number of tokens
	lastRef time.Time     // Time of the last request
	refill  time.Duration // Rate at which tokens are refilled
	m       sync.Mutex    // Mutex for synchronization
}

func newTokenBucket(capacity int, refill time.Duration) *tokenBucket {
	return &tokenBucket{
		cap:     capacity,
		tokens:  capacity,
		refill:  refill,
		lastRef: time.Now(),
	}
}

func (tb *tokenBucket) consumeToken() bool {
	tb.m.Lock()
	defer tb.m.Unlock()

	// Calculate how many tokens have been refilled since the last request
	refillElapsed := time.Now().Sub(tb.lastRef)
	newTokens := int(refillElapsed.Seconds() / tb.refill.Seconds() * float64(tb.cap))

	// Update the bucket
	if tb.tokens < newTokens {
		tb.tokens = newTokens
	}
	if tb.tokens > tb.cap {
		tb.tokens = tb.cap
	}

	// Decide whether to consume a token
	if tb.tokens == 0 {
		log.Println("No tokens available; rate limit exceeded")
		return false
	}

	tb.tokens--
	tb.lastRef = time.Now()

	return true
}

func simulateApiRequest(ctx context.Context, tb *tokenBucket) {
	select {
	case <-ctx.Done():
		return
	default:
		if tb.consumeToken() {
			log.Println("Processing API request...")
			time.Sleep(time.Duration(10) * time.Millisecond) // Simulate request processing
		}
	}
}

func main() {
	// Initialize token bucket with a capacity of 5 tokens refilled at a rate of 1 token per second
	tb := newTokenBucket(5, time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Generate goroutines to simulate API requests
	goFunc := func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				simulateApiRequest(ctx, tb)
			}
			time.Sleep(100 * time.Millisecond) // Simulate requests coming in rapidly
		}
	}

	// Start goroutines
	for i := 0; i < 10; i++ {
		go goFunc()
	}

	// Let the application run for a while to simulate some requests
	time.Sleep(5 * time.Second)
	log.Println("Simulation finished")
}
