package main

import "fmt"

// calculateSum calculates the sum of two numbers
func calculateSum(a float64, b float64) float64 {
	return a + b
}

func main() {
	sum := calculateSum(10.0, 20.0)
	fmt.Println("Sum:", sum)
}
