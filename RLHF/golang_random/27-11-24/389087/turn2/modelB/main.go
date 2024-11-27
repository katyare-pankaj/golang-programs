package main

import (
	"fmt"
)

// Define the interface Shape which requires all shapes to have an Area() float64 method
type Shape interface {
	Area() float64
}

// struct for Circle
type Circle struct {
	Radius float64
}

// Implement the Shape interface for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// struct for Square
type Square struct {
	Side float64
}

// Implement the Shape interface for Square
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// function to calculate the area of a given shape using the interface method
func calculateArea(shape Shape) float64 {
	return shape.Area()
}

func main() {
	circle := Circle{Radius: 5.5}
	square := Square{Side: 10.0}

	fmt.Println("Area of circle:", calculateArea(circle))
	fmt.Println("Area of square:", calculateArea(square))
}
