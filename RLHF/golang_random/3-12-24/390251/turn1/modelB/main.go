package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Create a rate limiter channel with a capacity of 5 requests per second
	rateLimiter := time.Tick(time.Second / 5)

	// Simulate sending 10 requests
	for i := 0; i < 10; i++ {
		<-rateLimiter // Wait for the rate limiter token

		success := sendRequest(i)
		if !success {
			// If the request failed, apply backoff strategy
			backoffRetry(i)
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

// Implements an exponential backoff strategy
func backoffRetry(requestID int) {
	for retries := 0; retries < 5; retries++ {
		backoffDuration := time.Duration(math.Pow(2, float64(retries))) * time.Second
		fmt.Printf("Request %d: Backoff for %v seconds\n", requestID, backoffDuration.Seconds())
		time.Sleep(backoffDuration)

		success := sendRequest(requestID)
		if success {
			return
		}
	}

	fmt.Printf("Request %d: Failed after retries\n", requestID)
}
