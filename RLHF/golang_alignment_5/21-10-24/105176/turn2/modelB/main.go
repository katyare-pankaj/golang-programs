package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers = 5
	numTasks   = 20
)

// task represents a unit of work to be performed concurrently
type task struct {
	id int
}

// worker processes tasks concurrently
func worker(id int, wg *sync.WaitGroup, tasks <-chan task, results chan<- int) {
	defer wg.Done()
	for task := range tasks {
		result := processTask(task)
		results <- result
	}
}

// processTask simulates work being done on a task
func processTask(task task) int {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	return task.id * 2
}

func main() {
	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// Create channels for tasks and results
	tasks := make(chan task, numTasks)
	results := make(chan int, numTasks)

	// Create a WaitGroup to wait for all workers to complete
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go worker(i, &wg, tasks, results)
	}

	// Generate and send tasks to the workers
	for i := 0; i < numTasks; i++ {
		tasks <- task{id: i + 1}
	}

	// Close the tasks channel to signal workers to stop
	close(tasks)

	// Wait for all workers to finish
	wg.Wait()

	// Receive and print results
	for i := 0; i < numTasks; i++ {
		result := <-results
		fmt.Printf("Task %d result: %d\n", i+1, result)
	}
}
