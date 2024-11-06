package main

import (
	"fmt"
)

// Shape interface defines the operations that all shapes must support
type Shape interface {
	Draw()
	Clone() Shape
}

// Concrete Shape implementations
type Circle struct {
	radius float64
}

func (c *Circle) Draw() {
	fmt.Println("Drawing a circle with radius", c.radius)
}

func (c *Circle) Clone() Shape {
	return &Circle{radius: c.radius}
}

type Rectangle struct {
	length float64
	width  float64
}

func (r *Rectangle) Draw() {
	fmt.Println("Drawing a rectangle with length", r.length, "and width", r.width)
}

func (r *Rectangle) Clone() Shape {
	return &Rectangle{length: r.length, width: r.width}
}

// ShapeFactory uses the prototype pattern to create shapes
type ShapeFactory struct {
	prototypes map[string]Shape
}

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{
		prototypes: make(map[string]Shape),
	}
}

func (sf *ShapeFactory) RegisterPrototype(name string, prototype Shape) {
	sf.prototypes[name] = prototype
}

func (sf *ShapeFactory) CreateShape(name string) Shape {
	prototype, ok := sf.prototypes[name]
	if !ok {
		return nil
	}
	return prototype.Clone()
}

func main() {
	// Create a shape factory and register prototypes
	factory := NewShapeFactory()
	factory.RegisterPrototype("circle", &Circle{radius: 5.0})
	factory.RegisterPrototype("rectangle", &Rectangle{length: 10.0, width: 20.0})

	// Create new shapes using the factory
	circle1 := factory.CreateShape("circle")
	circle1.Draw() // Output: Drawing a circle with radius 5

	rectangle1 := factory.CreateShape("rectangle")
	rectangle1.Draw() // Output: Drawing a rectangle with length 10 and width 20

	// Modify the prototype and create new shapes
	circlePrototype := factory.prototypes["circle"].(*Circle)
	circlePrototype.radius = 10.0

	circle2 := factory.CreateShape("circle")
	circle2.Draw() // Output: Drawing a circle with radius 10
}
