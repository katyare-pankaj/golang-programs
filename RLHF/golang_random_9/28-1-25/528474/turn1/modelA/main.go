package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulate a task that takes a variable amount of time to complete.
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Indicate that this goroutine is done once completed.

	// Simulate a workload with random duration.
	duration := time.Duration(rand.Intn(1000)) * time.Millisecond
	fmt.Printf("Worker %d starting, will take %v\n", id, duration)
	time.Sleep(duration)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	// Number of workers.
	numWorkers := 10

	// Add the number of workers to the WaitGroup.
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(i, &wg)
	}

	// Wait for all workers to complete.
	wg.Wait()

	fmt.Println("All workers completed.")
}
