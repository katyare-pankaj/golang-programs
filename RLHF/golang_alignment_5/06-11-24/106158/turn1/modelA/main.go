package main

import (
	"fmt"
)

// Shape interface defines the operations that can be performed on shapes
type Shape interface {
	Clone() Shape
	Draw()
}

// Concrete Shape implementations

type Circle struct {
	radius float64
}

func (c *Circle) Clone() Shape {
	return &Circle{radius: c.radius}
}

func (c *Circle) Draw() {
	fmt.Println("Drawing a circle with radius:", c.radius)
}

type Rectangle struct {
	length float64
	width  float64
}

func (r *Rectangle) Clone() Shape {
	return &Rectangle{length: r.length, width: r.width}
}

func (r *Rectangle) Draw() {
	fmt.Println("Drawing a rectangle with length:", r.length, "and width:", r.width)
}

// ShapeManager using the Prototype pattern
type ShapeManager struct {
	prototypes map[string]Shape
}

func NewShapeManager() *ShapeManager {
	return &ShapeManager{
		prototypes: make(map[string]Shape),
	}
}

func (sm *ShapeManager) RegisterPrototype(name string, prototype Shape) {
	sm.prototypes[name] = prototype
}

func (sm *ShapeManager) CreateShape(name string) Shape {
	p, ok := sm.prototypes[name]
	if !ok {
		return nil
	}
	return p.Clone()
}

func main() {
	manager := NewShapeManager()

	// Register prototypes
	manager.RegisterPrototype("circle", &Circle{radius: 5.0})
	manager.RegisterPrototype("rectangle", &Rectangle{length: 10.0, width: 20.0})

	// Create shapes from prototypes
	circle1 := manager.CreateShape("circle")
	circle1.Draw() // Output: Drawing a circle with radius: 5

	rectangle1 := manager.CreateShape("rectangle")
	rectangle1.Draw() // Output: Drawing a rectangle with length: 10 and width: 20

	// Modify the prototype and create new shapes
	manager.prototypes["circle"].(*Circle).radius = 10.0
	circle2 := manager.CreateShape("circle")
	circle2.Draw() // Output: Drawing a circle with radius: 10
}
