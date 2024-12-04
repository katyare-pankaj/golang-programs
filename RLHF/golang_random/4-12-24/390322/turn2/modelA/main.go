package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Worker simulates a task that takes some time to complete.
func worker(ctx context.Context, id int) {
	defer fmt.Println("Worker", id, "finished")
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Worker", id, "completed task")
	case <-ctx.Done():
		fmt.Println("Worker", id, "cancelled")
	}
}

func main() {
	// Number of workers to start
	numWorkers := 5

	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Ensure context is canceled even if the main exits early

	var wg sync.WaitGroup

	// Start multiple worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(ctx, id)
		}(i)
	}

	// Wait for workers to finish, but handle cancellation if it occurs
	select {
	case <-wg.Done():
		fmt.Println("All workers completed successfully")
	case <-ctx.Done():
		fmt.Println("Context timed out, cancelling workers")
		cancel()
		wg.Wait() // Ensure all goroutines are done
	}
}
