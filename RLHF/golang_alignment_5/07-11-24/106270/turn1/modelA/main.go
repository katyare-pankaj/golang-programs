// With Abstraction (Potentially Less Performant during Merging)

package main

import "fmt"

type Calculator interface {
	Add(a, b float64) float64
	Multiply(a, b float64) float64
}

type basicCalculator struct{}

func (c *basicCalculator) Add(a, b float64) float64 {
	return a + b
}

func (c *basicCalculator) Multiply(a, b float64) float64 {
	return a * b
}

func NewCalculator() Calculator {
	return &basicCalculator{}
}

func main() {

	calculator := NewCalculator()

	resultAdd := calculator.Add(5, 3)
	fmt.Println("5 + 3 =", resultAdd)

	resultMultiply := calculator.Multiply(5, 3)
	fmt.Println("5 * 3 =", resultMultiply)
}
