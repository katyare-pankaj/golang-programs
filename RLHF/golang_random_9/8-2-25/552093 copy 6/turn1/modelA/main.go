package main

import (
	"fmt"
	"sync"
)

// Global variable to simulate shared state.
var counter int

// Mutex to ensure safe access to the shared state.
var mutex sync.Mutex

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure that Done is called when the goroutine completes

	mutex.Lock()
	fmt.Printf("Worker %d is incrementing the counter.\n", id)
	counter++
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5

	for i := 0; i < numWorkers; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Printf("Final counter value: %d\n", counter)
}
