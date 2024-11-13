package main

import (
	"fmt"
)

func filterEvenNumbersReadable(nums []int) []int {
	evenNumbers := make([]int, 0)
	for _, num := range nums {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	return evenNumbers
}

func sumArrayReadable(array []int) int {
	total := 0
	for _, value := range array {
		total += value
	}
	return total
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := filterEvenNumbersReadable(nums)
	result := sumArrayReadable(evenNumbers)
	fmt.Println("Sum of even numbers:", result)
}
