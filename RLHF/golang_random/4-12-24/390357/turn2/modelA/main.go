package main

import (
	"fmt"
)

func sum(nums []int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum:", sum(numbers))
	fmt.Println("Contains 3:", contains(numbers, 3))
}
