package main

import (
	"context"
	"fmt"
	"time"
)

// worker simulates a task that takes variable time
func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Stopping due to cancellation\n", id)
			return
		default:
			// Simulating work by sleeping
			fmt.Printf("Worker %d: Working...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// Create a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Start multiple workers (goroutines)
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, i)
	}

	// Wait and then cancel the context manually (though it will timeout on its own)
	time.Sleep(5 * time.Second) // In a real scenario, you might decide to cancel earlier
	fmt.Println("Main: Cancellation complete")

	// Adding a small sleep to ensure output display
	time.Sleep(time.Second)
}
