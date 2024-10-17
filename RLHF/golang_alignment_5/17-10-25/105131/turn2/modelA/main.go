// circle.go
package main

import (
	"fmt"
	"math"
)

// CalculateArea calculates the area of a circle given its radius.
func CalculateArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {
	radius := 5.5
	area := CalculateArea(radius)
	fmt.Printf("Area of the circle with radius %.2f is: %.2f\n", radius, area)
}
