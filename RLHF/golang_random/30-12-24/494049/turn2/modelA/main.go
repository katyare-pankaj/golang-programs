package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func simulateWork(wg *sync.WaitGroup, id int) {
	defer wg.Done() // Decrement the WaitGroup count once the function exits

	// Simulate some work by sleeping for a random duration
	duration := time.Duration(rand.Intn(5)) * time.Second // Sleep between 0s and 4s
	fmt.Printf("Worker %d is working for %v\n", id, duration)
	time.Sleep(duration)

	fmt.Printf("Worker %d is done\n", id)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Set the number of goroutines
	numGoroutines := 5

	// Add the number of goroutines to the WaitGroup
	wg.Add(numGoroutines)

	// Launch the goroutines
	for i := 0; i < numGoroutines; i++ {
		go simulateWork(&wg, i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All workers have finished. Main function exiting.")
}
