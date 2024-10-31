package main

import "fmt"

// factorial calculates the factorial of a number
func factorial(n int) int {
	if n < 0 {
		return 0
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	number := 6
	result := factorial(number)
	fmt.Println("Factorial of", number, "is:", result)
}
