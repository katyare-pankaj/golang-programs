package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// worker simulates a task that takes some time to complete.
func worker(id int, wg *sync.WaitGroup, semaphore chan struct{}) {
	defer wg.Done()

	// Acquire a token from the semaphore to limit concurrency.
	semaphore <- struct{}{}
	defer func() {
		<-semaphore // Release the token.
	}()

	fmt.Printf("Worker %d starting\n", id)

	// Simulate work by sleeping for a random duration.
	sleepDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(sleepDuration)

	fmt.Printf("Worker %d done after %v\n", id, sleepDuration)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator.

	const numWorkers = 20    // Total number of workers.
	const maxConcurrency = 5 // Maximum number of concurrent workers.

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxConcurrency)

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, semaphore)
	}

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All workers completed.")
}
