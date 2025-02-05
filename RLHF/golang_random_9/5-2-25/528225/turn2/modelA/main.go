// main.go

package main

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
	// This main function can be used for manual testing if needed
}
