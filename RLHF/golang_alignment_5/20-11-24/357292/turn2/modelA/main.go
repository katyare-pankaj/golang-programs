package main

import (
	"fmt"
	"time"
)

// FactorialRecursive calculates the factorial of a non-negative integer n using recursion.
func FactorialRecursive(n int) int {
	if n < 0 {
		panic("Factorial is not defined for negative numbers")
	}
	if n == 0 || n == 1 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}

// FactorialIterative calculates the factorial of a non-negative integer n using iteration.
func FactorialIterative(n int) int {
	if n < 0 {
		panic("Factorial is not defined for negative numbers")
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	// Set a number to calculate the factorial of
	n := 10

	// Measure the execution time of the recursive function
	startTime := time.Now()
	resultRecursive := FactorialRecursive(n)
	endTime := time.Now()
	elapsedTimeRecursive := endTime.Sub(startTime)
	fmt.Printf("Recursive Factorial of %d: %d, Time: %s\n", n, resultRecursive, elapsedTimeRecursive)

	// Measure the execution time of the iterative function
	startTime = time.Now()
	resultIterative := FactorialIterative(n)
	endTime = time.Now()
	elapsedTimeIterative := endTime.Sub(startTime)
	fmt.Printf("Iterative Factorial of %d: %d, Time: %s\n", n, resultIterative, elapsedTimeIterative)
}
