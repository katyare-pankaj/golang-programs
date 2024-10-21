package main

import (
	"fmt"
)

// CalculateSum returns the sum of two integers.
func CalculateSum(a int, b int) int {
	// Check for integer overflow
	if a > 0 && b > (INT_MAX-a) {
		// Handle overflow scenario
		return INT_MAX
	} else if a < 0 && b < (INT_MIN-a) {
		// Handle underflow scenario
		return INT_MIN
	}
	return a + b
}

const (
	INT_MAX = int(^uint(0) >> 1)
	INT_MIN = -INT_MAX - 1
)

func main() {
	result := CalculateSum(3, 5)
	fmt.Println("Sum:", result)

	// Test for overflow and underflow
	overflowResult := CalculateSum(INT_MAX, 1)
	underflowResult := CalculateSum(INT_MIN, -1)
	fmt.Println("Overflow Result:", overflowResult)
	fmt.Println("Underflow Result:", underflowResult)
}
