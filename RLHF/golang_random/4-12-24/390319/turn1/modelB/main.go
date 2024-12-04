package main

import (
	"context"
	"fmt"
	"time"
)

// doWork simulates a long-running task
func doWork(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			// Context has been canceled
			fmt.Printf("Worker %d: Stopping work due to cancellation\n", id)
			return
		default:
			// Simulate doing some work
			fmt.Printf("Worker %d: Doing work...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Create a context that can be canceled
	ctx, cancel := context.WithCancel(context.Background())

	// Start a goroutine to do some work
	go doWork(ctx, 1)

	// Simulate some other work in main goroutine
	time.Sleep(3 * time.Second)

	// Cancel the context, which will signal the goroutine to stop
	fmt.Println("Main: Canceling context...")
	cancel()

	// Give the goroutine time to finish
	time.Sleep(1 * time.Second)

	fmt.Println("Main: Done")
}
