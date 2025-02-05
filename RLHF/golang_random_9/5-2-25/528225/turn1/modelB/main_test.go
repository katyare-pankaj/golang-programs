package main

import (
	"fmt"
	"testing"
)

// Factorial calculates the factorial of a non-negative integer n.
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Println("Factorial Test Results:")
	fmt.Println("Factorial(0) =", Factorial(0))   // Output: 1
	fmt.Println("Factorial(5) =", Factorial(5))   // Output: 120
	fmt.Println("Factorial(-3) =", Factorial(-3)) // Output: 0
}

func TestFactorial(t *testing.T) {
	// Initial test case: Factorial of 0 should be 1
	if Factorial(0) != 1 {
		t.Errorf("Factorial(0) failed. Expected: 1, Got: %d", Factorial(0))
	}

	// Add more test cases iteratively
	// Test case for Factorial of 5
	if Factorial(5) != 120 {
		t.Errorf("Factorial(5) failed. Expected: 120, Got: %d", Factorial(5))
	}

	// Test case for Factorial of a negative number
	if Factorial(-3) != 0 {
		t.Errorf("Factorial(-3) failed. Expected: 0, Got: %d", Factorial(-3))
	}

	// Add more test cases as needed
}
