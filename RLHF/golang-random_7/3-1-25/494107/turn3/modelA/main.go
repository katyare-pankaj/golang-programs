package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(w *sync.WaitGroup, index int) {
	// Simulate some work (in this case, just sleep for a short duration)
	time.Sleep(time.Second * time.Duration(index))

	// Increment the done call of WaitGroup when worker is done
	defer w.Done()

	fmt.Printf("Worker %d completed.\n", index)
}

func main() {
	// Initialize a new WaitGroup
	var wg sync.WaitGroup

	// Start 5 worker goroutines, each simulating a task
	for i := 0; i < 5; i++ {
		wg.Add(1)         // Increment WaitGroup counter
		go worker(&wg, i) // Pass the WaitGroup pointer to the worker
	}

	// Let the main goroutine know to wait for all workers to complete
	wg.Wait()

	// Once the WaitGroup completes, the main program proceeds
	fmt.Println("All worker goroutines have finished, main program continuing.")
}
