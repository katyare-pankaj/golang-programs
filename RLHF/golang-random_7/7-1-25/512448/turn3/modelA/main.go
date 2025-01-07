package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define the task struct which holds input, result, and a completion flag
type task struct {
	input     int
	result    int
	completed bool
}

// Worker function processes tasks concurrently
func worker(wg *sync.WaitGroup, taskCh chan task, resultCh chan int) {
	defer wg.Done()
	for task := range taskCh {
		// Simulate work with random delay
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		// Process the task (in this case, square the input)
		task.result = task.input * task.input
		task.completed = true
		fmt.Printf("Worker: Completed task %d: %d * %d = %d\n", task.input, task.input, task.input, task.result)

		// Send the result back to the result channel
		resultCh <- task.result
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Number of tasks
	numTasks := 10
	// Number of worker goroutines
	numWorkers := 3

	// Channels for sending tasks and receiving results
	taskCh := make(chan task, numTasks)
	resultCh := make(chan int)

	// Wait group for workers
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(&wg, taskCh, resultCh)
	}

	// Enqueue tasks
	for i := 1; i <= numTasks; i++ {
		task := task{input: i}
		taskCh <- task
	}
	// Close the task channel to indicate no more tasks will be sent
	close(taskCh)

	// Wait for all worker goroutines to finish
	wg.Wait()

	// Aggregate results
	totalResult := 0
	for result := range resultCh {
		totalResult += result
		fmt.Printf("Collected result: %d\n", result)
	}
	// Close the result channel
	close(resultCh)

	// Summarize results
	fmt.Println("\nAll tasks completed. Summarizing results:")
	fmt.Printf("Total sum of results: %d\n", totalResult)
}
