package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Custom error types
type NetworkError struct {
	Message string
}

func (e *NetworkError) Error() string {
	return e.Message
}

type DatabaseError struct {
	Message string
}

func (e *DatabaseError) Error() string {
	return e.Message
}

// Define a type for Task, which is a function returning an error
type Task func() error

// Function to execute tasks concurrently and handle errors
func RunTasksConcurrently(tasks []Task) {
	var wg sync.WaitGroup
	errChan := make(chan error, len(tasks))

	// Worker function to run a task and send any error to the channel
	runTask := func(task Task) {
		defer wg.Done()
		if err := task(); err != nil {
			errChan <- err
		}
	}

	// Start each task in a new Goroutine
	for _, task := range tasks {
		wg.Add(1)
		go runTask(task)
	}

	// Close the channel once all tasks have finished
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Process errors from the channel
	for err := range errChan {
		switch e := err.(type) {
		case *NetworkError:
			fmt.Printf("Network error occurred: %s\n", e.Message)
		case *DatabaseError:
			fmt.Printf("Database error occurred: %s\n", e.Message)
		default:
			fmt.Printf("Other error occurred: %s\n", e.Error())
		}
	}
}

func main() {
	// Define tasks with different error responses
	tasks := []Task{
		func() error {
			time.Sleep(1 * time.Second)
			return nil // simulate success
		},
		func() error {
			time.Sleep(1 * time.Second)
			return &NetworkError{"Network timeout"} // simulate network error
		},
		func() error {
			time.Sleep(2 * time.Second)
			return &DatabaseError{"Connection lost"} // simulate database error
		},
		func() error {
			time.Sleep(1 * time.Second)
			return errors.New("Unexpected error") // simulate general error
		},
	}

	// Run tasks concurrently
	RunTasksConcurrently(tasks)
}
