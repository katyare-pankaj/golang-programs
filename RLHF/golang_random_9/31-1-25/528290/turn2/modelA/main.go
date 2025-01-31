package main

import (
	"context"
	"fmt"
	"time"
)

// Callback is a functional interface type for processing tasks
type Callback func(ctx context.Context, id int) error

// worker is a goroutine function that processes a task with a given Callback
func worker(ctx context.Context, id int, callback Callback, done chan<- int) {
	select {
	case <-ctx.Done():
		fmt.Printf("Worker %d: canceled\n", id)
		done <- id
	default:
		err := callback(ctx, id)
		if err != nil {
			fmt.Printf("Worker %d: error %v\n", id, err)
		} else {
			fmt.Printf("Worker %d: completed successfully\n", id)
		}
	}
	done <- id
}

// differentTask1 is a task demonstrating a unique operation
func differentTask1(ctx context.Context, id int) error {
	fmt.Printf("Worker %d: performing task 1...\n", id)
	select {
	case <-time.After(1 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// differentTask2 is another task demonstrating a different operation
func differentTask2(ctx context.Context, id int) error {
	fmt.Printf("Worker %d: performing task 2...\n", id)
	select {
	case <-time.After(2 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// differentTask3 is yet another task demonstrating a unique operation
func differentTask3(ctx context.Context, id int) error {
	fmt.Printf("Worker %d: performing task 3...\n", id)
	select {
	case <-time.After(1 * time.Second):
		return fmt.Errorf("Worker %d had a simulated error", id)
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Channel to wait for all workers
	done := make(chan int, 3)

	tasks := []Callback{differentTask1, differentTask2, differentTask3}

	// Start each worker with different tasks
	for i, task := range tasks {
		go worker(ctx, i+1, task, done)
	}

	// Wait for all workers
	for range tasks {
		<-done
	}
}
