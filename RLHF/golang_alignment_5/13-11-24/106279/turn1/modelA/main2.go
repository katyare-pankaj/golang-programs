package main

import "fmt"

func sumArrayEfficient(array []int) int {
	var total int
	for i := 0; i < len(array); i++ {
		total += array[i]
	}
	return total
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := sumArrayEfficient(arr)
	fmt.Println("Sum:", result)
}
