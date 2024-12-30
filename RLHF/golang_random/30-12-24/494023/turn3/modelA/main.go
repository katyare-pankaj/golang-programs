package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// Simulate some work being done
	defer wg.Done()
	fmt.Printf("Worker %d is working...\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d has completed its work.\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Create a predefined number of workers
	numWorkers := 5

	// Add the number of workers to the WaitGroup
	wg.Add(numWorkers)

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	// Main program proceeds only after all workers are done
	fmt.Println("All workers have completed their tasks. Main program is exiting.")
}
