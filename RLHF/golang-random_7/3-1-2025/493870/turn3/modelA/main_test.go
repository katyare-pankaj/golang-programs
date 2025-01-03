package main

import (
	"fmt"
	"testing"
)

func reverseSliceForLoop(input []int) []int {
	result := make([]int, len(input))
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func reverseSliceCopy(input []int) []int {
	result := make([]int, len(input))
	copy(result[:0], input[len(input)-1:])
	return result
}

func BenchmarkReverseSliceForLoop(b *testing.B) {
	input := make([]int, 100000)
	for i := range input {
		input[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseSliceForLoop(input)
	}
}

func BenchmarkReverseSliceCopy(b *testing.B) {
	input := make([]int, 100000)
	for i := range input {
		input[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseSliceCopy(input)
	}
}

func main() {
	input := make([]int, 10)
	for i := range input {
		input[i] = i
	}
	fmt.Println("Reversed slice using for loop:", reverseSliceForLoop(input))
	fmt.Println("Reversed slice using copy:", reverseSliceCopy(input))
}
