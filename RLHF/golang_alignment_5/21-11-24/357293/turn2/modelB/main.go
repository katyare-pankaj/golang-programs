package main

import "fmt"

// A dictionary (map) that stores previously calculated Fibonacci numbers
var memo map[int]int = map[int]int{}

// A memoized recursive Fibonacci function
func fib(n int) int {
	// Base case: The first two Fibonacci numbers are 0 and 1.
	if n < 2 {
		return n
	}

	// Check if the answer is already stored in the memo
	if answer, ok := memo[n]; ok {
		return answer
	}

	// Recursive case: Calculate the Fibonacci number by calling the function with (n-1) and (n-2)
	result := fib(n-1) + fib(n-2)

	// Store the result in the memo for future reference
	memo[n] = result

	return result
}

func main() {
	// Driver Code
	n := 20
	fmt.Println("Nth Fibonacci Number: ", fib(n))

	// Test cases
	fmt.Println("fib(0) =", fib(0))   // Output: fib(0) = 0
	fmt.Println("fib(1) =", fib(1))   // Output: fib(1) = 1
	fmt.Println("fib(9) =", fib(9))   // Output: fib(9) = 34
	fmt.Println("fib(30) =", fib(30)) // Output: fib(30) = 832040
}
