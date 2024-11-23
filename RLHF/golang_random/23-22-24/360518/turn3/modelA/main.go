package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	maxTasksPerSecond = 5 // Set the rate limit to 5 tasks per second
	tokenBucketSize   = maxTasksPerSecond
	resetInterval     = time.Second
)

type Task struct {
	ID   int
	Data string
}

var (
	taskProcessorCh = make(chan Task) // Channel to process tasks
	wg              sync.WaitGroup
	tokenBucket     = make(chan struct{}, tokenBucketSize)
	taskCounter     int
)

func main() {
	// Start a goroutine to refill the token bucket every second
	go refillTokenBucket()

	// Start goroutines to process tasks from different sources
	go processTasks()

	// Process tasks from multiple sources
	go handleTasksFromSource1()
	go handleTasksFromSource2()

	// Wait for all tasks to be processed
	wg.Wait()
}

func refillTokenBucket() {
	ticker := time.NewTicker(resetInterval)
	defer ticker.Stop()

	for range ticker.C {
		// Reset the token bucket to the maximum size
		for i := 0; i < tokenBucketSize; i++ {
			select {
			case tokenBucket <- struct{}{}:
			default:
				// Bucket is full
				return
			}
		}
	}
}

func processTasks() {
	for task := range taskProcessorCh {
		// Take a token to process the task
		<-tokenBucket

		// Simulate processing the task
		fmt.Printf("Processing Task %d: %s\n", task.ID, task.Data)
		time.Sleep(time.Millisecond * 500) // Simulate task processing time

		// Return the token
		tokenBucket <- struct{}{}
	}
}

func handleTasksFromSource1() {
	for i := 1; i <= 10; i++ {
		// Send a task from source 1
		task := Task{ID: i, Data: fmt.Sprintf("Task from Source1 - %d", i)}
		taskProcessorCh <- task
	}
	wg.Done()
}

func handleTasksFromSource2() {
	for i := 11; i <= 20; i++ {
		// Send a task from source 2
		task := Task{ID: i, Data: fmt.Sprintf("Task from Source2 - %d", i)}
		taskProcessorCh <- task
	}
	wg.Done()
}
