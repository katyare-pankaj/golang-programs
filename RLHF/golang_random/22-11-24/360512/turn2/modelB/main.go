package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// goroutineFunction represents the function to be executed in a goroutine.
type goroutineFunction func(ctx context.Context) error

func main() {
	// Define the number of goroutines
	numGoroutines := 3

	// Create a WaitGroup to track the number of running goroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Create a background context
	ctx := context.Background()

	// Define a slice of functions to run in goroutines
	functions := []goroutineFunction{
		func(ctx context.Context) error {
			defer wg.Done()
			return work(ctx, "goroutine 1")
		},
		func(ctx context.Context) error {
			defer wg.Done()
			return work(ctx, "goroutine 2")
		},
		func(ctx context.Context) error {
			defer wg.Done()
			return work(ctx, "goroutine 3")
		},
	}

	// Launch the goroutines with a timeout of 2 seconds
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	for _, f := range functions {
		go func(f goroutineFunction) {
			err := f(timeoutCtx)
			if err != nil {
				// Handle the error (if any)
				fmt.Printf("Error: %v\n", err)
			}
		}(f)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines completed.")
}

func work(ctx context.Context, name string) error {
	// Simulate work with a random duration
	duration := time.Duration(rand.Intn(5)) * time.Second

	fmt.Printf("%s starting with duration %s...\n", name, duration)

	// Use the context for timeout handling
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			// Check if the context is canceled
			return ctx.Err()
		case <-ticker.C:
			// Perform work
			time.Sleep(1 * time.Second)
			duration -= time.Second
			if duration <= 0 {
				fmt.Printf("%s completed.\n", name)
				return nil
			}
		}
	}
}
