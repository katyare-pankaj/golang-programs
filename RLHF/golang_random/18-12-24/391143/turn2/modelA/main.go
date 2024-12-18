package main

import (
	"fmt"
	"sync"
	"time"
)

// Simple function that prints a message after a delay
func printMessage(msg string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println(msg)
}

// Function to demonstrate how blocked Goroutines don't affect concurrency
func main() {
	// Initialize a WaitGroup to wait for all Goroutines to complete
	var wg sync.WaitGroup

	// Start a Goroutine that will be blocked for 2 seconds
	wg.Add(1)
	go printMessage("This message will be delayed by 2 seconds.", 2*time.Second, &wg)

	// Start more Goroutines that will print messages immediately
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go printMessage(fmt.Sprintf("Message %d printed immediately.", i), 0*time.Second, &wg)
	}

	// Wait for all Goroutines to complete
	wg.Wait()

	// After all Goroutines finish, print a final message
	fmt.Println("All Goroutines completed.")
}
