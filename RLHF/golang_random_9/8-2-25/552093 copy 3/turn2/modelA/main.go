package main

import (
	"fmt"
	"sync"
	"time"
)

// Function that represents a task to be done by a goroutine
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrements the counter when the goroutine completes

	fmt.Printf("Worker %d starting\n", id)

	// Simulate some work with sleep
	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// Create a WaitGroup
	var wg sync.WaitGroup

	// Define the number of goroutines
	numWorkers := 3

	// Launch goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Increment the WaitGroup counter

		// Start a worker goroutine
		go worker(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All workers completed")
}
