package main

import (
	"fmt"
	"math"
)

type SliceOperation interface {
	Apply([]int) float64
	ProcessElement(int, float64) float64
	InitialResult() float64
}

type BaseOperation struct {
}

func (b *BaseOperation) Apply(slice []int) float64 {
	result := b.InitialResult()
	for _, element := range slice {
		result = b.ProcessElement(element, result)
	}
	return result
}

type SumOperation struct {
	BaseOperation
}

func (s *SumOperation) InitialResult() float64 {
	return 0
}

func (s *SumOperation) ProcessElement(element int, result float64) float64 {
	return result + float64(element)
}

type MaxOperation struct {
	BaseOperation
}

func (m *MaxOperation) InitialResult() float64 {
	return float64(math.MinInt64)
}

func (m *MaxOperation) ProcessElement(element int, result float64) float64 {
	if float64(element) > result {
		return float64(element)
	}
	return result
}

type MeanOperation struct {
	SumOperation
}

func (m *MeanOperation) Apply(slice []int) float64 {
	sum := m.SumOperation.Apply(slice)
	return sum / float64(len(slice))
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	sumOp := &SumOperation{}
	fmt.Println("Sum:", sumOp.Apply(slice)) // Output: Sum: 15

	maxOp := &MaxOperation{}
	fmt.Println("Maximum:", maxOp.Apply(slice)) // Output: Maximum: 5

	meanOp := &MeanOperation{}
	fmt.Println("Mean:", meanOp.Apply(slice)) // Output: Mean: 3
}
