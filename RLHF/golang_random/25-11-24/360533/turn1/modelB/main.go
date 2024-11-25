package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a background context
	ctx := context.Background()
	// Add a value to the context
	ctx = context.WithValue(ctx, "userID", 123)
	// Create a cancelable context
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Start the function with the context
	go doWork(ctx)
	// Simulate work being done
	time.Sleep(2 * time.Second)
	// Cancel the context
	cancel()
	fmt.Println("Main function exiting")
}
func doWork(ctx context.Context) {
	// Get the value from the context
	userID := ctx.Value("userID").(int)
	fmt.Println("Worker started with user ID:", userID)
	// Simulate work being done
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker canceled")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Worker doing work...")
		}
	}
}
