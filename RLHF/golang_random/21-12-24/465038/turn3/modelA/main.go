package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	numGoroutines := 5

	// Increment the WaitGroup counter for each goroutine.
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go worker(i, &wg)
	}

	// Wait for all goroutines to finish.
	wg.Wait()

	fmt.Println("All goroutines completed.")
}

// worker is a function that will be executed concurrently.
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the wait group counter when the goroutine ends.
	fmt.Printf("Goroutine %d started.\n", id)

	// Simulate work by sleeping for a random duration.
	time.Sleep(time.Duration(id*2) * time.Second)

	fmt.Printf("Goroutine %d finished.\n", id)
}
