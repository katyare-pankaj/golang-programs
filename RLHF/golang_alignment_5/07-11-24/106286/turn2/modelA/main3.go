package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make([]int, len(data))

	// Reusing result buffer to avoid allocation
	for i, value := range data {
		result[i] = value * value
	}
	fmt.Println("Result:", result)
}
