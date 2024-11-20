package main

import (
	"context"
	"fmt"
	"time"
)

// Simulate a blocking operation from a third-party library
func blockingOperation(ctx context.Context, duration time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(duration):
		return nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Attempt the blocking operation
	err := blockingOperation(ctx, 5*time.Second)
	if err != nil {
		fmt.Printf("Blocking operation timed out: %v\n", err)
	} else {
		fmt.Println("Blocking operation completed successfully.")
	}
}
