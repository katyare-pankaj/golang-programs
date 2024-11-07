package main

import (
	"fmt"
)

func main() {
	// Sample data
	data := []int{1, 2, 3, 4, 5}
	sum := 0

	// In-place computation to minimize memory allocation
	for _, value := range data {
		sum += value
	}

	fmt.Println("Sum:", sum)
}
