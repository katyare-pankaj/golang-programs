package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Creating a buffered channel to communicate between Goroutines
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	// Worker Goroutine that processes jobs
	worker := func(id int, jobs <-chan int, results chan<- int, done chan<- bool) {
		for job := range jobs {
			fmt.Printf("Worker %d processing job %d\n", id, job)
			time.Sleep(time.Second) // Simulating work
			results <- job * 2      // Sending result back
		}
		done <- true
	}

	// Launching multiple workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			done := make(chan bool)
			go worker(id, jobs, results, done)
			<-done // Wait for worker to finish
		}(w)
	}

	// Sending jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // No more jobs to send

	// Closing results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collecting results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}

	fmt.Println("All jobs processed.")
}
