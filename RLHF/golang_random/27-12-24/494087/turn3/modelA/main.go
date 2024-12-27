package main

import (
	"fmt"
	"sync"
	"time"
)

// Function that represents a task to be performed by a goroutine
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when this goroutine finishes
	fmt.Printf("Worker %d started.\n", id)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Worker %d finished.\n", id)
}

func main() {
	var wg sync.WaitGroup // Initialize the WaitGroup
	numWorkers := 5

	// Start the workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1) // Increment the counter for each worker
		go worker(i+1, &wg)
	}

	// Wait for all workers to complete
	fmt.Println("Main function is waiting for workers to finish...")
	wg.Wait() // Block the main function until all workers call wg.Done()

	fmt.Println("All workers have completed. Main function continues...")
}
