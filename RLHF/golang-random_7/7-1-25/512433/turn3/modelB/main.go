package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, errCh chan error) {
	defer wg.Done()
	fmt.Printf("Worker %d is starting\n", id)
	// Simulate work being done
	time.Sleep(time.Duration(id) * time.Second)

	// Introduce an error for demonstration purposes
	if id == 3 {
		errCh <- errors.New("worker 3 encountered an error")
		return
	}

	// Simulate successful completion
	errCh <- nil
}
func main() {
	var wg sync.WaitGroup
	errCh := make(chan error, 3) // Buffered channel to hold errors
	// Add 3 workers to the waitgroup
	wg.Add(3)
	// Start 3 workers
	for i := 0; i < 3; i++ {
		go worker(i, &wg, errCh)
	}
	// Wait for all workers to finish
	wg.Wait()
	close(errCh) // Close the error channel after all workers have completed

	// Process errors
	for err := range errCh {
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	fmt.Println("All workers are done")
}
