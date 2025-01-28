package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	// Maximum duration for a task
	MaxDuration = 5 * time.Second
	// Number of workers
	NumWorkers = 10
	// Buffer size for the work channel
	WorkerChannelBuffer = 10
)

func worker(id int, work <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for taskID := range work {
		// Simulate a variable workload using a random duration
		duration := time.Duration(rand.Int63n(int64(MaxDuration)))
		fmt.Printf("Worker %d starting task %d with duration: %v\n", id, taskID, duration)
		time.Sleep(duration)
		fmt.Printf("Worker %d completed task %d\n", id, taskID)
	}
}
func main() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Create a work channel with a buffer
	work := make(chan int, WorkerChannelBuffer)
	var wg sync.WaitGroup

	// Launch worker goroutines
	for i := 0; i < NumWorkers; i++ {
		wg.Add(1)
		go worker(i, work, &wg)
	}

	// Generate and send tasks to workers
	const totalTasks = 100
	for taskID := 0; taskID < totalTasks; taskID++ {
		work <- taskID
	}

	// Close the work channel to signal workers to exit
	close(work)

	fmt.Println("Sent all tasks, waiting for workers to finish...")

	// Wait for all workers to complete
	wg.Wait()

	fmt.Println("All workers completed.")
}
