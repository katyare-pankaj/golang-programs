package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Add a wait group for each goroutine we're starting.
	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second) // Simulate work
		fmt.Println("Worker 1 completed.")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second) // Simulate work
		fmt.Println("Worker 2 completed.")
	}()

	// Wait for all goroutines to finish before exiting the main function.
	fmt.Println("Waiting for workers to complete...")
	wg.Wait()
	fmt.Println("All workers completed. Exiting main function.")
}
