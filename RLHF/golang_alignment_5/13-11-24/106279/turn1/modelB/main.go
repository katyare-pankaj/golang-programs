package main

import "fmt"

func calculateSumReadable(array []int) int {
	totalSum := 0
	for _, value := range array {
		totalSum += value
	}
	return totalSum
}

func main() {
	exampleArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	calculatedSum := calculateSumReadable(exampleArray)
	fmt.Println("Readable Sum:", calculatedSum)
}
