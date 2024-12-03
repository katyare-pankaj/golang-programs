package main

import (
	"fmt"
	"sync"
	"time"
)

// Define a type for the task, which is a function that returns an error
type Task func() error

// Higher-order function to run tasks concurrently using Goroutines
func RunConcurrently(tasks []Task) []error {
	var wg sync.WaitGroup
	errors := make([]error, len(tasks))

	// Function to run each task and capture its error
	taskRunner := func(i int, task Task) {
		defer wg.Done()
		errors[i] = task()
	}

	// Start each task in a new Goroutine
	for i, task := range tasks {
		wg.Add(1)
		go taskRunner(i, task)
	}

	// Wait for all Goroutines to finish
	wg.Wait()
	return errors
}

func main() {
	// Define some tasks to run
	tasks := []Task{
		func() error {
			fmt.Println("Task 1 is running")
			time.Sleep(1 * time.Second)
			return nil // simulate success
		},
		func() error {
			fmt.Println("Task 2 is running")
			time.Sleep(2 * time.Second)
			return fmt.Errorf("Task 2 failed") // simulate failure
		},
		func() error {
			fmt.Println("Task 3 is running")
			time.Sleep(1 * time.Second)
			return nil // simulate success
		},
	}

	// Run tasks concurrently and capture results
	errors := RunConcurrently(tasks)

	// Display results
	for i, err := range errors {
		if err != nil {
			fmt.Printf("Error in task %d: %v\n", i+1, err)
		} else {
			fmt.Printf("Task %d completed successfully\n", i+1)
		}
	}
}
