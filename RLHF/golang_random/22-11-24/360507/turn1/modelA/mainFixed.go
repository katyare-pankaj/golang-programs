package main

import (
	"fmt"
	"time"
)

func worker(work <-chan int, done chan bool) {
	for {
		select {
		case num := <-work:
			// Simulate work
			fmt.Printf("Worker received %d\n", num)
			time.Sleep(time.Second) // Simulate processing time
		case <-done:
			// Shut down worker
			fmt.Println("Worker shutting down")
			return
		}
	}
}

func main() {
	// Create a channel to send work to the worker
	work := make(chan int)
	// Create a channel to signal the worker to shutdown
	done := make(chan bool)

	// Start a worker goroutine
	go worker(work, done)

	// Send some work to the worker
	for i := 0; i < 5; i++ {
		work <- i
	}
	time.Sleep(5 * time.Second) // Wait for all work to be processed

	// Signal the worker to stop
	done <- true

	// Wait for the worker goroutine to exit
	select {} // Keeps the main goroutine alive until the worker exits
}
