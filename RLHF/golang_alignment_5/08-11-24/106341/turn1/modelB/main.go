package main

import "fmt"

func fibonacciDP(n int) int {
	// Create a slice to store pre-computed Fibonacci numbers
	dp := make([]int, n+1)

	// Base case: The first two Fibonacci numbers are 0 and 1
	dp[0] = 0
	dp[1] = 1

	// Recursive case: Use pre-computed values to calculate Fibonacci numbers
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func main() {

	fib := fibonacciDP(100)
	fmt.Println("Fibonacci :", fib)
}
