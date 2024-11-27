package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Call Done() when the goroutine exits
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Duration(id) * time.Second) // Simulate work
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	numWorkers := 5

	// Add the number of goroutines to the WaitGroup
	wg.Add(numWorkers)

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All workers have completed. Main function exiting.")
}
