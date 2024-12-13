package main

import (
	"fmt"
	"sync"
	"time"
)

// Shared counter variable
var sharedCounter int

// Mutex to protect access to sharedCounter
var mutex sync.Mutex

func incrementCounter(n int) {
	for i := 0; i < n; i++ {
		mutex.Lock() // Acquire the lock
		sharedCounter++
		mutex.Unlock() // Release the lock
	}
}

func main() {
	const numGoroutines = 10
	const iterationsPerGoroutine = 1_000_000

	fmt.Println("Starting with sharedCounter:", sharedCounter)

	// Start Goroutines to increment the counter
	for i := 0; i < numGoroutines; i++ {
		go incrementCounter(iterationsPerGoroutine)
	}

	// Wait for all Goroutines to finish
	time.Sleep(2 * time.Second)

	fmt.Println("Final value of sharedCounter:", sharedCounter)
}
