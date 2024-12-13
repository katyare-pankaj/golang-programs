package main

import (
	"fmt"
	"sync"
)

// Shared counter variable
var sharedCounter int

// Mutex to protect access to sharedCounter
var mutex sync.Mutex

func incrementCounter(n int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure that Done is called at the end of the function

	for i := 0; i < n; i++ {
		mutex.Lock() // Acquire the lock
		sharedCounter++
		mutex.Unlock() // Release the lock
	}
}

func main() {
	const numGoroutines = 10
	const iterationsPerGoroutine = 1_000_000

	var wg sync.WaitGroup // Create a WaitGroup instance

	fmt.Println("Starting with sharedCounter:", sharedCounter)

	// Start Goroutines to increment the counter
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)                                        // Increment WaitGroup counter
		go incrementCounter(iterationsPerGoroutine, &wg) // Start Goroutine with WaitGroup
	}

	wg.Wait() // Wait for all Goroutines to finish

	fmt.Println("Final value of sharedCounter:", sharedCounter)
}
