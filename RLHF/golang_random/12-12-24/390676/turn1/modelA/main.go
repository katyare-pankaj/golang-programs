package main

import (
	"fmt"
	"time"
)

// Data structure to simulate work being done
type Job struct {
	ID      int
	Payload string
}

// Error to simulate failure
type JobError struct {
	ID       int
	ErrorMsg string
}

// Worker function to process jobs
func worker(id int, jobs <-chan Job, errors chan<- JobError, done chan<- struct{}) {
	for job := range jobs {
		// Simulate processing time
		time.Sleep(time.Second)

		// Simulate error condition
		if job.ID%2 == 0 { // Simulate error for even job IDs
			errors <- JobError{ID: job.ID, ErrorMsg: "failed to process job"}
		} else {
			fmt.Printf("Worker %d processed job: %d\n", id, job.ID)
		}
	}
	done <- struct{}{} // Signal completion
}

func main() {
	jobs := make(chan Job)
	errors := make(chan JobError)
	done := make(chan struct{})

	// Start worker goroutines
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, errors, done)
	}

	// Send jobs
	go func() {
		for i := 0; i < 10; i++ {
			jobs <- Job{ID: i, Payload: fmt.Sprintf("data-%d", i)}
		}
		close(jobs) // Close jobs channel when done sending
	}()

	// Wait for workers to finish processing
	go func() {
		for i := 0; i < 3; i++ {
			<-done
		}
		close(errors) // Close errors channel when all workers are done
	}()

	// Handle errors
	for err := range errors {
		fmt.Printf("Error processing job %d: %s\n", err.ID, err.ErrorMsg)
	}

	fmt.Println("All workers have finished processing.")
}
