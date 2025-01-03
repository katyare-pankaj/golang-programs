package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(w *sync.WaitGroup, id int, message string) {
	defer w.Done()                              // Signal that this goroutine is complete
	time.Sleep(time.Duration(id) * time.Second) // Simulate work
	fmt.Printf("Worker %d: %s\n", id, message)
}

func main() {
	// Create a new WaitGroup
	var wg sync.WaitGroup

	// Start multiple goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the counter
		go worker(&wg, i, "Working...")
	}

	// The main goroutine waits for all worker goroutines to complete
	wg.Wait()

	// Continue with the main program after all workers are done
	fmt.Println("All workers have completed their tasks.")
}
