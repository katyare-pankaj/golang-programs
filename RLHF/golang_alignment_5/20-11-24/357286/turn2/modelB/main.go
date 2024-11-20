package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	maxRetries  = 3                      // Maximum number of retries
	maxBackoff  = 5 * time.Second        // Maximum backoff duration
	baseBackoff = 250 * time.Millisecond // Base backoff duration
)

// retryableError is an interface that returns true if an error is transient and should be retried.
type retryableError interface {
	IsTransient() bool
}

// TransientError is a basic struct that implements retryableError.
type TransientError struct {
	error
}

// IsTransient returns true, indicating that this error is transient.
func (e TransientError) IsTransient() bool {
	return true
}

// retry executes the provided function up to maxRetries times, with exponential backoff on transient errors.
func retry(f func() error) error {
	for i := 0; i <= maxRetries; i++ {
		err := f()
		if err == nil {
			return nil
		}

		if re, ok := err.(retryableError); ok && re.IsTransient() {
			fmt.Printf("Transient error: %v, Retrying in...\n", err)

			// Calculate the backoff duration using exponential backoff with a jitter to avoid retrying all at once
			backoff := baseBackoff * time.Duration(2*i+rand.Intn(i+1))
			if backoff > maxBackoff {
				backoff = maxBackoff
			}

			time.Sleep(backoff)
			continue
		}

		// If the error is not transient, return it
		return err
	}
	return fmt.Errorf("all retries failed")
}

// Function with a transient error
func someOperation() error {
	// Simulate a transient error with a probability of 50%
	if rand.Intn(2) == 0 {
		return TransientError{fmt.Errorf("some transient error occurred")}
	}
	return nil
}

func main() {
	err := retry(someOperation)
	if err != nil {
		fmt.Println("Failed:", err)
	} else {
		fmt.Println("Success!")
	}
}
