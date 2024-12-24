package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the wait group counter when the goroutine finishes
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5

	fmt.Println("Starting workers...")
	// Add workers to the wait group
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers completed. Exiting the program.")
}
