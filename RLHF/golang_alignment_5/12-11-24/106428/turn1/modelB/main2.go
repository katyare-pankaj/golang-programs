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

	// Unnecessary loop to introduce a hotspot
	for i := 0; i < 10000000; i++ {
	}
	return result
}

func main() {
	start := time.Now()
	factorial(1000)
	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed)
}
