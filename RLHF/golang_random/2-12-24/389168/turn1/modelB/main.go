package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// Define your numbers
	a, b := 5.2, 3.8

	// Create a matrix with these numbers
	data := []float64{a, b}
	matrix := mat.NewDense(1, len(data), data) // A 1x2 matrix

	// Calculate the sum using Gonum's methods
	sum := mat.Sum(matrix)

	// Print the result
	fmt.Printf("The sum is: %.2f\n", sum)
}
