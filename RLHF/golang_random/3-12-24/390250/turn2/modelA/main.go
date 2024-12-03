package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Initialize the worker pool
	const numWorkers = 3
	const numTasks = 10
	var wg sync.WaitGroup
	workCh := make(chan int)
	doneCh := make(chan bool)

	// Start worker Goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, workCh, doneCh)
	}

	// Add tasks to the channel
	for i := 0; i < numTasks; i++ {
		workCh <- i
	}

	// Close the work channel when all tasks are added
	close(workCh)

	// Wait for all workers to finish
	for i := 0; i < numWorkers; i++ {
		<-doneCh
	}

	// Terminate the program
	wg.Wait()
	fmt.Println("All tasks are completed.")
}

func worker(id int, workCh <-chan int, doneCh chan<- bool) {
	defer wg.Done()
	for job := range workCh {
		fmt.Printf("Worker %d processing task %d\n", id, job)
		time.Sleep(time.Duration(job) * time.Second) // Simulate work with variable duration
	}
	doneCh <- true
}
