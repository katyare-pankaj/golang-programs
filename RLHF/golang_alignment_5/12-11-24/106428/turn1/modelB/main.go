package main

import (
	"fmt"
	"time"
)

func factorial(n int) int {
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
	// Function call to measure the execution time and identify hotspots
	start := time.Now()
	factorial(1000)
	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed)
}
