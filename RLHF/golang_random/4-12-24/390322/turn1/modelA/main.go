package main

import (
	"context"
	"fmt"
	"time"
)

// simulateWork simulates some work that may take time to complete.
func simulateWork(ctx context.Context, name string) {
	// Simulate some work taking time
	time.Sleep(2 * time.Second)

	select {
	case <-ctx.Done():
		fmt.Printf("Work for %s cancelled\n", name)
	default:
		fmt.Printf("Work for %s completed\n", name)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Start the work with the provided context
	go simulateWork(ctx, "Task 1")

	// Start another task (as an example, to illustrate more concurrent use)
	go simulateWork(ctx, "Task 2")

	time.Sleep(1.5 * time.Second) // Wait a bit to see the work progress

	// Cancel the context to stop any ongoing work
	cancel()

	fmt.Println("Main thread finished")
}
