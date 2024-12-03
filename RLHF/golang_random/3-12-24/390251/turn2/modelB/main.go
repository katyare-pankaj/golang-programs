package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create a rate limiter channel that allows 5 requests per second
	rateLimiter := time.Tick(time.Second / 5)

	// Simulate sending 10 requests
	for i := 0; i < 10; i++ {
		<-rateLimiter // Wait for the rate limiter token

		success := sendRequest(i)
		if !success {
			// If the request failed, apply fixed backoff strategy
			fixedBackoffRetry(i)
		}
	}
}

// Simulates sending a request and returns false if rate limited
func sendRequest(requestID int) bool {
	// Simulate a random rate limit failure
	if rand.Float32() < 0.3 { // 30% chance of being rate limited
		fmt.Printf("Request %d: Rate limited, will retry...\n", requestID)
		return false
	}

	fmt.Printf("Request %d: Successfully processed.\n", requestID)
	return true
}

// Implements a fixed backoff strategy
func fixedBackoffRetry(requestID int) {
	const maxRetries = 5
	const backoffDuration = 2 * time.Second // Fixed backoff duration

	for retries := 0; retries < maxRetries; retries++ {
		fmt.Printf("Request %d: Backoff for %v seconds\n", requestID, backoffDuration.Seconds())
		time.Sleep(backoffDuration)

		success := sendRequest(requestID)
		if success {
			return
		}
	}

	fmt.Printf("Request %d: Failed after retries\n", requestID)
}
