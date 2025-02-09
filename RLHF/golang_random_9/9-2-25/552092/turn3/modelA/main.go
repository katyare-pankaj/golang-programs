package main

import (
	"fmt"
	"sync"
)

// task simulates a running task in a goroutine
func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure Done is called to decrement the counter

	// Simulate some work
	fmt.Printf("Goroutine %d is starting\n", id)
	// In actual work, there might be operations such as HTTP requests, file processing, etc.
	// Here we simulate it with a simple print statement, but any operation can be placed here.
	fmt.Printf("Goroutine %d is working\n", id)

	// Simulate more work if necessary

	fmt.Printf("Goroutine %d has finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 5

	// Add the number of goroutines to the WaitGroup counter
	wg.Add(numGoroutines)

	for i := 1; i <= numGoroutines; i++ {
		// Launch goroutines
		go task(i, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("All goroutines have completed")
}
