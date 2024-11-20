package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const (
	maxRetries    = 3
	initialDelay  = 100 * time.Millisecond
	backoffFactor = 2.0
)

func retry(ctx context.Context, f func(context.Context) error) error {
	for attempt := 0; attempt < maxRetries; attempt++ {
		err := f(ctx)
		if err == nil {
			return nil
		}

		// Calculate the delay for the next attempt
		delay := initialDelay * time.Duration(rand.Float64()*backoffFactor)
		fmt.Println("Attempt", attempt+1, "failed, retrying in", delay)

		// Check if the context is canceled
		select {
		case <-ctx.Done():
			fmt.Println("Context canceled, stopping retry")
			return ctx.Err()
		case <-time.After(delay):
		}
	}
	return fmt.Errorf("all retries failed")
}

func externalService(ctx context.Context) error {
	// Simulate a transient error with a certain probability
	if rand.Intn(10) < 4 {
		return fmt.Errorf("external service failed transiently")
	}
	fmt.Println("External service successful")
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := retry(ctx, externalService)
	if err != nil {
		fmt.Println("Final error:", err)
	} else {
		fmt.Println("Operation successful")
	}
}
