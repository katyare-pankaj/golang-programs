package main

import (
	"fmt"
	"time"
)

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
func main() {
	start := time.Now()
	result := factorial(1000000)
	elapsed := time.Since(start)
	fmt.Printf("Factorial of 1000000: %d\n", result)
	fmt.Printf("Execution time: %s\n", elapsed)
}
