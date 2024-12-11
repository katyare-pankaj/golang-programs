package counter

import (
	"sync"
)

// Counter is a simple implementation of a thread-safe counter.
type Counter struct {
	mu  sync.Mutex
	val int
}

// Increment safely increments the counter.
func (c *Counter) Increment() {
	c.mu.Lock()
	c.val++
	c.mu.Unlock()
}

// Value returns the current value of the counter.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}
