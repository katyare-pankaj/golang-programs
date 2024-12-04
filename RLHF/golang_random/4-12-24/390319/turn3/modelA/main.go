package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// task simulates a task that can be canceled
func task(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Task %d: Canceling due to context error: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Task %d: Working...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// Create a context with a 2-second deadline
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Start multiple tasks
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go task(ctx, &wg, i)
	}

	// Simulate waiting for a while
	fmt.Println("Waiting for tasks...")
	wg.Wait()
}
