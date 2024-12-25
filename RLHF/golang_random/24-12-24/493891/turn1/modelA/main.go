package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	defer wg.Wait() // Wait for all goroutines to finish before the main function exits

	numWorkers := 5
	ch := make(chan int)

	// Worker function that processes jobs and sends a completion signal
	for i := 0; i < numWorkers; i++ {
		wg.Add(1) // Increment WaitGroup count for each worker
		go func() {
			defer wg.Done() // Decrement WaitGroup count when the worker is done
			for job := range ch {
				time.Sleep(time.Duration(job) * time.Millisecond)
				fmt.Printf("Completed job: %d\n", job)
			}
		}()
	}

	// Generate jobs and send them to the worker channel
	numJobs := 10
	for i := 1; i <= numJobs; i++ {
		ch <- i * 100 // Send job time in milliseconds
	}

	close(ch) // Signal the workers that there are no more jobs
}
