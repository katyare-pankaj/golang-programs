package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate work by sleeping for a random duration
	sleepDuration := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(sleepDuration)
	fmt.Printf("Worker %d finished after sleeping for %s\n", id, sleepDuration)
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 10
	// Add the number of goroutines to the WaitGroup
	wg.Add(numGoroutines)

	// Start goroutines
	for i := 0; i < numGoroutines; i++ {
		go worker(i, &wg)
	}

	// Now, the main goroutine will block until all worker goroutines have called Done.
	fmt.Println("Main goroutine waiting for workers to finish...")
	wg.Wait()

	fmt.Println("All workers have finished their tasks. Main function exiting now.")
}
