package main

import (
	"fmt"
	"time"
)

// Expensive operation we want to optimize
func slowFunction(n int) int {
	time.Sleep(100 * time.Millisecond) // Simulate a slow operation
	return n * n
}

// Create a closure-based cache
func createCachedFunction(f func(int) int) func(int) int {
	cache := make(map[int]int) // Internal state to store results

	return func(n int) int {
		if result, exists := cache[n]; exists {
			fmt.Printf("Fetching from cache for %d: %d\n", n, result)
			return result
		}
		result := f(n)
		cache[n] = result
		fmt.Printf("Computing and caching result for %d: %d\n", n, result)
		return result
	}
}

// Example use of closure in callbacks
func useFunctionWithCallback(n int, callback func(int) int) {
	fmt.Printf("Result for %d is: %d\n", n, callback(n))
}

func main() {
	// Create a cached version of the slow function
	cachedSlowFunction := createCachedFunction(slowFunction)

	// Demonstrate the use of closures in callbacks
	useFunctionWithCallback(5, cachedSlowFunction) // Computes and caches
	useFunctionWithCallback(5, cachedSlowFunction) // Fetches from cache
	useFunctionWithCallback(3, cachedSlowFunction) // Computes and caches
	useFunctionWithCallback(3, cachedSlowFunction) // Fetches from cache
	useFunctionWithCallback(5, cachedSlowFunction) // Fetches from cache
}
