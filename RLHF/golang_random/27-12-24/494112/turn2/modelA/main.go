package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// A simple task that simulates work
func computeWork(id int) {
	rand.Seed(int64(time.Now().UnixNano()))
	for i := 0; i < 1_000_000; i++ {
		// Perform a computationally intensive task
		result := 0
		for j := 0; j < 1_000_000; j++ {
			result += rand.Int()
		}
	}
	fmt.Printf("Worker %d completed work\n", id)
}

func main() {
	// Set the number of CPU cores to use
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			computeWork(id)
		}(i)
	}

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers completed work")
}
