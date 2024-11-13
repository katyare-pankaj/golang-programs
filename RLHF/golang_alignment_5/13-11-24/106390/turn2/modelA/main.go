package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// worker represents a goroutine that processes tasks from a channel
func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d started task: %d\n", id, task)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Worker %d completed task: %d\n", id, task)
	}
}

func main() {
	const numWorkers = 3
	const numTasks = 10

	// Create a wait group to wait for all workers to finish
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Create a channel to send tasks to workers
	tasks := make(chan int)

	// Start the worker goroutines
	for i := 0; i < numWorkers; i++ {
		go worker(i, tasks, &wg)
	}

	// Generate and send tasks to the channel
	for i := 0; i < numTasks; i++ {
		tasks <- i
	}

	// Close the channel to signal to workers that no more tasks will be sent
	close(tasks)

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All tasks completed.")
}
