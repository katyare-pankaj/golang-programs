package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine exits
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Worker %d completed\n", id)
}

func main() {
	var wg sync.WaitGroup // Create a WaitGroup

	// Number of workers we want to spawn
	numWorkers := 5

	// Increment the WaitGroup counter for each worker we are starting
	wg.Add(numWorkers)

	// Spawn goroutines to do the work
	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg)
	}

	// Wait for all workers to complete their tasks
	wg.Wait()

	fmt.Println("All workers completed. Main goroutine exiting.")
}
