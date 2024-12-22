package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Duration(id) * time.Second) // Simulate work
	fmt.Printf("Worker %d completed\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Number of workers to spawn
	numWorkers := 3

	// Add the main goroutine and all workers to the WaitGroup
	wg.Add(numWorkers + 1)

	// Start the workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg)
	}

	// Do some main task here
	fmt.Println("Main task is doing some work...")

	// Wait for all workers to finish before proceeding
	wg.Wait()

	fmt.Println("All workers have finished. Exiting main function...")
}
