// Example usage:
package main

import (
	"fmt"
)

type Observable interface {
	RegisterObserver(Observer)
	RemoveObserver(Observer)
	NotifyObservers()
}

type Observer interface {
	Update()
}

type ConcreteObservable struct {
	observers []Observer
	data      string
}

func (c *ConcreteObservable) RegisterObserver(o Observer) {
	c.observers = append(c.observers, o)
}

func (c *ConcreteObservable) RemoveObserver(o Observer) {
	for i, observer := range c.observers {
		if observer == o {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			return
		}
	}
}

func (c *ConcreteObservable) NotifyObservers() {
	for _, observer := range c.observers {
		observer.Update()
	}
}

func (c *ConcreteObservable) SetData(data string) {
	c.data = data
	c.NotifyObservers()
}

func (c *ConcreteObservable) GetData() string {
	return c.data
}

type ConcreteObserverA struct{}

func (a *ConcreteObserverA) Update() {
	fmt.Println("Observer A updated:")
}

type ConcreteObserverB struct{}

func (b *ConcreteObserverB) Update() {
	fmt.Println("Observer B updated:")
}

// Strategy Pattern
type Strategy interface {
	DoSomething()
}

type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) DoSomething() {
	fmt.Println("Strategy A is executed.")
}

type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) DoSomething() {
	fmt.Println("Strategy B is executed.")
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c *Context) DoSomething() {
	c.strategy.DoSomething()
}

// Factory Pattern
type Component interface {
	Operation()
}

type ConcreteComponentA struct{}

func (c *ConcreteComponentA) Operation() {
	fmt.Println("Component A is doing operation.")
}

type ConcreteComponentB struct{}

func (c *ConcreteComponentB) Operation() {
	fmt.Println("Component B is doing operation.")
}

type ComponentFactory interface {
	CreateComponent() Component
}

type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateComponent() Component {
	return &ConcreteComponentA{}
}

type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateComponent() Component {
	return &ConcreteComponentB{}
}

func main() {
	// Observer Pattern
	observable := &ConcreteObservable{}
	observerA := &ConcreteObserverA{}
	observerB := &ConcreteObserverB{}

	observable.RegisterObserver(observerA)
	observable.RegisterObserver(observerB)

	observable.SetData("New data")

	// Strategy Pattern
	context := &Context{}
	context.SetStrategy(&ConcreteStrategyA{})
	context.DoSomething()
	context.SetStrategy(&ConcreteStrategyB{})
	context.DoSomething()

	// Factory Pattern
	factory := &ConcreteFactoryA{}
	component := factory.CreateComponent()
	component.Operation()

}
