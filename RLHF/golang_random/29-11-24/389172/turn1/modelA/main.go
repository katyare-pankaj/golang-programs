package main

import (
	"context"
	"fmt"
	"time"
)

// Simulate a long-running task
func longRunningTask(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Task cancelled due to timeout")
		return
	case <-time.After(10 * time.Second):
	}
	fmt.Println("Task completed successfully")
}

func main() {
	// Create a context with a 5-second deadline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Start the goroutine with the context
	go longRunningTask(ctx)

	// Wait for the goroutine to finish or the timeout to trigger
	select {
	case <-time.After(20 * time.Second):
		fmt.Println("Task did not finish within the given time")
	}
}
