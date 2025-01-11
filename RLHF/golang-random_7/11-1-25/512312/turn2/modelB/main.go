package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	evenNums := []int{}

	// Using range loop to filter even numbers
	for _, num := range nums {
		if num%2 == 0 {
			evenNums = append(evenNums, num)
		}
	}

	fmt.Println("Original slice:", nums)
	fmt.Println("Even numbers: ", evenNums)
}
