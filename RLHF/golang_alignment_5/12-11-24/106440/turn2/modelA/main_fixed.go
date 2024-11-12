package main

import (
	"fmt"
	"sync"
)

// SharedResource represents a resource accessed by multiple goroutines
type SharedResource struct {
	mu      sync.Mutex
	counter int
}

func (r *SharedResource) Increment() {
	r.mu.Lock()
	r.counter++
	r.mu.Unlock()
}

func main() {
	resource := &SharedResource{}

	// Goroutines simulating concurrent work
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) // Add to the WaitGroup before launching the goroutine
		go func() {
			defer wg.Done() // Signal that the goroutine is done when it finishes
			for j := 0; j < 1000; j++ {
				resource.Increment()
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final Counter Value:", resource.counter) // Expected output: 10000
}
