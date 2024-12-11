package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	defer fmt.Printf("Worker %d exiting\n", id)

	for {
		select {
		case <-ctx.Done():
			// When the context is cancelled, exit the Goroutine
			fmt.Printf("Worker %d received cancellation signal\n", id)
			return
		default:
			// Simulate doing some work
			fmt.Printf("Worker %d is working...\n", id)
			time.Sleep(1 * time.Second) // Simulating work
		}
	}
}

func main() {
	// Create a context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())

	// Start a Goroutine with the cancellable context
	go worker(ctx, 1)

	// Let the worker run for a few seconds
	time.Sleep(3 * time.Second)

	// Cancel the context to signal the worker to stop
	fmt.Println("Main: cancelling worker")
	cancel()

	// Wait a little to allow the worker to finish
	time.Sleep(1 * time.Second)

	fmt.Println("Main: exiting")
}
