package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker", id, "cancelled")
			return
		default:
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println("worker", id, "working")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// Start some workers
	for i := 0; i < 5; i++ {
		wg.Add(1) // Increment the wait group counter
		go worker(ctx, &wg, i)
	}

	// Simulate some work that might take a long time in the main thread
	time.Sleep(2 * time.Second)

	// Cancel the context after 2 seconds
	cancel()

	// Wait for all workers to finish their current iteration
	wg.Wait()

	// Proceed to the next stage of execution
	fmt.Println("All workers have finished. Proceeding to next stage.")
}
