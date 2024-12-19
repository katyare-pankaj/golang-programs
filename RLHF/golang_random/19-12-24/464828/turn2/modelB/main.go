package main

import (
	"fmt"
)

// function to calculate area and circumference of a circle
func circleCalc(radius float64) (area float64, circumference float64) {
	const pi = 3.14
	area = pi * radius * radius
	circumference = 2 * pi * radius
	return
}

func main() {
	// using named return values
	radius := 5.5
	a, c := circleCalc(radius)
	fmt.Printf("Area of circle with radius %.2f is: %.2f\n", radius, a)
	fmt.Printf("Circumference of circle with radius %.2f is: %.2f\n", radius, c)

	// using unnamed return values (not recommended)
	_, _ = circleCalc(3.0)

	// Using named return values again, to make the code more readable and understandable.
	_, circumferenceOfNewCircle := circleCalc(4.0)
	fmt.Printf("Circumference of circle with radius 4.0 is: %.2f\n", circumferenceOfNewCircle)

}
