package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	numGoroutines := 5

	// Add a count to the wait group for each goroutine we're about to start.
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		// Launch a goroutine that will decrement the wait group counter when it finishes.
		go worker(i, &wg)
	}

	// Wait for all goroutines to complete before proceeding.
	wg.Wait()

	fmt.Println("All goroutines completed.")
}

// worker represents the function that will be executed concurrently.
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the wait group counter when the goroutine ends.
	fmt.Printf("Goroutine %d started.\n", id)
	// Simulate work by sleeping for a random duration.
	time.Sleep(time.Duration(id*2) * time.Second)
	fmt.Printf("Goroutine %d finished.\n", id)
}
