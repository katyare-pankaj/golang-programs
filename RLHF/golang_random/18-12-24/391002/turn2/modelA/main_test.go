package main

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

func work(id int, ch chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	if id%2 == 0 {
		ch <- fmt.Errorf("error with even id: %d", id)
	} else {
		time.Sleep(100 * time.Millisecond) // Simulate some work
		ch <- nil
	}
}

func runWorkers(numWorkers int) (errors []error) {
	var wg sync.WaitGroup
	errorCh := make(chan error, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go work(i, errorCh, &wg)
	}

	wg.Wait()

	close(errorCh)

	for err := range errorCh {
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func TestErrorPropagation(t *testing.T) {
	numWorkers := 10
	errors := runWorkers(numWorkers)

	// Validate that exactly half the workers encountered errors
	if len(errors) != numWorkers/2 {
		t.Errorf("expected %d errors, but got %d", numWorkers/2, len(errors))
	}

	// Ensure all errors are related to even IDs
	for _, err := range errors {
		if !strings.Contains(err.Error(), "error with even id") {
			t.Errorf("expected error for even ID, got %v", err)
		}
	}
}

func main() {
	// For demonstration purposes, run the test
	t := &testing.T{}
	TestErrorPropagation(t)
	fmt.Println("All tests passed.")
}
