package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Callback is a functional interface type for processing tasks
type Callback func(ctx context.Context) error

// worker function to pick tasks from the channel and execute them
func worker(ctx context.Context, id int, tasks <-chan Callback, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: canceled\n", id)
			return
		default:
			if err := task(ctx); err != nil {
				fmt.Printf("Worker %d: error %v\n", id, err)
			} else {
				fmt.Printf("Worker %d: completed a task successfully\n", id)
			}
		}
	}
}

func main() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Define different tasks that can be performed by workers
	tasks := []Callback{
		func(ctx context.Context) error {
			fmt.Println("Task 1: Working...")
			select {
			case <-time.After(2 * time.Second):
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		},
		func(ctx context.Context) error {
			fmt.Println("Task 2: Working...")
			select {
			case <-time.After(1 * time.Second):
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		},
		func(ctx context.Context) error {
			fmt.Println("Task 3: Working...")
			select {
			case <-time.After(3 * time.Second):
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		},
		func(ctx context.Context) error {
			fmt.Println("Task 4: Working...")
			select {
			case <-time.After(2 * time.Second):
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		},
	}

	taskChan := make(chan Callback)
	var wg sync.WaitGroup

	// Start multiple workers
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, taskChan, &wg)
	}

	// Distribute tasks to workers via channel
	go func() {
		defer close(taskChan)
		for _, task := range tasks {
			select {
			case taskChan <- task:
			case <-ctx.Done():
				fmt.Println("Task distribution canceled")
				return
			}
		}
	}()

	// Wait for all workers to complete their tasks (or be canceled)
	wg.Wait()
	fmt.Println("All workers completed or canceled.")
}
