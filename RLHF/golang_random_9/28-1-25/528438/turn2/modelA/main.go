package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Worker function that processes tasks from the tasks channel
func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup on completion of the worker

	for task := range tasks {
		// Simulating task processing by sleeping for a random time
		fmt.Printf("Worker %d started task %d\n", id, task)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Worker %d finished task %d\n", id, task)
	}
}

func main() {
	const numWorkers = 3
	const numTasks = 10

	// Create a channel for tasks
	tasks := make(chan int, numTasks)
	var wg sync.WaitGroup

	// Start a fixed number of workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Generate tasks and send them to the workers
	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}
	close(tasks) // Close the tasks channel to signal workers no more tasks are coming

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All tasks have been processed.")
}
