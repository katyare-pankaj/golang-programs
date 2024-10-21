package main

import (
	"fmt"
)

// CalculateSum returns the sum of two integers.
// We should add comments to explain the purpose of the function and its parameters
func CalculateSum(a int, b int) int {
	// Use meaningful variable names
	total := a + b
	return total
}

func main() {
	// Provide context for the input values
	num1 := 3
	num2 := 5
	result := CalculateSum(num1, num2)
	fmt.Println("Sum of", num1, "and", num2, "is:", result)
}
