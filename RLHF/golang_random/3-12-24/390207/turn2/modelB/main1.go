package main

import (
	"fmt"
	"sync"
)

// Counter is protected by a mutex to prevent race conditions
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment safely increments the counter
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
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

	// Incrementing the counter concurrently
	incrementCounter := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter.Increment()
		}
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go incrementCounter()
	}

	wg.Wait()
	fmt.Printf("Final Counter Value: %d\n", counter.Value())
}
