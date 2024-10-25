package main

import (
	"fmt"
	"sync"
)

// A LazyMap is a map that computes values lazily on demand.
type LazyMap struct {
	mu      sync.Mutex
	m       map[int]int
	compute func(int) int
}

// NewLazyMap creates a new LazyMap with the given computation function.
func NewLazyMap(compute func(int) int) *LazyMap {
	return &LazyMap{m: make(map[int]int), compute: compute}
}

// Get retrieves the value for the given key from the map.
// The value is computed lazily if it hasn't been computed before.
func (lm *LazyMap) Get(key int) int {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	if value, ok := lm.m[key]; ok {
		return value // Value has already been computed
	}

	value := lm.compute(key)
	lm.m[key] = value // Store the computed value
	return value
}

func main() {
	// Example usage: Real-time data processing with lazy evaluation
	dataStream := []int{1, 2, 3, 4, 5}

	// Define a computationally expensive function to be applied lazily
	expensiveComputation := func(x int) int {
		fmt.Printf("Computing value for %d...\n", x)
		// Simulate expensive computation
		result := 0
		for i := 0; i < 1000000; i++ {
			result += x * i
		}
		return result
	}

	lazyMap := NewLazyMap(expensiveComputation)

	for _, key := range dataStream {
		result := lazyMap.Get(key)
		fmt.Printf("Result for %d: %d\n", key, result)
	}
}
