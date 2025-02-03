package main

import "fmt"

func expensiveFunction(n int) int {
	// Simulate an expensive operation.  In a real case, this would do some useful work.
	if n < 0 {
		return 0
	}
	return n * n
}

// Memoize returns a closure that memoizes the results of the given function.
func Memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(key int) int {
		if result, ok := cache[key]; ok {
			return result // If the result is in the cache, return it
		}
		result := f(key)    // Otherwise, call the original function
		cache[key] = result // Store the result in the cache
		return result
	}
}

func main() {
	memoizedFunction := Memoize(expensiveFunction)

	result := memoizedFunction(10)
	fmt.Println("Memoized result:", result) // Output: Memoized result: 100

	result = memoizedFunction(10)
	fmt.Println("Reusing cached result:", result) // Output: Reusing cached result: 100
}
