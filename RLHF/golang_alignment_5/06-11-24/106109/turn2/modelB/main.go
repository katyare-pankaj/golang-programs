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
	return &ShapeFactory{prototypes: make(map[string]Shape)}
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

// TestShapeFactory uses the prototype pattern to create test shapes
type TestShapeFactory struct {
	ShapeFactory
}

func NewTestShapeFactory() *TestShapeFactory {
	tf := NewShapeFactory()
	// Register test prototypes
	tf.RegisterPrototype("test_circle", &Circle{radius: 1.0})
	tf.RegisterPrototype("test_rectangle", &Rectangle{length: 2.0, width: 3.0})
	return &TestShapeFactory{*tf}
}

func main() {
	// Create a shape factory and register prototypes
	factory := NewShapeFactory()
	factory.RegisterPrototype("circle", &Circle{radius: 5.0})
	factory.RegisterPrototype("rectangle", &Rectangle{length: 10.0, width: 20.0})

	// Use the factory to create shapes
	circle := factory.CreateShape("circle")
	circle.Draw() // Output: Drawing a circle with radius 5.0

	// Create test shapes using the TestShapeFactory
	testFactory := NewTestShapeFactory()
	testCircle := testFactory.CreateShape("test_circle")
	testCircle.Draw() // Output: Drawing a circle with radius 1.0

	// Perform tests on test shapes
	// ...
}
