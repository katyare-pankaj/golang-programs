package main

import (
	"fmt"
	"sync"
)

// A function that simulates a worker goroutine.
func worker(id int, task string, done chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate a failure condition.
	if task == "failingTask" {
		done <- fmt.Errorf("task %d failed: %s", id, task)
	} else {
		done <- nil
	}
}

func main() {
	var wg sync.WaitGroup
	done := make(chan error, 3)

	// Launch three worker goroutines.
	wg.Add(3)
	go worker(1, "succeedingTask", done, &wg)
	go worker(2, "failingTask", done, &wg)
	go worker(3, "succeedingTask", done, &wg)

	// Wait for all workers to complete.
	wg.Wait()

	// Collect errors from the channel.
	var errors []error
	for err := range done {
		if err != nil {
			errors = append(errors, err)
		}
	}

	// Print out any collected errors.
	if len(errors) > 0 {
		fmt.Println("Errors occurred:")
		for _, err := range errors {
			fmt.Println(err)
		}
	} else {
		fmt.Println("All tasks completed successfully.")
	}
}
