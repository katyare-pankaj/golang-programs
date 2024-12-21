package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func longRunningTask(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("Long-running task started.")

	for {
		select {
		case <-ctx.Done():
			log.Println("Long-running task canceled by context.")
			return
		default:
			// Simulate work
			log.Println("Doing work...")
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the long-running task
	wg.Add(1)
	go longRunningTask(ctx, &wg)

	// Simulate a condition where the task needs to be canceled
	time.Sleep(time.Second * 2)
	fmt.Println("Canceling the task...")
	cancel()

	// Wait for the task to finish
	wg.Wait()
	fmt.Println("Task finished gracefully.")
}
