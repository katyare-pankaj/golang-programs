package main

import (
	"fmt"
	"math"
)

// Pure function to calculate the area of a circle
func calculateArea(radius float64) float64 {
	return math.Pi * radius * radius
}

// Higher-order function to apply a transformation to a list of numbers
func mapFloat64(numbers []float64, transform func(float64) float64) []float64 {
	result := make([]float64, len(numbers))
	for i, num := range numbers {
		result[i] = transform(num)
	}
	return result
}

// Function to calculate the areas of multiple circles
func calculateAreas(radii []float64) []float64 {
	return mapFloat64(radii, calculateArea)
}

func main() {
	radii := []float64{1.0, 2.0, 3.0}
	areas := calculateAreas(radii)
	fmt.Println("Areas:", areas) // Output: Areas: [3.14 12.57 28.27]
}
