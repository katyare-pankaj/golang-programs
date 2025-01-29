package main

import (
	"fmt"
	"sync"
	"time"
)

// ProcessItems simulates processing multiple items.
func ProcessItems(id int, wg *sync.WaitGroup, items int) {
	defer wg.Done()

	fmt.Printf("Worker %d: Starting\n", id)

	// Define cleanup logic once.
	cleanup := func() {
		fmt.Printf("Worker %d: Cleaning up resources\n", id)
		// Place cleanup logic here, like closing files, connections, etc.
	}

	// Perform work with items, defer just once for the cleanup.
	defer cleanup() // Defer called outside the loop.

	for i := 0; i < items; i++ {
		// Simulate work and resource management.
		fmt.Printf("Worker %d: Processing item %d\n", id, i)
		time.Sleep(10 * time.Millisecond)
		// Normally, a task might open a resource here.
	}

	fmt.Printf("Worker %d: Finished\n", id)
}

func main() {
	const (
		numWorkers = 5
		numItems   = 100
	)
	var wg sync.WaitGroup

	wg.Add(numWorkers)
	fmt.Println("Starting optimized workers:")
	for i := 0; i < numWorkers; i++ {
		go ProcessItems(i, &wg, numItems)
	}
	wg.Wait()
	fmt.Println("All optimized workers finished.")
}
