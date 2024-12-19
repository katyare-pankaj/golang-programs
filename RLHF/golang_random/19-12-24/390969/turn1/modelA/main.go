package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	id   int
	data string
}

func main() {
	const numWorkers = 4
	const numTasks = 100
	var wg sync.WaitGroup

	// Create a worker pool
	workerPool := make(chan struct{}, numWorkers)

	// Create a channel for tasks
	taskChannel := make(chan task, numTasks)

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(workerPool, taskChannel, &wg)
	}

	// Generate and send tasks
	for i := 0; i < numTasks; i++ {
		taskChannel <- task{id: i + 1, data: fmt.Sprintf("Task %d", i+1)}
	}

	// Close the task channel when all tasks are sent
	close(taskChannel)

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All tasks processed.")
}

func worker(workerPool chan struct{}, taskChannel <-chan task, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// Wait for a worker slot to become available
		<-workerPool
		defer func() { workerPool <- struct{}{} }()

		select {
		case task, ok := <-taskChannel:
			if !ok {
				return // Channel is closed, exit worker
			}
			fmt.Printf("Worker processing task: %d with data: %s\n", task.id, task.data)
			time.Sleep(time.Duration(int64(task.id)) * time.Millisecond)
		default:
			time.Sleep(1 * time.Millisecond)
		}
	}
}
