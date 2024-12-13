package main

import (
	"fmt"
	"sync"
)

// Shared counter variable
var sharedCounter int

func incrementCounter(n int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the function completes
	for i := 0; i < n; i++ {
		sharedCounter++
	}
}

func main() {
	const numGoroutines = 10
	const iterationsPerGoroutine = 1_000_000

	var wg sync.WaitGroup
	fmt.Println("Starting with sharedCounter:", sharedCounter)

	// Start Goroutines to increment the counter
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1) // Increment the WaitGroup counter before starting the Goroutine
		go incrementCounter(iterationsPerGoroutine, &wg)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	fmt.Println("Final value of sharedCounter:", sharedCounter)
}
