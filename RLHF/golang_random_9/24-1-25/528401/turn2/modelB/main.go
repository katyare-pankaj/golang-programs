package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	numRequests    = 10
	requestTimeout = 2 * time.Second
)

func processRequest(ctx context.Context, wg *sync.WaitGroup, requestID int) {
	defer wg.Done()

	// Simulate some work for the request
	select {
	case <-time.After(time.Duration((requestID+1)%4) * time.Second):
		fmt.Printf("Request %d completed successfully\n", requestID)
	case <-ctx.Done():
		fmt.Printf("Request %d timed out\n", requestID)
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	// Start processing multiple requests concurrently
	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go processRequest(ctx, &wg, i)
	}

	// Wait for all requests to complete or timeout
	wg.Wait()

	fmt.Println("All requests processed.")
}
