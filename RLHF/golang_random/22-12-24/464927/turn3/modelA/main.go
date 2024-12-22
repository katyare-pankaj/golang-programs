package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	// Create a background context
	backgroundCtx := context.Background()

	// Create a cancelable context
	ctx, cancel := context.WithCancel(backgroundCtx)
	defer cancel() // Cancel the context on program exit

	// Set up a WaitGroup to wait for goroutines
	var wg sync.WaitGroup

	// Add the goroutines to the WaitGroup
	wg.Add(2)

	// Start the first goroutine
	go func(ctx context.Context) {
		defer wg.Done() // Decrement the counter once the goroutine finishes
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 1 received cancel signal, exiting")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Goroutine 1 is working")
			}
		}
	}(ctx)

	// Start the second goroutine
	go func(ctx context.Context) {
		defer wg.Done() // Decrement the counter once the goroutine finishes
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 2 received cancel signal, exiting")
				return
			case <-time.After(2 * time.Second):
				fmt.Println("Goroutine 2 is working")
			}
		}
	}(ctx)

	// Handle termination signals
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	for sig := range ch {
		fmt.Printf("Received signal: %s, cancelling context\n", sig)
		cancel()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines finished, program exiting")
}
