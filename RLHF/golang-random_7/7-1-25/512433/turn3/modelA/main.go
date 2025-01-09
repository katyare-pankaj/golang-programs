package main

import (
	"fmt"
	"sync"
	"time"
)

func simulateWork(id int, wg *sync.WaitGroup, errorCh chan error, workDuration time.Duration) {
	defer wg.Done()
	fmt.Printf("Worker %d is starting and will work for %s\n", id, workDuration)
	time.Sleep(workDuration) // Simulate work by sleeping

	// Introduce a random error for demonstration purposes
	if workDuration%2 == 0 {
		errorCh <- fmt.Errorf("worker %d encountered an error", id)
	}

	fmt.Printf("Worker %d is done\n", id)
}

func main() {
	var wg sync.WaitGroup           // Initialize the wait group
	errorCh := make(chan error, 10) // Create a channel to pass errors

	workDurations := []time.Duration{
		time.Second * 1,
		time.Second * 2,
		time.Second * 3,
		time.Second * 4,
		time.Second * 5,
	}

	// Add workers to the wait group
	for i, duration := range workDurations {
		wg.Add(1)                                  // Increment the wait count for each worker
		go simulateWork(i, &wg, errorCh, duration) // Start the worker goroutine
	}

	// Wait for all workers to finish
	fmt.Println("Main goroutine is waiting for workers...")
	wg.Wait() // Block until the wait count reaches zero

	// Collect and print errors
	close(errorCh) // Close the channel after all workers have finished
	var errors []error
	for err := range errorCh {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		fmt.Println("Errors encountered:")
		for _, err := range errors {
			fmt.Printf("- %v\n", err)
		}
	} else {
		fmt.Println("All workers completed without errors.")
	}
}
