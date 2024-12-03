package main

import (
	"fmt"
	"time"
)

func main() {
	workQueue := make(chan int)
	done := make(chan struct{})

	// Start worker goroutines
	go func() {
		for work := range workQueue {
			fmt.Println("Processing work:", work)
			time.Sleep(1 * time.Second) // Simulate work
		}
		close(done)
	}()

	// Add work items
	for i := 0; i < 5; i++ {
		workQueue <- i
	}
	close(workQueue) // Signal workers to finish

	// Wait for workers to complete
	<-done
}
