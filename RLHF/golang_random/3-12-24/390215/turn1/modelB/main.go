package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Shared counter
	counter := 0

	// Number of goroutines to spawn
	goroutines := 100

	wg.Add(goroutines)

	// Increment function that is run by each goroutine
	increment := func() {
		defer wg.Done()
		mu.Lock()
		counter = counter + 1
		mu.Unlock()
	}

	// Start multiple goroutines
	for i := 0; i < goroutines; i++ {
		go increment()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Printf("Final Counter: %d\n", counter)
}
