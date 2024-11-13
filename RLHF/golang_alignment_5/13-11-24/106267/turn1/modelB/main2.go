package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Point struct {
	X float64
	Y float64
}

func (p Point) Area() float64 {
	return 0 // Area of a point is 0
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func calculateArea(shape Shape) float64 {
	return shape.Area()
}

func main() {
	point := Point{X: 1.0, Y: 2.0}
	rect := Rectangle{Width: 3.0, Height: 4.0}

	fmt.Println("Area of point:", calculateArea(point))    // Output: Area of point: 0
	fmt.Println("Area of rectangle:", calculateArea(rect)) // Output: Area of rectangle: 12
}
