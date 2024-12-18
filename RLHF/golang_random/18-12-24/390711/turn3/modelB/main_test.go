package main

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func worker(wg *sync.WaitGroup, num int, errCh chan<- error) {
	defer wg.Done()
	// Simulate some work
	time.Sleep(time.Duration(num) * 50 * time.Millisecond)

	// Introduce a random error for demonstration purposes
	if num == 3 {
		errCh <- errors.New("worker failed for number 3")
		return
	}

	fmt.Printf("Processed %d\n", num)
}

func TestWorker(t *testing.T) {
	var wg sync.WaitGroup
	errCh := make(chan error, 1) // Buffered channel to hold one error
	numbers := []int{1, 2, 3, 4, 5}

	for _, num := range numbers {
		wg.Add(1)
		go worker(&wg, num, errCh)
	}

	go func() {
		wg.Wait()
		close(errCh) // Close the error channel when all workers are done
	}()

	// Wait for either all workers to complete or an error to occur
	var err error
	select {
	case err = <-errCh:
		// An error occurred
		t.Errorf("Error occurred: %v", err)
	default:
		// All workers completed successfully
		fmt.Println("All workers completed successfully")
	}
}
