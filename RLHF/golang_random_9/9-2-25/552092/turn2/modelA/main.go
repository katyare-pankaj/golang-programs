package main

import (
	"fmt"
	"sync"
	"time"
)

// simulateWork simulates a task taking some time to complete
func simulateWork(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure Done is called when the goroutine completes

	fmt.Printf("Goroutine %d is starting work\n", id)
	time.Sleep(time.Second) // Simulate work by sleeping for 1 second
	fmt.Printf("Goroutine %d has finished work\n", id)
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 3

	// Define how many goroutines we're waiting for
	wg.Add(numGoroutines)

	for i := 1; i <= numGoroutines; i++ {
		// Launch goroutines
		go simulateWork(i, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All goroutines have completed")
}
