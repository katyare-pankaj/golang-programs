package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate varying work duration
	fmt.Printf("Worker %d completed\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 10

	// Start a separate WaitGroup for partial synchronization
	var partialWg sync.WaitGroup
	partialWorkers := 3
	partialWg.Add(partialWorkers)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Add to overall WaitGroup
		go func(i int) {
			defer wg.Done()
			if i <= partialWorkers {
				partialWg.Done() // Decrement the partial WaitGroup if worker is in the first phase
			}
			worker(i, &wg)
		}(i)
	}

	// Wait for the first phase of workers to complete
	partialWg.Wait()
	fmt.Println("First phase completed.")

	// Do some other work here...

	// Wait for all workers to complete their tasks
	wg.Wait()
	fmt.Println("All workers completed. Main goroutine exiting.")
}
