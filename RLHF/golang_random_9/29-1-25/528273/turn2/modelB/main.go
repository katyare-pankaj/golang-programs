package main

import (
	"fmt"
	"sync"
	"time"
)

func optimizedWorker(id int, wg *sync.WaitGroup) {
	// Define a cleanup function
	cleanup := func() {
		wg.Done()
	}

	// Defer the cleanup function once outside the loop
	defer cleanup()

	fmt.Printf("Optimized Worker %d: Starting\n", id)

	// Simulate work with a loop and multiple tasks.
	for i := 0; i < 5; i++ {
		// Some actual work is done here...
		fmt.Printf("Worker %d: Processing task %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("Optimized Worker %d: Finished\n", id)
}

func main() {
	const numWorkers = 10
	var wg sync.WaitGroup

	wg.Add(numWorkers)
	fmt.Println("Starting optimized workers:")
	for i := 0; i < numWorkers; i++ {
		go optimizedWorker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All optimized workers finished.")
}
