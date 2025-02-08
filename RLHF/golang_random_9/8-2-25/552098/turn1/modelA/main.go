package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker simulates processing of jobs from a jobs channel and signals completion on a results channel
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure the wait group is decremented when work is done
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Send the result
	}
}

func main() {
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start the workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Prepare a deferred function to clean up the results channel
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(results)
		close(done) // Signal that all work is done
	}()

	// Send jobs to the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // No more jobs will be sent

	// Collect the results
	go func() {
		<-done // Wait for all workers to be done
		for result := range results {
			fmt.Printf("Result: %d\n", result)
		}
	}()

	// Adding a small sleep to prevent the main from ending too soon
	time.Sleep(2 * time.Second)
	fmt.Println("All jobs processed")
}
