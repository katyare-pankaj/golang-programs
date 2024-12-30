package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure that wg.Done() is called even if an error occurs

	// Simulate work by sleeping for a random duration
	duration := time.Duration(100+id*50) * time.Millisecond
	fmt.Printf("Worker %d starting, will sleep for %s\n", id, duration)
	time.Sleep(duration)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 5

	// Add to the WaitGroup for each worker goroutine
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, &wg)
	}

	// Wait for all goroutines to call wg.Done()
	wg.Wait()

	fmt.Println("All workers have completed their tasks. Exiting main function.")
}
