package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Mock blocking library function
func blockingLibraryOperation(ctx context.Context, id uuid.UUID) (string, error) {
	select {
	case <-ctx.Done():
		return "", errors.Wrap(ctx.Err(), "operation canceled")
	default:
		// Simulate a blocking operation that takes 2 seconds
		time.Sleep(2 * time.Second)
		return fmt.Sprintf("Result for %s", id), nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	id := uuid.New()
	go func() {
		result, err := blockingLibraryOperation(ctx, id)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}
	}()

	// Do other work in the main goroutine while the blocking operation is in progress
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main goroutine doing other work.")
}
