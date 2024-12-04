package main

import (
	"context"
	"fmt"
	"time"
)

// Worker simulates a task that takes time and can respond to cancellation.
func worker(ctx context.Context, result chan<- string) {
	select {
	case <-time.After(5 * time.Second):
		result <- "Task completed"
	case <-ctx.Done():
		result <- "Task cancelled"
	}
}

func main() {
	// Create a context that will be canceled after 3 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Ensure resources are cleaned up

	result := make(chan string, 1)

	// Start a worker goroutine.
	go worker(ctx, result)

	select {
	case res := <-result:
		fmt.Println(res)
	case <-ctx.Done():
		fmt.Println("Main: Context canceled or timed out")
	}
}
