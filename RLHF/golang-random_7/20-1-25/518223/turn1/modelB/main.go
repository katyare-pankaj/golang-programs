package main

import "fmt"

// Shape interface defines a method to calculate area
type Shape interface {
	Area() float64
}

// Rectangle struct implements the Shape interface
type Rectangle struct {
	length float64
	width  float64
}

// Area method calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Circle struct implements the Shape interface
type Circle struct {
	radius float64
}

// Area method calculates the area of a circle
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// DisplayAreaFunction is a functional interface that represents a function taking a Shape and printing its area
type DisplayAreaFunction func(Shape)

// PrintArea function takes a DisplayAreaFunction as an argument and applies it to each shape in a slice
func PrintArea(shapes []Shape, displayArea DisplayAreaFunction) {
	for _, shape := range shapes {
		displayArea(shape)
	}
}

// main function demonstrates the use of functional interfaces
func main() {
	// Create a slice of shapes
	shapes := []Shape{Rectangle{length: 5, width: 3}, Circle{radius: 4}}

	// Use an anonymous function as a callback to calculate and display the area of each shape
	PrintArea(shapes, func(shape Shape) {
		fmt.Printf("Area of %T: %f\n", shape, shape.Area())
	})

	// Alternatively, you can define a named function as a callback
	displayAreaByType := func(shape Shape) {
		fmt.Printf("Area of %T: %f\n", shape, shape.Area())
	}
	PrintArea(shapes, displayAreaByType)
}
