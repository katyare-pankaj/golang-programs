package main

import "fmt"

// MemoizedFibonacci uses memoization to store pre-computed Fibonacci numbers
func MemoizedFibonacci(n int, memo map[int]int) int {
	// Base case: The first two Fibonacci numbers are 0 and 1.
	if n < 2 {
		return n
	}

	// If the result is already memoized, return it
	if result, ok := memo[n]; ok {
		return result
	}

	// Recursive case: Calculate the n-th Fibonacci number by calling the function with (n-1) and (n-2).
	result := MemoizedFibonacci(n-1, memo) + MemoizedFibonacci(n-2, memo)

	// Memoize the result
	memo[n] = result

	return result
}

func main() {
	//big data n
	n := 100
	// Create a memoization map
	memo := make(map[int]int)

	// Measure the time taken without memoization
	// ... (code not shown for brevity)

	// Calculate the Fibonacci number using memoization
	result := MemoizedFibonacci(n, memo)

	fmt.Println("Fibonacci(", n, ") =", result)

	// Measure the time taken with memoization
	// ... (code not shown for brevity)
}
