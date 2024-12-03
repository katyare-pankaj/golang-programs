package main

import (
	"errors"
	"fmt"
	"sync"
)

// Define a type for Task, which is a function returning an error
type Task func() error

// Function to execute tasks concurrently and collect errors
func RunTasksConcurrently(tasks []Task) []error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(tasks))

	// Worker function to run a task and send the error to a channel
	runTask := func(task Task) {
		defer wg.Done()
		errChan <- task() // Send the result to the channel
	}

	// Start Goroutines for each task
	for _, task := range tasks {
		wg.Add(1)
		go runTask(task)
	}

	// Close the channel once all tasks have finished
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Collect errors from the channel
	var errors []error
	for err := range errChan {
		// Only append non-nil errors to filter out successful results
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func main() {
	// Define a list of tasks
	tasks := []Task{
		func() error {
			return nil // simulate a successful task
		},
		func() error {
			return errors.New("error in task 2") // simulate an error
		},
		func() error {
			return nil // simulate a successful task
		},
		func() error {
			return errors.New("error in task 4") // simulate another error
		},
	}

	// Execute tasks concurrently and handle errors
	errors := RunTasksConcurrently(tasks)

	// Output errors
	if len(errors) > 0 {
		fmt.Println("Errors occurred:")
		for i, err := range errors {
			fmt.Printf("Task %d: %v\n", i+1, err)
		}
	} else {
		fmt.Println("All tasks completed successfully.")
	}
}
