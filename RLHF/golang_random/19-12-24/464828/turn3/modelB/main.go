package main

import "fmt"

// function to calculate area and perimeter of a rectangle
func calculateRectangleProperties(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}

func main() {
	width := 5.0
	height := 3.0
	area, perimeter := calculateRectangleProperties(width, height)

	// Using named return values makes the code more readable
	fmt.Println("Area:", area)
	fmt.Println("Perimeter:", perimeter)

	// Now, let's see what happens if we swap the assignment order
	perimeter, area = calculateRectangleProperties(width, height)

	// Oops! We accidentally swapped the values.
	// Using named return values helps catch this error early
	fmt.Println("\nSwapped Values:")
	fmt.Println("Area:", area)           // Output: Perimeter: 16.0
	fmt.Println("Perimeter:", perimeter) // Output: Area: 15.0
}
