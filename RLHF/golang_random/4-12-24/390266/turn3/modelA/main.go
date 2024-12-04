package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Retry retries the given function a fixed number of times with a constant delay between retries.
// It returns the result of the function if it succeeds, otherwise it returns an error.
func Retry(maxRetries int, delay time.Duration, f func() error) error {
	for retries := 0; retries < maxRetries; retries++ {
		if err := f(); err == nil {
			return nil // Success
		}

		fmt.Printf("Retry %d: Failed, waiting %v before retrying...\n", retries+1, delay)
		time.Sleep(delay)
	}

	return fmt.Errorf("Maximum retries reached")
}

// Simulate a function that may fail
func simulateFailingOperation() error {
	// Randomly simulate a failure
	if rand.Float32() < 0.3 {
		return fmt.Errorf("Operation failed")
	}
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Set maximum retries and delay
	maxRetries := 3
	delay := 1 * time.Second

	// Retry the operation
	err := Retry(maxRetries, delay, simulateFailingOperation)
	if err != nil {
		fmt.Printf("Operation failed after %d retries: %v\n", maxRetries, err)
	} else {
		fmt.Println("Operation succeeded after retry.")
	}
}
