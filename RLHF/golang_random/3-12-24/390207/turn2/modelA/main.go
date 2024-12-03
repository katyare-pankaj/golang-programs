package main

import (
	"fmt"
	"sync"
)

// Counter is a struct that holds a count and a mutex to protect it
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment safely increments the counter
func (c *Counter) Increment() {
	c.mu.Lock()         // Acquire the lock
	defer c.mu.Unlock() // Ensure the lock is released

	c.value++
}

// Value safely retrieves the current value of the counter
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// Function that increments the counter 1000 times
	incrementCounter := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter.Increment()
		}
	}

	// Launch 10 goroutines to increment the counter concurrently
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go incrementCounter()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final counter value
	fmt.Printf("Final Counter Value: %d\n", counter.Value())
}
