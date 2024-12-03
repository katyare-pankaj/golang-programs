package main

import (
	"fmt"
	"sync"
	"time"
)

// SharedResource represents a simple shared resource.
type SharedResource struct {
	count int
	mu    *sync.Mutex // Mutex to protect access to the count.
}

func newSharedResource() *SharedResource {
	return &SharedResource{
		count: 0,
		mu:    &sync.Mutex{},
	}
}

// Increment increments the count of the shared resource.
func (sr *SharedResource) Increment() {
	sr.mu.Lock() // Acquire the mutex.
	sr.count++
	sr.mu.Unlock() // Release the mutex.
}

// Decrement decrements the count of the shared resource.
func (sr *SharedResource) Decrement() {
	sr.mu.Lock() // Acquire the mutex.
	if sr.count > 0 {
		sr.count--
	}
	sr.mu.Unlock() // Release the mutex.
}

// GetCount returns the current count of the shared resource.
func (sr *SharedResource) GetCount() int {
	sr.mu.RLock() // Acquire the read lock.
	count := sr.count
	sr.mu.RUnlock() // Release the read lock.
	return count
}

func main() {
	// Create a shared resource.
	resource := newSharedResource()

	// Create worker goroutines to increment and decrement the resource.
	numWorkers := 100
	numIterations := 1000

	wg := &sync.WaitGroup{}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numIterations; j++ {
				resource.Increment()
				resource.Decrement()
				// Simulate some work
				time.Sleep(time.Microsecond)
			}
		}(i)
	}

	// Wait for all worker goroutines to complete.
	wg.Wait()

	// Print the final count.
	fmt.Printf("Final count: %d\n", resource.GetCount())
}
