package main

import (
	"fmt"
	"sync"
	"time"
)

// TaskFunction represents a simple asynchronous task
func taskFunction(wg *sync.WaitGroup, id int, sleep time.Duration) {
	defer wg.Done()

	// Simulate a task that takes some time to complete
	time.Sleep(sleep)
	fmt.Printf("Task %d completed after %s\n", id, sleep)
}

func main() {
	var wg sync.WaitGroup
	numTasks := 5
	sleepDurations := []time.Duration{time.Second, 2 * time.Second, 1 * time.Second, 3 * time.Second, 2 * time.Second}

	// Start tasks asynchronously
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go taskFunction(&wg, i+1, sleepDurations[i])
	}

	// Wait for all tasks to complete before proceeding
	wg.Wait()

	// Once all tasks are done, gracefully shut down
	fmt.Println("All tasks completed, shutting down gracefully.")
}
