package main

import (
	"fmt"
	"sync"
	"testing"
)

// worker performs a task that might fail and returns an error via the error channel.
func worker(wg *sync.WaitGroup, num int, errCh chan error) {
	defer wg.Done()
	if num == 3 {
		// Simulate an error for number 3
		errCh <- fmt.Errorf("error processing number %d", num)
		return
	}
	fmt.Printf("Processed %d successfully\n", num)
}

// TestWorkers demonstrates how to handle errors in Goroutines using a sync.WaitGroup and a channel.
func TestWorkers(t *testing.T) {
	var wg sync.WaitGroup
	numbers := []int{1, 2, 3, 4, 5}
	errCh := make(chan error)

	for _, num := range numbers {
		wg.Add(1)
		go worker(&wg, num, errCh)
	}

	// Wait for all workers to complete
	wg.Wait()

	// Close the error channel
	close(errCh)

	// Check if there were any errors
	var encounteredError error
	for err := range errCh {
		if err != nil {
			encounteredError = err
			break
		}
	}

	if encounteredError != nil {
		t.Fatalf("An error occurred: %v", encounteredError)
	}
}

func TestMain(m *testing.M) {
	// Call the testing framework's "main"
	m.Run()
}
