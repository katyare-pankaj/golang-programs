package main

import (
	"fmt"
)

// Define the Product interface
type Product interface {
	Use()
	Clone() Product
}

// Concrete Product implementations
type ConcreteProductA struct {
	name string
}

func (p *ConcreteProductA) Use() {
	fmt.Println("Using ConcreteProduct A")
}

func (p *ConcreteProductA) Clone() Product {
	return &ConcreteProductA{name: p.name}
}

type ConcreteProductB struct {
	name string
}

func (p *ConcreteProductB) Use() {
	fmt.Println("Using ConcreteProduct B")
}

func (p *ConcreteProductB) Clone() Product {
	return &ConcreteProductB{name: p.name}
}

// Manager class uses the prototype pattern
type Manager struct {
	prototypes map[string]Product
}

func NewManager() *Manager {
	return &Manager{
		prototypes: make(map[string]Product),
	}
}

func (m *Manager) RegisterPrototype(name string, prototype Product) {
	m.prototypes[name] = prototype
}

func (m *Manager) CreateProduct(name string) Product {
	p, ok := m.prototypes[name]
	if !ok {
		return nil
	}
	return p.Clone()
}

func main() {
	manager := NewManager()

	// Register prototypes
	manager.RegisterPrototype("productA", &ConcreteProductA{name: "Default A"})
	manager.RegisterPrototype("productB", &ConcreteProductB{name: "Default B"})

	// Creating products via cloning
	product1 := manager.CreateProduct("productA")
	product1.Use() // Output: Using ConcreteProduct A

	product2 := manager.CreateProduct("productB")
	product2.Use() // Output: Using ConcreteProduct B

	// Modify the prototype and create new instances
	productAPrototype := manager.prototypes["productA"].(*ConcreteProductA)
	productAPrototype.name = "Custom A"

	product3 := manager.CreateProduct("productA")
	product3.Use() // Output: Using ConcreteProduct A with name Custom A
}
