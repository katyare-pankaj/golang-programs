package main

import (
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func calculateArea(shape interface{}) float64 {
	switch shape.(type) {
	case Point:
		return 0 // Area of a point is 0
	case Rectangle:
		rect := shape.(Rectangle)
		return rect.Width * rect.Height
	default:
		panic("Unsupported shape type")
	}
}

func main() {
	point := Point{X: 1.0, Y: 2.0}
	rect := Rectangle{Width: 3.0, Height: 4.0}

	fmt.Println("Area of point:", calculateArea(point))    // Output: Area of point: 0
	fmt.Println("Area of rectangle:", calculateArea(rect)) // Output: Area of rectangle: 12
}
