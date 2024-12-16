package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()                             // Decrement the WaitGroup counter when the goroutine is done
	time.Sleep(time.Duration(id) * time.Second) // Simulate work
	fmt.Println("Worker", id, "finished.")
}

func main() {
	var wg sync.WaitGroup // Create a WaitGroup

	// Start several worker goroutines
	wg.Add(3) // Add 3 to the WaitGroup counter
	go worker(1, &wg)
	go worker(2, &wg)
	go worker(3, &wg)

	fmt.Println("Main function starting other tasks...")
	// Simulate additional work in the main goroutine
	time.Sleep(1 * time.Second)

	fmt.Println("Main function waiting for workers to finish...")
	wg.Wait() // Wait for all goroutines to complete

	fmt.Println("All workers have finished. Main function exiting.")
}
