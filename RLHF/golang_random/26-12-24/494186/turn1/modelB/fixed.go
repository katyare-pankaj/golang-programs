package main

import (
	"errors"
	"fmt"
	"time"
)

func worker(id int, done chan bool, errChan chan error) {
	time.Sleep(time.Second * 2)
	if id == 2 {
		errChan <- errors.New("worker 2 failed")
		return
	}
	fmt.Println("Worker", id, "completed")
	done <- true // Indicate worker completion
}

func main() {
	numWorkers := 3
	done := make(chan bool, numWorkers)     // Buffered channel to avoid blocking
	errChan := make(chan error, numWorkers) // Buffered channel to avoid blocking

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, done, errChan)
	}

	// Wait for workers to finish and handle any errors
	for i := 1; i <= numWorkers; i++ {
		select {
		case <-done:
			fmt.Println("Worker", i, "done")
		case err := <-errChan:
			fmt.Println("Error:", err)
		}
	}

	// Close channels after all workers are done
	close(done)
	close(errChan)
}
