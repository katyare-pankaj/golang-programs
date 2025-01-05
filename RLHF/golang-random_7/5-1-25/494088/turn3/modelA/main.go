package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func longRunningTask(ctx context.Context, id int) {
	defer fmt.Printf("Task %d: Exiting\n", id)

	for {
		select {
		case <-ctx.Done():
			// Check the error to see if the context was canceled
			if err := ctx.Err(); err != nil {
				fmt.Printf("Task %d: Canceled by context: %v\n", id, err)
			}
			return
		default:
			// Simulate work being done
			time.Sleep(1 * time.Second)
			fmt.Printf("Task %d: Working...\n", id)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// Create a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Start multiple tasks
	numTasks := 2
	wg.Add(numTasks)
	for i := 0; i < numTasks; i++ {
		go longRunningTask(ctx, i)
	}

	// Wait for tasks to complete or timeout
	wg.Wait()
	fmt.Println("All tasks have finished or timed out.")
}
