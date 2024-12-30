package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate work by sleeping for a random duration
	sleepDuration := time.Duration(100+id*50) * time.Millisecond
	time.Sleep(sleepDuration)
	fmt.Printf("Worker %d finished after sleeping for %s\n", id, sleepDuration)
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 5

	// Add the number of goroutines to the WaitGroup
	wg.Add(numGoroutines)

	// Start goroutines
	for i := 0; i < numGoroutines; i++ {
		go worker(i, &wg)
	}

	// Wait for all goroutines to complete before exiting the main function
	wg.Wait()

	fmt.Println("All workers have finished their tasks.")
}
