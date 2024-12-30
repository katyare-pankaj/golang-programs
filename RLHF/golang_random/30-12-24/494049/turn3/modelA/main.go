package main

import (
	"fmt"
)

// SliceOperation defines the interface for slice operations.
type SliceOperation interface {
	Apply([]int) int
	InitialValue() int
	ProcessElement(int, int) int
	ResultFrom(int, int) int
}

// BaseOperation implements the basic template method for slice operations.
type BaseOperation struct {
}

// Apply is the template method that processes the slice using the custom element processing step.
func (b *BaseOperation) Apply(slice []int) int {
	result := b.InitialValue()
	for _, element := range slice {
		result = b.ProcessElement(result, element)
	}
	return result
}

// SumOperation overrides the BaseOperation to implement the sum operation.
type SumOperation struct {
	BaseOperation
}

// InitialValue returns the initial value for sum, which is 0.
func (s *SumOperation) InitialValue() int {
	return 0
}

// ProcessElement overrides the default processing to add the current element to the result.
func (s *SumOperation) ProcessElement(result, element int) int {
	return result + element
}

// ResultFrom overrides the default result extraction to return the result as is.
func (s *SumOperation) ResultFrom(current, next int) int {
	return current
}

// MaxOperation overrides the BaseOperation to implement the maximum value operation.
type MaxOperation struct {
	BaseOperation
}

// InitialValue returns the initial value for max, which is the first element.
func (m *MaxOperation) InitialValue() int {
	if len(slice) == 0 {
		panic("cannot find max of an empty slice")
	}
	return slice[0]
}

// ProcessElement overrides the default processing to update the result if the current element is greater.
func (m *MaxOperation) ProcessElement(result, element int) int {
	return m.ResultFrom(result, element)
}

// ResultFrom returns the greater of the current and next values.
func (m *MaxOperation) ResultFrom(current, next int) int {
	if next > current {
		return next
	}
	return current
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	sumOp := &SumOperation{}
	fmt.Println("Sum:", sumOp.Apply(slice)) // Output: Sum: 15

	maxOp := &MaxOperation{}
	fmt.Println("Max:", maxOp.Apply(slice)) // Output: Max: 5
}
