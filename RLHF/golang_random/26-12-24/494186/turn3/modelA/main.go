package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Simulates a long-running task
func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker received shutdown signal. Exiting...")
			return
		default:
			fmt.Println("Worker doing work...")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	// Create a wait group to keep track of goroutines
	var wg sync.WaitGroup

	// Create a context with a cancellation function
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start some worker goroutines
	numWorkers := 3
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	// Register a signal handler for SIGINT and SIGTERM
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	defer signal.Stop(sigChan)

	fmt.Println("Program started. Press Ctrl+C to terminate...")

	// Wait for a signal
	select {
	case sig := <-sigChan:
		fmt.Println("Received signal:", sig)

		// Cancel the context, signaling workers to shut down
		cancel()

		// Wait for all worker goroutines to finish
		wg.Wait()

	case <-ctx.Done():
		fmt.Println("Context canceled, exiting...")
	}

	fmt.Println("Program terminated gracefully.")
}
