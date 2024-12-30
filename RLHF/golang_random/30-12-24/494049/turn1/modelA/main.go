package main

import (
	"fmt"
)

// Define an interface for slice operations
type SliceOperations interface {
	// Template method defining the algorithm
	ProcessSlice()
	// Steps that subclasses can override
	Filter([]int) []int
	Map([]int) []int
	Reduce([]int) int
}

// Implement a base structure for SliceOperations
type BaseSliceOperations struct {
	data []int
}

// Implement the template method
func (b *BaseSliceOperations) ProcessSlice() {
	fmt.Println("Applying filter...")
	filtered := b.Filter(b.data)

	fmt.Println("Applying map...")
	mapped := b.Map(filtered)

	fmt.Println("Applying reduce...")
	result := b.Reduce(mapped)

	fmt.Println("Result:", result)
}

// Default implementations of steps that can be overridden
func (b *BaseSliceOperations) Filter(data []int) []int {
	return data
}

func (b *BaseSliceOperations) Map(data []int) []int {
	return data
}

func (b *BaseSliceOperations) Reduce(data []int) int {
	return 0
}

// Customize the behavior by creating a subclass
type CustomSliceOperations struct {
	BaseSliceOperations
}

// Override specific steps
func (c *CustomSliceOperations) Filter(data []int) []int {
	return data[:len(data)/2]
}

func (c *CustomSliceOperations) Map(data []int) []int {
	result := make([]int, len(data))
	for i, value := range data {
		result[i] = value * value
	}
	return result
}

func (c *CustomSliceOperations) Reduce(data []int) int {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return sum
}

func main() {
	// Create an instance of CustomSliceOperations
	customOps := &CustomSliceOperations{
		BaseSliceOperations{
			data: []int{1, 2, 3, 4, 5},
		},
	}

	// Use the template method
	customOps.ProcessSlice()
}
