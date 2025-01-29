package main

import (
	"fmt"
	"sync"
)

type memoizedFunc func(int) int

func memoize(f memoizedFunc) memoizedFunc {
	cache := make(map[int]int)
	mu := &sync.Mutex{}
	return func(n int) int {
		mu.Lock()
		defer mu.Unlock()
		if result, ok := cache[n]; ok {
			return result // Return the cached result if available
		}
		result := f(n) // Calculate the result and store it in the cache
		cache[n] = result
		return result
	}
}

// Factorial is an expensive computation to be memoized
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func main() {
	memoizedFactorial := memoize(Factorial)

	// Test the memoized factorial function
	fmt.Println("Factorial of 5:", memoizedFactorial(5))   // Output: 120
	fmt.Println("Factorial of 5:", memoizedFactorial(5))   // Output: 120 (Cached result)
	fmt.Println("Factorial of 10:", memoizedFactorial(10)) // Output: 3628800
	fmt.Println("Factorial of 10:", memoizedFactorial(10)) // Output: 3628800 (Cached result)
}
