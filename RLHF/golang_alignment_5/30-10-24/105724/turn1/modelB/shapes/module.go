// shapes/module.go

package shapes

import (
	"math"
)

// Rectangle represents a rectangle shape.
type Rectangle struct {
	Length float64
	Width  float64
}

// Area calculates the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Circle represents a circle shape.
type Circle struct {
	Radius float64
}

// Area calculates the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
