package main

import (
	"fmt"
	"go-programs/RLHF/golang_random_9/22-1-25/528320/turn2/modelA/funclib"
)

func main() {
	// Sample data
	numbers := []int{1, 2, 3, 4, 5}

	// Map: Square each number
	squaredNumbers := funclib.Map(numbers, func(x int) int { return x * x })
	fmt.Println("Squared numbers:", squaredNumbers)

	// Filter: Select even numbers
	evenNumbers := funclib.Filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Println("Even numbers:", evenNumbers)

	// Reduce: Sum all numbers
	sumOfNumbers := funclib.Reduce(numbers, func(acc, x int) int { return acc + x }, 0)
	fmt.Println("Sum of numbers:", sumOfNumbers)

	// Reduce: Find the maximum number
	maxNumber := funclib.Reduce(numbers, func(acc, x int) int {
		if x > acc {
			return x
		}
		return acc
	}, numbers[0])
	fmt.Println("Max number:", maxNumber)
}
