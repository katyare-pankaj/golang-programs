package main

import (
	"fmt"
)

// A function that calculates the area and perimeter of a rectangle using named return values
func calculateRectangleProperties(width float64, height float64) (area float64, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}

func main() {
	// Example usage
	width, height := 5.5, 3.0
	area, perimeter := calculateRectangleProperties(width, height)

	fmt.Printf("Width: %.1f\n", width)
	fmt.Printf("Height: %.1f\n", height)
	fmt.Printf("Area: %.2f\n", area)
	fmt.Printf("Perimeter: %.2f\n", perimeter)
}
