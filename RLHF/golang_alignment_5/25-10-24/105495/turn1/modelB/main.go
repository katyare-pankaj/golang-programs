//go:build go1.18
// +build go1.18

package main

import (
	"fmt"
	"runtime"
)

// Vpsubd performs packed subtraction of 32-bit integers in 'x' and 'y' and stores the results in 'dst'.
//func Vpsubd(dst *[8]int32, x *[8]int32, y *[8]int32)

func sumArrayBranchAware(arr []int32) int32 {
	total := int32(0)
	length := len(arr)

	// branch-aware optimization for SIMD instruction
	if length >= 8 {
		// Loop iterating with a step of 8, which corresponds to the size of the SIMD register.
		for i := 0; i < length-7; i += 8 {
			// Load 8 int32 values from memory to a SIMD register.
			x := arr[i : i+8]

			// Use SIMD instruction to calculate the sum of 8 integers simultaneously.
			total += vpsubd(&x, &[8]int32{}, &[8]int32{})
		}
	}

	// Add the remaining elements that are not processed by SIMD
	for i := length &^ 7; i < length; i++ {
		total += arr[i]
	}

	return total
}

func main() {
	// Ensure that the code is compiled for the Intel AVX-512 instructionset.
	runtime.CPUProfile()

	// Create a large array to test the performance of the function.
	arrSize := 1000000
	arr := make([]int32, arrSize)

	// Initialize the array with some data.
	for i := range arr {
		arr[i] = int32(i)
	}

	sum := sumArrayBranchAware(arr)
	fmt.Println("Sum:", sum)
}
