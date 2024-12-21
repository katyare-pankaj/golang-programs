package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Cancelled\n", id)
			return
		default:
			// Simulate work
			time.Sleep(time.Second)
			fmt.Printf("Worker %d is doing work...\n", id)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// Start multiple workers with a time-bound context
	for i := 0; i < 3; i++ {
		workerCtx, _ := context.WithTimeout(ctx, time.Second*5)
		go worker(workerCtx, i)
	}

	fmt.Println("Main routine is working...")

	time.Sleep(time.Second * 7)

	// After 7 seconds, cancel all workers
	fmt.Println("Cancelling all workers...")
	cancel()

	// Wait for all workers to finish
	time.Sleep(time.Second * 2)

	fmt.Println("All workers terminated. Exiting.")
}
