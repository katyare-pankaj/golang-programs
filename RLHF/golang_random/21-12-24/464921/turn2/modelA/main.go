package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Worker simulates a worker task.
func Worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d canceled.\n", id)
			return
		default:
			fmt.Printf("Worker %d working...\n", id)
			time.Sleep(time.Second) // Simulate work
		}
	}
}

func main() {
	// Initialize a wait group to track workers
	var wg sync.WaitGroup

	// Create a background context
	ctx := context.Background()

	// Create a channel to receive user signals
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	// Start multiple worker tasks
	numWorkers := 3
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Worker(ctx, i, &wg)
	}

	// Wait for a signal or all workers to finish
	select {
	case sig := <-ch:
		fmt.Printf("Received signal: %s\n", sig)
		// Cancel the context to stop the worker tasks
		ctx, cancel := context.WithCancel(ctx)
		cancel()
	case <-wg.Done():
		fmt.Println("All workers finished.")
	}

	// Wait for all workers to stop
	wg.Wait()
}
