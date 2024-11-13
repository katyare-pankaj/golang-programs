package main

import "fmt"

func sumEvenNumbersEfficient(nums []int) int {
	sum := 0
	for _, num := range nums {
		if num%2 == 0 {
			sum += num
		}
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := sumEvenNumbersEfficient(nums)
	fmt.Println("Sum of even numbers:", result)
}
