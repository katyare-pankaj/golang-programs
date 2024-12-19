package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a task to be processed
type Task struct {
	ID   int
	Data string
}

// Worker processes tasks from a channel
func Worker(wg *sync.WaitGroup, work chan Task) {
	defer wg.Done()
	for task := range work {
		fmt.Printf("Processing task %d: %s\n", task.ID, task.Data)
		// Simulate task processing time
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5

	// Start worker goroutines
	work := make(chan Task, 100) // Buffer the work channel
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Worker(&wg, work)
	}

	// Generate and send tasks to the work channel
	for i := 1; i <= 100; i++ {
		work <- Task{ID: i, Data: fmt.Sprintf("Task %d", i)}
	}

	// Close the work channel to signal workers to exit
	close(work)
	wg.Wait()

	fmt.Println("All tasks completed.")
}
