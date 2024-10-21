package main

import "fmt"

func transformData(data []int) []int {
	transformedData := make([]int, len(data))
	for i, value := range data {
		transformedData[i] = value * 2
	}
	return transformedData
}

func main() {
	inputData := []int{1, 2, 3, 4, 5}
	result := transformData(inputData)
	fmt.Println("Transformed data:", result)
}
