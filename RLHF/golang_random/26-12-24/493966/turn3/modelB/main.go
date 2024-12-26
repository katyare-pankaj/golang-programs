package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker(id int) {
	for {
		fmt.Printf("Worker %d is running\n", id)
		// Simulate work
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}

func main() {
	runtime.GOMAXPROCS(1) // Limit Go to use only 1 thread for simplicity

	numWorkers := 3
	for i := 0; i < numWorkers; i++ {
		go worker(i) // Create and start worker goroutines
	}

	// The main goroutine runs continuously while other goroutines are scheduled
	for {
		// Let's introduce a slight delay in the main goroutine
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
}
