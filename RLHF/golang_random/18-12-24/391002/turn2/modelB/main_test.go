package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// Worker processes tasks and sends errors to the error channel.
func Worker(id int, tasks <-chan int, errs chan<- error) {
	for task := range tasks {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Simulate work
		if task%3 == 0 {
			errs <- fmt.Errorf("worker %d: failed to process task %d", id, task)
			continue
		}
		fmt.Printf("Worker %d completed task %d\n", id, task)
	}
	close(errs)
}

// CentralErrorHandler receives errors from multiple workers and accumulates them.
func CentralErrorHandler(errs <-chan error) []error {
	var allErrs []error
	for err := range errs {
		allErrs = append(allErrs, err)
	}
	return allErrs
}

// TestErrorPropagation demonstrates error propagation with Goroutines and channels.
func TestErrorPropagation(t *testing.T) {
	const numWorkers = 3
	tasks := make(chan int, 10)
	errs := make(chan error, numWorkers)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start workers
	for i := 0; i < numWorkers; i++ {
		go func(i int) {
			Worker(i, tasks, errs)
			wg.Done()
		}(i)
	}

	// Generate tasks and send them to workers
	for i := 0; i < 20; i++ {
		tasks <- i
	}
	close(tasks)

	wg.Wait()

	// Gather errors from central error handler
	allErrs := CentralErrorHandler(errs)

	// Assert the expected number of errors
	expectedErrors := 0
	for _, task := range tasks {
		if task%3 == 0 {
			expectedErrors++
		}
	}

	if len(allErrs) != expectedErrors {
		t.Errorf("Expected %d errors, but got %d", expectedErrors, len(allErrs))
	}
}
