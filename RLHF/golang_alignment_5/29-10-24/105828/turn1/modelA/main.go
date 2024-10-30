package main

import (
	"fmt"
)

// Shape interface with a Area() method
type Shape interface {
	Area() float64
}

// Rectangle struct implementing Shape interface
type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Circle struct implementing Shape interface
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// CalculateArea function using dynamic typing
func CalculateArea(shape Shape) float64 {
	return shape.Area()
}

func main() {
	// Usage with dynamic typing
	shapes := []interface{}{
		Rectangle{length: 10, width: 20},
		Circle{radius: 5},
	}

	totalArea := 0.0
	for _, shape := range shapes {
		area := CalculateArea(shape.(Shape)) // Type assertion required here
		fmt.Printf("Area of %T: %f\n", shape, area)
		totalArea += area
	}
	fmt.Println("Total Area:", totalArea)
}
