package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// worker function that takes a context and a WaitGroup as parameters
func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done() // Call Done on WaitGroup when goroutine exits
	for {
		select {
		case <-ctx.Done():
			// Check if context is canceled
			fmt.Printf("Worker %d: Exiting due to cancellation.\n", id)
			return
		default:
			// Simulate work by printing a message and sleeping
			fmt.Printf("Worker %d: Performing task...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	const numWorkers = 5  // Number of workers to create
	var wg sync.WaitGroup // Create a WaitGroup to track worker lifecycles

	// Create a parent context with a time limit of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Call cancel to release resources when main function exits

	// Add the number of workers to the WaitGroup
	wg.Add(numWorkers)

	// Create worker Goroutines with the context and WaitGroup
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, &wg, i)
	}

	// Wait for all worker Goroutines to complete or until the context is canceled
	wg.Wait()
	fmt.Println("All workers have completed.")
}
