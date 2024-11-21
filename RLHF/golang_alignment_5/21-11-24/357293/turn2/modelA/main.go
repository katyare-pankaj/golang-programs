package main

import (
	"fmt"
	"testing"
)

// memoizedFibonacci uses a map to store the computed Fibonacci numbers for memoization.
var memoizedFibonacci = make(map[int]int)

func fibonacci(n int) int {
	if result, ok := memoizedFibonacci[n]; ok {
		return result // If the result is already computed, return it from the map
	}

	// Base cases
	var result int
	if n < 2 {
		result = n
	} else {
		result = fibonacci(n-1) + fibonacci(n-2)
	}

	// Store the result in the map before returning
	memoizedFibonacci[n] = result
	return result
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{19, 4181}, // Large input to test the memoization
	}

	for _, test := range tests {
		got := fibonacci(test.n)
		if got != test.want {
			t.Errorf("fibonacci(%d) = %d, want %d", test.n, got, test.want)
		}
	}
}

// Main function to demonstrate the memoized Fibonacci function.
func main() {
	// Example usage
	n := 19
	fmt.Println("Fibonacci(", n, ") =", fibonacci(n))

	// Run tests
	testing.Run("fibonacciTest", TestFibonacci)
}
