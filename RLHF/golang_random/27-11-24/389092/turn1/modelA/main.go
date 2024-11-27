package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Create a task function that takes a WaitGroup as an argument
	task := func(wg *sync.WaitGroup, num int) {
		defer wg.Done() // Decrement the WaitGroup counter when the task is complete
		fmt.Printf("Starting task %d\n", num)
		time.Sleep(1 * time.Second) // Simulate work being done
		fmt.Printf("Finishing task %d\n", num)
	}

	// Increment the WaitGroup counter before starting each goroutine
	wg.Add(3)

	// Start 3 goroutines
	go task(&wg, 1)
	go task(&wg, 2)
	go task(&wg, 3)

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All tasks completed.")
}
