package calculator

import "fmt"

// Calculator struct encapsulates the calculator logic
type Calculator struct{}

// Add performs addition
func (c Calculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract performs subtraction
func (c Calculator) Subtract(a, b float64) float64 {
	return a - b
}

func ExampleCalculator() {
	calc := Calculator{}
	result := calc.Add(10.0, 20.0)
	fmt.Println("Result:", result) // Output: Result: 30
}
