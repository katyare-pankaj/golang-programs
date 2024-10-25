// Example usage of compiler flags
//go:build go1.18
// +build go1.18

package main

import "fmt"

//go:noinline
func sumArrayAutoVectorized(arr []int32) int32 {
	total := int32(0)
	for _, v := range arr {
		total += v
	}
	return total
}

func main() {
	// Enable compiler flags for AVX-512 support and optimization
	// The exact flags may vary depending on the Go compiler version and platform
	//runtime.SetCPUProfileRate(1)

	arrSize := 10000000
	arr := make([]int32, arrSize)

	for i := range arr {
		arr[i] = int32(i)
	}

	sum := sumArrayAutoVectorized(arr)
	fmt.Println("Sum:", sum)
}
