package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()   // Lock the mutex
		counter++   // Increment the counter
		mu.Unlock() // Unlock the mutex
	}
}

func main() {
	var wg sync.WaitGroup

	// Start 10 goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final counter value:", counter) // Output will be consistent
}
