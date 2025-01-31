package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Callback is a functional interface type for processing tasks
type Callback func(ctx context.Context) error

// worker is a goroutine function that processes tasks from a shared channel
func worker(ctx context.Context, id int, taskChan chan Callback, processedTasks map[int]bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: canceled\n", id)
			return
		case task, ok := <-taskChan:
			if !ok {
				// Channel closed, no more tasks
				return
			}
			// Check if the task has already been processed
			taskID := rand.Int() // Generate a unique task ID
			if _, ok := processedTasks[taskID]; ok {
				continue
			}
			processedTasks[taskID] = true

			err := task(ctx)
			if err != nil {
				fmt.Printf("Worker %d: error %v\n", id, err)
			} else {
				fmt.Printf("Worker %d: completed a task successfully\n", id)
			}
		}
	}
}

func main() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define different tasks
	task1 := func(ctx context.Context) error {
		fmt.Println("Task 1: Working...")
		select {
		case <-time.After(2 * time.Second):
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	task2 := func(ctx context.Context) error {
		fmt.Println("Task 2: Working...")
		select {
		case <-time.After(1 * time.Second):
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	task3 := func(ctx context.Context) error {
		fmt.Println("Task 3: Working...")
		select {
		case <-time.After(3 * time.Second):
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	// Create a slice of tasks to be shared by workers
	tasks := []Callback{task1, task2, task3}

	// Number of workers
	numWorkers := 3

	// Create a channel to share tasks
	taskChan := make(chan Callback)

	// Map to keep track of processed tasks
	processedTasks := make(map[int]bool)

	// WaitGroup to synchronize main goroutine with workers
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, i, taskChan, processedTasks, &wg)
	}

	// Add tasks to the channel
	for _, task := range tasks {
		taskChan <- task
	}

	close(taskChan) // Signal that no more tasks will be added

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All workers completed.")
}
