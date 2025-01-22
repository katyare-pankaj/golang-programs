package main

import (
	"fmt"
	"math"
)

func squareSliceElements(input []int) []int {
	result := make([]int, len(input))
	for i, v := range input {
		if v < 0 {
			result[i] = int(math.Abs(float64(v))) * int(math.Abs(float64(v)))
		} else {
			result[i] = v * v
		}
	}
	return result
}

func main() {
	testInput := []int{-1, 2, -3, 4, -5, 0, 6}
	fmt.Println("Original Data: ", testInput)
	result := squareSliceElements(testInput)
	fmt.Println("Squared Data: ", result)
}
