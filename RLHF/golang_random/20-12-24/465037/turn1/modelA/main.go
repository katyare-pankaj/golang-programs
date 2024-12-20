package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	wg = &sync.WaitGroup{}
)

func worker(id int) {
	defer wg.Done()
	// Simulate work
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d done.\n", id)
}

func addWorkers(count int) {
	// Ensure the main goroutine waits for all workers to complete
	wg.Add(count)
	for i := 0; i < count; i++ {
		go worker(i)
	}
	wg.Wait()
}

func main() {
	// Initialize wait group
	wg = &sync.WaitGroup{}

	// Start workers
	addWorkers(5)

	fmt.Println("All workers completed.")

	// Clean up resources (if any)
	defer func() {
		log.Println("Cleaning up resources...")
	}()
}
