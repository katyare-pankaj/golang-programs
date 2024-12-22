package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

// Function to simulate a worker task
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d: Starting task...\n", id)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate varying task execution times
	fmt.Printf("Worker %d: Task completed.\n", id)
}

func main() {
	// Number of worker goroutines
	const numWorkers = 5
	// Initialize WaitGroup
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1) // Increment the WaitGroup count for each worker
		go worker(i+1, &wg)
	}

	// Wait for all workers to complete
	fmt.Println("Waiting for all workers to complete...")
	wg.Wait() // The main goroutine will block here until all workers have called wg.Done()

	fmt.Println("All workers have completed their tasks.")
}
