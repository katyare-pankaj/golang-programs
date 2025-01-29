package main

import (
	"fmt"
	"sync"
	"time"
)

// Memoize accepts a function and returns a memoized version of it
func Memoize(fn func(int) int) func(int) int {
	cache := make(map[int]int)
	var mu sync.Mutex

	return func(n int) int {
		mu.Lock()
		defer mu.Unlock()

		// Check if the result is already cached
		if result, found := cache[n]; found {
			fmt.Printf("Cache hit for input: %d -> %d\n", n, result)
			return result
		}

		// Compute the result, cache it, and return
		result := fn(n)
		cache[n] = result
		fmt.Printf("Computed and cached for input: %d -> %d\n", n, result)
		return result
	}
}

// ExpensiveComputation simulates a time-consuming computation
func ExpensiveComputation(n int) int {
	fmt.Printf("Performing expensive computation for input: %d...\n", n)
	time.Sleep(2 * time.Second) // Simulate delay
	return n * n
}

func main() {
	// Create a memoized version of ExpensiveComputation
	memoizedComputation := Memoize(ExpensiveComputation)

	// Perform some computations
	fmt.Println("Result:", memoizedComputation(5)) // This will calculate and cache the result
	fmt.Println("Result:", memoizedComputation(5)) // This should return the cached result
	fmt.Println("Result:", memoizedComputation(6)) // This will calculate and cache the result
	fmt.Println("Result:", memoizedComputation(5)) // This should return the cached result again
	fmt.Println("Result:", memoizedComputation(6)) // This should return the cached result
	fmt.Println("Result:", memoizedComputation(7)) // This will calculate and cache the result
}
