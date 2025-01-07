package main

import (
	"fmt"
)

func squareFirstHalf(slice []int) {
	for i := 0; i < len(slice)/2; i++ {
		slice[i] *= slice[i]
	}
}

func doubleSecondHalf(slice []int) {
	for i := len(slice) / 2; i < len(slice); i++ {
		slice[i] *= 2
	}
}

func reverseSlice(slice []int) {
	n := len(slice)
	for i := 0; i < n/2; i++ {
		slice[i], slice[n-1-i] = slice[n-1-i], slice[i]
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Calls the functions to manipulate the slice
	squareFirstHalf(numbers)
	doubleSecondHalf(numbers)
	reverseSlice(numbers)

	// Outputs the final result
	fmt.Println("Modified numbers:", numbers)
}
