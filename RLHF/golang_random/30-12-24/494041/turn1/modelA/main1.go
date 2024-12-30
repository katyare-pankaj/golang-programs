package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup    // Initialize a WaitGroup
	const numGoroutines = 10 // Number of goroutines to create

	// Create and start goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go func(id int) {
			defer wg.Done()                   // Signal the completion of this goroutine
			time.Sleep(time.Millisecond * 50) // Simulate work
			fmt.Printf("Task %d completed.\n", id)
		}(i)
	}

	fmt.Println("Main thread starting to wait.")
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All tasks have completed.")
}
