package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	oddNums := []int{}

	for _, num := range nums {
		if num%2 != 0 {
			oddNums = append(oddNums, num)
		}
	}

	fmt.Println("Original slice: ", nums)
	fmt.Println("Odd numbers: ", oddNums)
}
