package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a new WaitGroup
	var wg sync.WaitGroup

	// Add a count of 3 to the WaitGroup, one for each goroutine
	wg.Add(3)

	// Start the first goroutine
	go func() {
		defer wg.Done() // Call Done() when the goroutine is finished
		fmt.Println("Goroutine 1 is running...")
		time.Sleep(1 * time.Second) // Simulate work by sleeping for 1 second
		fmt.Println("Goroutine 1 has finished.")
	}()

	// Start the second goroutine
	go func() {
		defer wg.Done() // Call Done() when the goroutine is finished
		fmt.Println("Goroutine 2 is running...")
		time.Sleep(2 * time.Second) // Simulate work by sleeping for 2 seconds
		fmt.Println("Goroutine 2 has finished.")
	}()

	// Start the third goroutine
	go func() {
		defer wg.Done() // Call Done() when the goroutine is finished
		fmt.Println("Goroutine 3 is running...")
		time.Sleep(3 * time.Second) // Simulate work by sleeping for 3 seconds
		fmt.Println("Goroutine 3 has finished.")
	}()

	// Wait for all goroutines to complete
	wg.Wait()

	// Print a message to indicate the main function is finishing
	fmt.Println("All goroutines have finished. Main function is exiting.")
}
