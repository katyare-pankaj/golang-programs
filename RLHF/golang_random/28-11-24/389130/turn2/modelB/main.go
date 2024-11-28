package main

import (
	"errors"
	"fmt"
)

// Custom error type with context
type MyError struct {
	Message string
	Context map[string]string
}

func (e MyError) Error() string {
	return e.Message
}

func doSomething() error {
	// Simulate an error
	if err := someOperation(); err != nil {
		return fmt.Errorf("something failed: %w", err)
	}
	return nil
}

func someOperation() error {
	return errors.New("operation failed due to network error")
}

func main() {
	err := doSomething()
	if err != nil {
		// Handle the error
		if varErr, ok := err.(*MyError); ok {
			fmt.Printf("Error: %s\n", varErr.Message)
			for key, value := range varErr.Context {
				fmt.Printf("Context: %s = %s\n", key, value)
			}
		} else {
			fmt.Printf("Error: %s\n", err)
		}
	}
}
