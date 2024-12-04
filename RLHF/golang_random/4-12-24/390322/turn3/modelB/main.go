package main

import (
	"context"
	"fmt"
	"time"
)

// Worker simulates a task which can be cancelled
func worker(ctx context.Context, id int, results chan<- string) {
	select {
	case <-time.After(time.Duration(id+1) * time.Second): // Simulate variable completion times
		results <- fmt.Sprintf("Worker %d completed", id)
	case <-ctx.Done(): // Listen for context cancellation
		results <- fmt.Sprintf("Worker %d cancelled", id)
	}
}

func main() {
	// Create a context with cancel and timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure resources are cleaned up after the program exits

	n := 3 // Number of workers
	results := make(chan string, n)

	// Start multiple worker goroutines
	for i := 0; i < n; i++ {
		go worker(ctx, i, results)
	}

	// Collect results or handle task cancellations by respective timeout
	for i := 0; i < n; i++ {
		select {
		case res := <-results:
			fmt.Println(res)
		case <-ctx.Done():
			fmt.Println("Main: Context reached timeout, cancelling remaining tasks if any")
			break
		}
	}

	// Wait a short moment to see cancel messages (this is optional)
	time.Sleep(1 * time.Second)
}
