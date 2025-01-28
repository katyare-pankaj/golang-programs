package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SimulatedTask represents a task with an ID and duration.
type SimulatedTask struct {
	id       int
	duration time.Duration
}

// worker is a function that processes tasks from the task channel.
func worker(id int, tasks <-chan SimulatedTask, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d starting task %d with duration %v\n", id, task.id, task.duration)
		time.Sleep(task.duration) // Simulate the workload
		fmt.Printf("Worker %d completed task %d\n", id, task.id)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const (
		numWorkers    = 5
		numTasks      = 20
		maxTaskTimeMs = 2000 // Maximum task duration in milliseconds
	)

	var wg sync.WaitGroup
	taskChannel := make(chan SimulatedTask, numTasks)

	// Start the worker pool
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, taskChannel, &wg)
	}

	// Generate and send tasks with varying durations
	for i := 0; i < numTasks; i++ {
		duration := time.Duration(rand.Intn(maxTaskTimeMs)) * time.Millisecond
		taskChannel <- SimulatedTask{id: i, duration: duration}
	}

	// Close task channel once all tasks have been sent
	close(taskChannel)

	// Wait for all workers to finish processing
	wg.Wait()

	fmt.Println("All tasks have been processed.")
}
