package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine exits.

	// Simulate work by sleeping for a random duration.
	sleepDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
	fmt.Printf("Worker %d: Sleeping for %v...\n", id, sleepDuration)
	time.Sleep(sleepDuration)

	fmt.Printf("Worker %d: Done!\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5 // Number of worker goroutines to create.

	// Add the number of worker goroutines to the WaitGroup.
	wg.Add(numWorkers)

	// Start worker goroutines.
	for i := 1; i <= numWorkers; i++ {
		go worker(&wg, i)
	}

	fmt.Println("Main goroutine: Waiting for workers to finish...")

	// Wait for all worker goroutines to call Done() on the WaitGroup.
	wg.Wait()

	fmt.Println("Main goroutine: All workers are done. Exiting...")
}
