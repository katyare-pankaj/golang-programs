package main

import "fmt"

// AbstractionLayer represents a high-level abstraction in game development
type AbstractionLayer struct {
	value int
}

func (a *AbstractionLayer) Process() int {
	// Some complex processing logic
	a.value = a.value * 2
	return a.value
}

func main() {
	layer := &AbstractionLayer{value: 5}
	result := layer.Process()
	fmt.Println("Result:", result) // Output: Result: 20
}
