package main

import (
	"fmt"
	"math"
)

func main() {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	sum := 0.0

	// Using built-in Sum function for better performance
	for _, value := range data {
		sum += value
	}
	fmt.Println("Sum:", sum)

	// Using built-in Pow function
	for i, value := range data {
		data[i] = math.Pow(value, 2)
	}
	fmt.Println("Modified Data:", data)
}
