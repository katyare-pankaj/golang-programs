package main

import "fmt"

func calculateSumEfficient(array []int) int {
	return sum(array, 0)
}

func sum(arr []int, start int) int {
	if start == len(arr) {
		return 0
	}
	return arr[start] + sum(arr, start+1)
}

func main() {
	exampleArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	calculatedSum := calculateSumEfficient(exampleArray)
	fmt.Println("Efficient Sum:", calculatedSum)
}
