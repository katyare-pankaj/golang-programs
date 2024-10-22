package main

import "fmt"

// Shape interface
type Shape interface {
	draw()
}

// Rectangle struct
type Rectangle struct{}

func (r *Rectangle) draw() {
	fmt.Println("Drawing a rectangle.")
}

// Circle struct
type Circle struct{}

func (c *Circle) draw() {
	fmt.Println("Drawing a circle.")
}

// ShapeFactory interface
type ShapeFactory interface {
	createShape() Shape
}

// RectangleFactory struct
type RectangleFactory struct{}

func (rf *RectangleFactory) createShape() Shape {
	return &Rectangle{}
}

// CircleFactory struct
type CircleFactory struct{}

func (cf *CircleFactory) createShape() Shape {
	return &Circle{}
}

func main() {
	var factory ShapeFactory

	factory = &RectangleFactory{} // Change factory to create different shapes
	shape := factory.createShape()
	shape.draw() // Output: Drawing a rectangle.

	factory = &CircleFactory{}
	shape = factory.createShape()
	shape.draw() // Output: Drawing a circle.
}
