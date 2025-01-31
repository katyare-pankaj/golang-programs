package main

import (
	"context"
	"fmt"
	"time"
)

// Callback is a functional interface type for processing tasks
type Callback func(ctx context.Context, id int) error

// worker is a goroutine function that processes a task with a given Callback
func worker(ctx context.Context, id int, callback Callback) {
	select {
	case <-ctx.Done():
		fmt.Printf("Worker %d: canceled\n", id)
	default:
		err := callback(ctx, id)
		if err != nil {
			fmt.Printf("Worker %d: error %v\n", id, err)
		} else {
			fmt.Printf("Worker %d: completed successfully\n", id)
		}
	}
}

func main() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Define a callback function
	callback := func(ctx context.Context, id int) error {
		// Simulate some work
		fmt.Printf("Worker %d: working...\n", id)
		select {
		case <-time.After(1 * time.Second):
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	// Start multiple workers
	for i := 1; i <= 5; i++ {
		go worker(ctx, i, callback)
	}

	// Wait some time to observe the outputs
	time.Sleep(4 * time.Second)
}
