package main

import (
	"context"
	"fmt"
	"time"
)

// Callback is a functional interface type for processing tasks
type Callback func(ctx context.Context) error

// worker is a goroutine function that processes a task with a given Callback
func worker(ctx context.Context, id int, tasks []Callback) {
	for _, task := range tasks {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: canceled\n", id)
			return
		default:
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Define different tasks that can be performed by workers
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

	// Create a slice of tasks to be performed by workers
	tasks := []Callback{task1, task2, task3}

	// Start multiple workers, each with a slice of tasks to perform
	for i := 1; i <= 3; i++ {
		go worker(ctx, i, tasks)
	}

	// Wait for all workers to complete their tasks (or be canceled)
	<-ctx.Done()
	fmt.Println("All workers completed or canceled.")
}
