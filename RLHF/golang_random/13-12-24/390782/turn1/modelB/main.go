package main

import (
	"fmt"
	"time"
)

// Simulates a function that might encounter an error
func doWork(id int, workChannel chan<- string, errorChannel chan<- error) {
	defer close(errorChannel)

	time.Sleep(time.Duration(id) * time.Second)

	// Simulate an error
	if id%2 == 0 {
		errorChannel <- fmt.Errorf("error in goroutine %d: failed to complete work", id)
		return
	}

	// Send work completion message
	workChannel <- fmt.Sprintf("goroutine %d completed work", id)
}

func main() {
	workChannel := make(chan string)
	errorChannel := make(chan error)
	numGoroutines := 5

	// Launch goroutines
	for i := 0; i < numGoroutines; i++ {
		go doWork(i, workChannel, errorChannel)
	}

	// Wait for all goroutines to finish and collect errors
	errorsOccurred := false
	for i := 0; i < numGoroutines; i++ {
		select {
		case workResult := <-workChannel:
			fmt.Println(workResult)
		case err := <-errorChannel:
			fmt.Printf("Error: %v\n", err)
			errorsOccurred = true
		}
	}

	// Check if any errors occurred
	if errorsOccurred {
		fmt.Println("Some errors occurred during work execution.")
	} else {
		fmt.Println("All work completed successfully.")
	}
}
