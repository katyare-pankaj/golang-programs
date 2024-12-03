package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

// define variables for rate limiting
var (
	maxRequests  = 5               // maximum number of requests allowed per second
	requestLimit = 1 * time.Second // limit period (1 second in this case)
	rateLimit    = int64(0)        // current number of requests made
	lastPeriod   = time.Now()
)

func RateLimitedFunction(ctx context.Context, id int) error {
	// Check if we can make the request or need to sleep
	if err := rateLimitPermit(); err != nil {
		fmt.Printf("Request %d violated rate limit; delaying for %s.\n", id, err)
		backoffDelay(ctx)
		// If needed, retry or handle the error
		return err
	}

	// Simulate work
	time.Sleep(time.Duration(int64(100+id)) * time.Millisecond)
	fmt.Printf("Request %d completed.\n", id)

	return nil
}

// Get the current period (seconds since the epoch)
func getCurrentPeriod() time.Time {
	return time.Now().Truncate(requestLimit)
}

// Recalculate the rate and adjust lastPeriod as needed
func updateRate() {
	now := getCurrentPeriod()
	if now.Before(lastPeriod) {
		return
	}
	lastPeriod = now
	atomic.StoreInt64(&rateLimit, 0)
}

// Permit a request only if the rate limit is not exceeded
func rateLimitPermit() error {
	updateRate()

	// We want to enforce the maximum rate which is maxRequests/second
	if rateLimit < int64(maxRequests) {
		if atomic.AddInt64(&rateLimit, 1); rateLimit <= int64(maxRequests) {
			return nil
		}
	}

	// Exponentially back off for rate limit violations
	delayTime, err := exponentialBackoff()
	if err != nil {
		return err
	}

	return fmt.Errorf("rate limit exceeded, backing off for %s", delayTime)
}

// Generate an exponential backoff delay
func exponentialBackoff() (time.Duration, error) {
	return time.Second << atomic.AddUint32((*uint32)(&lastPeriod), 1), nil
}

// Check context to see if we should delay or return immediately if canceled
func backoffDelay(ctx context.Context) {
	if ctx.Err() != nil {
		return
	}
	// Implement delaying mechanism with Checking context for cancellation
	delayTime, _ := exponentialBackoff()
	select {
	case <-time.After(delayTime):
	case <-ctx.Done():
		// Handle context cancellation
	}
}

func main() {
	ctx := context.Background()

	// Spin up goroutines that attempt to access the rate-limited function
	for i := 0; i < 20; i++ {
		go func(id int) {
			if err := RateLimitedFunction(ctx, id); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}(i)
	}

	// Keep the main thread running to see the results
	time.Sleep(10 * time.Second)
}
