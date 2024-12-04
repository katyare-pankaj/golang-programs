package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// simulateOperation simulates an operation that might fail randomly.
func simulateOperation() error {
	// Simulate a random failure
	if rand.Float32() < 0.3 { // 30% chance of failure
		return errors.New("operation failed")
	}
	return nil
}

// retryWithFixedBackoffRetries an operation with retries on failure and a fixed delay.
func retryWithFixedBackoff(operation func() error, retries int, delay time.Duration) error {
	var err error
	for i := 0; i < retries; i++ {
		err = operation()
		if err == nil {
			return nil
		}
		fmt.Printf("Attempt %d: operation failed with error: %v. Retrying...\n", i+1, err)
		time.Sleep(delay)
	}
	return fmt.Errorf("operation failed after %d attempts: %s", retries, err)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed random number generator

	err := retryWithFixedBackoff(simulateOperation, 5, 2*time.Second)
	if err != nil {
		fmt.Println("Final outcome: error")
		fmt.Printf("Operation failed after retries: %v\n", err)
	} else {
		fmt.Println("Final outcome: success")
	}
}
