package main

import (
	"context"
	"fmt"
	"time"
)

// ProcessWork simulates processing work and returns an error if an issue occurs.
func ProcessWork(ctx context.Context, id int, workCh <-chan int, resultCh chan<- struct{ error }) {
	defer close(resultCh)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: context cancelled\n", id)
			return
		case task, ok := <-workCh:
			if !ok {
				fmt.Printf("Worker %d: work channel closed\n", id)
				return
			}

			// Simulate work
			time.Sleep(time.Duration(task) * time.Millisecond)

			// Simulate an occasional error
			if task%3 == 0 {
				resultCh <- struct{ error }{fmt.Errorf("error processing task %d", task)}
			} else {
				resultCh <- struct{ error }{nil}
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	workCh := make(chan int)
	results := make(chan struct{ error })

	// Start workers
	for i := 1; i <= 3; i++ {
		go ProcessWork(ctx, i, workCh, results)
	}

	// Send work tasks
	for i := 1; i <= 10; i++ {
		workCh <- i
	}

	// Close work channel
	close(workCh)

	// Collect and handle results
	for i := 1; i <= 3; i++ {
		select {
		case result := <-results:
			if result.error != nil {
				fmt.Printf("Error: %v\n", result.error)
			} else {
				fmt.Printf("Task completed successfully\n")
			}
		case <-ctx.Done():
			fmt.Println("Main context cancelled")
			return
		}
	}
}
