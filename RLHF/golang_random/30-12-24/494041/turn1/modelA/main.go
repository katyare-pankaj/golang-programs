package main

import (
	"fmt"
	"sync"
	"time"
)

// NonBlockingTask simulates a non-blocking task done by a goroutine
func NonBlockingTask(wg *sync.WaitGroup) {
	defer wg.Done()                   // Signal the completion of this goroutine
	time.Sleep(time.Millisecond * 50) // Simulate work
	fmt.Println("Task completed.")
}

func main() {
	var wg sync.WaitGroup    // Initialize a WaitGroup
	const numGoroutines = 10 // Number of goroutines to create

	// Create and start goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)               // Increment the WaitGroup counter
		go NonBlockingTask(&wg) // Start a goroutine for the task
	}

	fmt.Println("Main thread starting to wait.")
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All tasks have completed.")
}
