package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// performTask simulates a task that takes a random amount of time
func performTask(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	// Indicate the start of the goroutine
	fmt.Printf("Goroutine %d starting\n", id)

	// Simulate a random workload
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// Indicate the completion of the goroutine
	fmt.Printf("Goroutine %d completed\n", id)
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 5

	// Add the number of goroutines to the WaitGroup counter
	wg.Add(numGoroutines)

	for i := 1; i <= numGoroutines; i++ {
		// Launch goroutines
		go performTask(i, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("All goroutines have completed")
}
