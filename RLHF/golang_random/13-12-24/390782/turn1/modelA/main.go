package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function simulates a task that may produce an error.
func worker(id int, ch chan<- error) {
	defer func() {
		if r := recover(); r != nil {
			ch <- fmt.Errorf("goroutine %d panicked: %v", id, r)
		}
	}()

	// Simulating some work
	time.Sleep(time.Duration(id) * time.Second)

	// Simulate an error on the worker with id 2
	if id == 2 {
		ch <- fmt.Errorf("worker %d encountered an error", id)
		return
	}

	// Successful execution
	ch <- nil
}

func main() {
	const numWorkers = 5
	errorChannel := make(chan error, numWorkers)
	var wg sync.WaitGroup

	// Start several worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, errorChannel)
		}(i)
	}

	wg.Wait()
	close(errorChannel)

	// Handle errors from the error channel
	var hasError bool
	for err := range errorChannel {
		if err != nil {
			fmt.Println("Error:", err)
			hasError = true
		}
	}

	if hasError {
		fmt.Println("One or more workers encountered an error.")
	} else {
		fmt.Println("All workers completed successfully.")
	}
}
