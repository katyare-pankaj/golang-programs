package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

func main() {
	// Create a WaitGroup instance
	var wg sync.WaitGroup

	// Number of goroutines to start
	numGoroutines := 5

	// Increment the WaitGroup counter for each goroutine
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			fmt.Printf("Goroutine %d is starting...\n", id)
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second) // Simulate some work
			fmt.Printf("Goroutine %d has finished.\n", id)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print a message once all goroutines are done
	fmt.Println("All goroutines completed.")
}
