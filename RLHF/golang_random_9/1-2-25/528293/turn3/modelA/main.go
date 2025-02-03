package main

import (
	"fmt"
	"sync"
)

// A simulation of an expensive computation
func expensiveComputation(x int) int {
	// Simulating a complex calculation
	return x * x
}

// Create a cache mechanism with synchronization using closures
func createCache(f func(int) int) func(int) int {
	cache := make(map[int]int)
	var mu sync.Mutex

	return func(x int) int {
		mu.Lock()
		defer mu.Unlock()

		// Check if the result is already in the cache
		if result, found := cache[x]; found {
			fmt.Printf("Cache hit for input %d: %d\n", x, result)
			return result
		}

		// If not cached, compute and store the result
		result := f(x)
		cache[x] = result
		fmt.Printf("Cache miss. Computed and cached result for %d: %d\n", x, result)
		return result
	}
}

func main() {
	// Create a cached version of the expensive computation
	cachedComputation := createCache(expensiveComputation)

	// Test the cached computation
	fmt.Println("Result:", cachedComputation(10)) // Computes and caches
	fmt.Println("Result:", cachedComputation(10)) // Returns cached
	fmt.Println("Result:", cachedComputation(5))  // Computes and caches
	fmt.Println("Result:", cachedComputation(10)) // Returns cached
	fmt.Println("Result:", cachedComputation(5))  // Returns cached
}
