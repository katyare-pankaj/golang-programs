package main

import (
	"fmt"
)

type SliceProcessor interface {
	ProcessSlice([]int) int
	CustomizeStep(int, int) int
}

type BaseSliceProcessor struct {
}

func (b *BaseSliceProcessor) ProcessSlice(slice []int) int {
	total := 0
	for i := 0; i < len(slice); i++ {
		currentElement := slice[i]
		total = b.CustomizeStep(total, currentElement)
	}
	return total
}

func (b *BaseSliceProcessor) CustomizeStep(total, currentElement int) int {
	return total + currentElement
}

type SumProcessor struct {
	BaseSliceProcessor
}

func (s *SumProcessor) CustomizeStep(total, currentElement int) int {
	return total + currentElement
}

type ProductProcessor struct {
	BaseSliceProcessor
}

func (p *ProductProcessor) CustomizeStep(total, currentElement int) int {
	if total == 0 {
		return currentElement
	}
	return total * currentElement
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	sumProcessor := &SumProcessor{}
	fmt.Println("Sum:", sumProcessor.ProcessSlice(slice)) // Output: Sum: 15

	productProcessor := &ProductProcessor{}
	fmt.Println("Product:", productProcessor.ProcessSlice(slice)) // Output: Product: 120
}
