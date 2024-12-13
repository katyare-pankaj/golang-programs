package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func doWork(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify that this Goroutine is done when it exits

	for {
		select {
		case <-ctx.Done():
			// Cleanup code here if needed
			fmt.Printf("Goroutine %d cancelled and cleaning up.\n", id)
			return
		default:
			// Simulate work
			fmt.Printf("Goroutine %d is working...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// Create a cancellable context with a signal interrupt listener
	ctx, cancel := context.WithCancel(context.Background())

	// Listen for OS interrupt signals to trigger cancellation
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	// Start multiple Goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go doWork(ctx, i, &wg)
	}

	// Wait for an interrupt signal
	<-sigs
	fmt.Println("Received an interrupt signal, cancelling Goroutines...")

	// Call cancel to signal all Goroutines to stop
	cancel()

	// Wait for all Goroutines to finish
	wg.Wait()
	fmt.Println("All Goroutines have completed.")
}
