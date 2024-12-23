package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done() // Decrement the wait count when the worker is done
	fmt.Printf("Worker %d starting...\n", id)
	time.Sleep(time.Second) // Simulate some work
	fmt.Printf("Worker %d completed.\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Start multiple worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the wait count
		go worker(&wg, i)
	}

	// Wait for all workers to complete
	fmt.Println("Main waiting for workers...")
	wg.Wait()

	// Continue with main execution
	fmt.Println("All workers completed. Main continuing...")
}
