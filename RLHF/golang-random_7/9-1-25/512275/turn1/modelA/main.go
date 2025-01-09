package main

import "fmt"

func main() {
	// Create a large array of integers
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = i
	}

	// Use a range loop with a pointer to iterate and modify the array
	for i := range nums {
		nums[i] *= nums[i]
	}

	// Print the first few elements to verify the results
	fmt.Println("First few elements after squaring:")
	for _, num := range nums[:10] {
		fmt.Printf("%d ", num)
	}
}
