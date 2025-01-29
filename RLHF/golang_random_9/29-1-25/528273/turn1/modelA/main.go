package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate some work being done by a worker.
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Deferred function call

	fmt.Printf("Worker %d: Starting\n", id)

	// Simulate some work by sleeping.
	time.Sleep(2 * time.Second)

	fmt.Printf("Worker %d: Finished\n", id)
}

// Optimized worker without using defer
func optimizedWorker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Optimized Worker %d: Starting\n", id)
	defer wg.Done() // This can be inside function scope without affecting performance significantly.

	// Simulate some work by sleeping.
	time.Sleep(2 * time.Second)

	fmt.Printf("Optimized Worker %d: Finished\n", id)
}

func main() {
	const numWorkers = 5
	var wg sync.WaitGroup

	fmt.Println("Starting workers with defer:")

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, &wg)
	}
	wg.Wait()

	fmt.Println("All workers with defer finished\n")

	fmt.Println("Starting optimized workers without defer inside loop:")
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go optimizedWorker(i, &wg)
	}
	wg.Wait()

	fmt.Println("All optimized workers finished")
}
