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

type SumOperation struct{}

func (s *SumOperation) InitialResult() float64 {
	return 0
}

func (s *SumOperation) ProcessElement(element int, result float64) float64 {
	return result + float64(element)
}

func (s *SumOperation) Apply(slice []int) float64 {
	result := s.InitialResult()
	for _, element := range slice {
		result = s.ProcessElement(element, result)
	}
	return result
}

type MaxOperation struct{}

func (m *MaxOperation) InitialResult() float64 {
	return float64(math.MinInt64)
}

func (m *MaxOperation) ProcessElement(element int, result float64) float64 {
	if float64(element) > result {
		return float64(element)
	}
	return result
}

func (m *MaxOperation) Apply(slice []int) float64 {
	result := m.InitialResult()
	for _, element := range slice {
		result = m.ProcessElement(element, result)
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
