package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	// Simulate some work being done
	// Replace this with your actual work
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	numWorkers := 5
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Launch Goroutines that call the worker function
	for i := 0; i < numWorkers; i++ {
		go worker(i+1, &wg)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	fmt.Println("All workers finished")
}
