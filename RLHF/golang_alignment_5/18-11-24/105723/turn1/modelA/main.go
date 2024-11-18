package main

import (
	"fmt"
)

// Factorial function with tail recursion
func factorial(n int) int {
	result := 1
	return factorialHelper(n, &result)
}

func factorialHelper(n int, result *int) int {
	if n == 0 {
		return *result
	}
	*result *= n
	return factorialHelper(n-1, result)
}

func main() {
	n := 7
	fmt.Println("Factorial of", n, "is:", factorial(n))
}
