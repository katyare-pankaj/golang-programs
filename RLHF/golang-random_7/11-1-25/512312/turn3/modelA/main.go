package main

import (
	"fmt"
)

func main() {
	// Original slice of integers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	oddNumbers := []int{}

	// Use range loop to filter odd numbers
	for _, num := range numbers {
		if num%2 != 0 {
			oddNumbers = append(oddNumbers, num)
		}
	}

	fmt.Println("Original slice:", numbers)
	fmt.Println("Odd numbers:", oddNumbers)
}
