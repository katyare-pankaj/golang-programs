package main

import "fmt"

// Shape is an interface that defines the behavior of shapes.
type Shape interface {
	Area() float64
}

// Circle is a struct that implements the Shape interface.
type Circle struct {
	Radius float64
}

// Area calculates the area of the circle.
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Rectangle is a struct that implements the Shape interface.
type Rectangle struct {
	Length float64
	Width  float64
}

// Area calculates the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Triangle is a struct that implements the Shape interface.
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates the area of the triangle.
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5.0},
		Rectangle{Length: 10.0, Width: 20.0},
		Triangle{Base: 6.0, Height: 7.0},
	}

	for _, shape := range shapes {
		fmt.Printf("Area of %T is %.2f\n", shape, shape.Area())
	}
}
