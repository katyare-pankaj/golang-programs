package main

import (
	"fmt"
	"sync"
	"time"
)

// Shared counter that multiple goroutines will access.
var sharedCounter int
var mux sync.Mutex

// Function to increment the shared counter.
func incrementCounter() {
	mux.Lock()         // Acquire lock
	defer mux.Unlock() // Release lock
	sharedCounter++
}

func main() {
	// Create a large number of goroutines to increment the counter.
	numGoroutines := 100000
	for i := 0; i < numGoroutines; i++ {
		go incrementCounter()
	}

	// Wait for all goroutines to complete.
	time.Sleep(5 * time.Second)

	// Print the final value of the shared counter.
	fmt.Printf("Final shared counter value: %d\n", sharedCounter)
}
