package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Define the number of goroutines
	numGoroutines := 5

	// Create a WaitGroup to track the completion of goroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Create a slice to store goroutine IDs
	goroutineIDs := make([]int, numGoroutines)

	// Launch goroutines that will perform work
	for i := 0; i < numGoroutines; i++ {
		goroutineIDs[i] = i + 1
		go func(id int) {
			defer wg.Done()
			// Simulate work being done
			time.Sleep(time.Duration(id) * time.Second)
			fmt.Printf("Goroutine %d completed its work.\n", id)
		}(goroutineIDs[i])
	}

	// Start the progress bar
	fmt.Println("Starting Goroutines...")
	for i := 0; i < numGoroutines; i++ {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
	fmt.Println()

	// Wait for all goroutines to complete
	wg.Wait()

	// All goroutines are done, print a final message
	fmt.Println("All goroutines have completed their work.")
}
