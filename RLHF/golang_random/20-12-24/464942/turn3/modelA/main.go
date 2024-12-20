package main

import (
	"fmt"
	"sync"
)

func incrementCounter(counter *int, iterations int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < iterations; i++ {
		// Lock the mutex before accessing the shared counter
		counterLock.Lock()
		*counter++
		counterLock.Unlock()
	}
}

func main() {
	var counter = 0
	var wg sync.WaitGroup
	var iterations = 100000 // Number of iterations per goroutine
	var numGoroutines = 10  // Number of goroutines

	// Initialize the mutex
	counterLock := &sync.Mutex{}

	// Add a task for each goroutine
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementCounter(&counter, iterations, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Verify the final counter value
	expectedCounterValue := numGoroutines * iterations
	if counter != expectedCounterValue {
		fmt.Printf("Mismatch: expected %d, got %d\n", expectedCounterValue, counter)
	} else {
		fmt.Printf("The final counter value is correct: %d\n", counter)
	}
}
