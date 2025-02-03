package main

import (
	"fmt"
	"time"
)

// Expensive function to be memoized
func expensiveOperation(n int) int {
	time.Sleep(100 * time.Millisecond) // Simulate a slow operation
	return n * n
}

// Memoize function using closures
func memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)

	return func(n int) int {
		if result, found := cache[n]; found {
			fmt.Printf("Returning cached result for %d: %d\n", n, result)
			return result
		}
		result := f(n)
		cache[n] = result
		fmt.Printf("Computing and caching result for %d: %d\n", n, result)
		return result
	}
}

func main() {
	// Create a memoized version of the expensive function
	memoizedExpensiveOperation := memoize(expensiveOperation)

	// Test memoization
	fmt.Println("Result:", memoizedExpensiveOperation(10)) // Computes and caches
	fmt.Println("Result:", memoizedExpensiveOperation(10)) // Returns cached
	fmt.Println("Result:", memoizedExpensiveOperation(5))  // Computes and caches
	fmt.Println("Result:", memoizedExpensiveOperation(5))  // Returns cached
	fmt.Println("Result:", memoizedExpensiveOperation(10)) // Returns cached
}
