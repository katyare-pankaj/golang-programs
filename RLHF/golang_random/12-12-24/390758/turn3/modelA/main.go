package main

import (
	"fmt"
	"time"
)

func main() {
	// Create channels
	resultChannel := make(chan string)
	timeoutChannel := make(chan struct{})

	// Simulate an operation in a goroutine
	go func() {
		// Simulate a long-running operation
		time.Sleep(3 * time.Second) // Simulating workload
		resultChannel <- "Operation completed successfully"
	}()

	// Set a timeout for 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		close(timeoutChannel) // Signal timeout
	}()

	// Wait for the result or timeout
	select {
	case result := <-resultChannel:
		fmt.Println(result)
	case <-timeoutChannel:
		fmt.Println("Operation timed out.")
	}
}
