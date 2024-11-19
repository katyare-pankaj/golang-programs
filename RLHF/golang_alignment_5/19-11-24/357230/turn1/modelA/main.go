package main

import (
	"fmt"
)

// Define the Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define the Rectangle struct that implements the Shape interface
type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

// Define the Circle struct that implements the Shape interface
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// Function to calculate the total area of a list of shapes
func calculateTotalArea(shapes []Shape) float64 {
	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

// Function to calculate the total perimeter of a list of shapes
func calculateTotalPerimeter(shapes []Shape) float64 {
	totalPerimeter := 0.0
	for _, shape := range shapes {
		totalPerimeter += shape.Perimeter()
	}
	return totalPerimeter
}

func main() {
	// Create instances of Rectangle and Circle
	rect := Rectangle{length: 5.0, width: 10.0}
	circle := Circle{radius: 4.0}

	// Create a slice of Shape to hold different shapes
	shapes := []Shape{rect, circle}

	// Calculate the total area and perimeter
	totalArea := calculateTotalArea(shapes)
	totalPerimeter := calculateTotalPerimeter(shapes)

	// Output the results
	fmt.Printf("Total area: %.2f\n", totalArea)
	fmt.Printf("Total perimeter: %.2f\n", totalPerimeter)
}
