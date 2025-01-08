package main

import (
	"fmt"
	"sync"
	"time"
)

func simulateWork(id int, wg *sync.WaitGroup, workDuration time.Duration) {
	defer wg.Done() // Signal completion once the work is done
	fmt.Printf("Worker %d is starting and will work for %s\n", id, workDuration)
	time.Sleep(workDuration) // Simulate work by sleeping
	fmt.Printf("Worker %d is done\n", id)
}

func main() {
	var wg sync.WaitGroup // Initialize the wait group

	workDurations := []time.Duration{
		time.Second * 1,
		time.Second * 2,
		time.Second * 3,
		time.Second * 4,
		time.Second * 5,
	}

	// Add workers to the wait group
	for i, duration := range workDurations {
		wg.Add(1)                         // Increment the wait count for each worker
		go simulateWork(i, &wg, duration) // Start the worker goroutine
	}

	// Wait for all workers to finish
	fmt.Println("Main goroutine is waiting for workers...")
	wg.Wait() // Block until the wait count reaches zero

	fmt.Println("All workers have completed their tasks.")
}
