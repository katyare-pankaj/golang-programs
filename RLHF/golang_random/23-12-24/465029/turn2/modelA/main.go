package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sharedCounter = 0
	mu            sync.Mutex
)

func incrementCounter(n int) {
	for i := 0; i < n; i++ {
		mu.Lock() // Acquire the lock before updating the counter
		sharedCounter++
		mu.Unlock() // Release the lock after updating the counter
	}
}

func main() {
	const (
		numGoroutines          = 100
		iterationsPerGoroutine = 1000000
	)

	// Start multiple goroutines to increment the counter
	start := time.Now()
	for i := 0; i < numGoroutines; i++ {
		go incrementCounter(iterationsPerGoroutine)
	}

	// Wait for all goroutines to finish
	for i := 0; i < numGoroutines; i++ {
		time.Sleep(10 * time.Millisecond) // Allow some time for goroutines to complete
	}

	// Print the final value of the shared counter
	fmt.Printf("Final shared counter value: %d\n", sharedCounter)
	fmt.Printf("Elapsed time: %v\n", time.Since(start))
}
