package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Exiting due to context cancellation\n", id)
			return
		default:
			// Simulate work being done
			time.Sleep(time.Millisecond * 100)
			fmt.Printf("Worker %d: Working...\n", id)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Start multiple workers
	numWorkers := 5
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(ctx, &wg, i)
	}

	// Wait for all workers to complete or timeout
	wg.Wait()
	fmt.Println("All workers have finished or timed out.")
}
