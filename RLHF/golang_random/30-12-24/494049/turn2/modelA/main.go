package main

import (
	"fmt"
)

// SliceOperation defines the interface for slice operations.
type SliceOperation interface {
	Apply([]int) int
	ProcessElement(int, int) int
}

// BaseOperation implements the basic template method for slice operations.
type BaseOperation struct {
}

// Apply is the template method that processes the slice using the custom element processing step.
func (b *BaseOperation) Apply(slice []int) int {
	result := 0
	for _, element := range slice {
		result = b.ProcessElement(result, element)
	}
	return result
}

// ProcessElement is the default element processing step (sum).
func (b *BaseOperation) ProcessElement(result, element int) int {
	return result + element
}

// SumOperation overrides the ProcessElement method to implement the sum operation.
type SumOperation struct {
	BaseOperation
}

// ProductOperation overrides the ProcessElement method to implement the product operation.
type ProductOperation struct {
	BaseOperation
}

// ProcessElement overrides the default sum to calculate the product.
func (p *ProductOperation) ProcessElement(result, element int) int {
	if result == 0 {
		return element
	}
	return result * element
}

// CustomOperation demonstrates how to define a new custom operation.
type CustomOperation struct {
	BaseOperation
}

// ProcessElement overrides the default sum to implement a custom operation (e.g., sum of squares).
func (c *CustomOperation) ProcessElement(result, element int) int {
	return result + element*element
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	sumOp := &SumOperation{}
	fmt.Println("Sum:", sumOp.Apply(slice)) // Output: Sum: 15

	productOp := &ProductOperation{}
	fmt.Println("Product:", productOp.Apply(slice)) // Output: Product: 120

	customOp := &CustomOperation{}
	fmt.Println("Sum of Squares:", customOp.Apply(slice)) // Output: Sum of Squares: 55
}
