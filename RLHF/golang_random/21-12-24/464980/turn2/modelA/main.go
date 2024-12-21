package main

import (
	"fmt"
)

// DataProcessor is a function type that takes an int and returns a string.
type DataProcessor func(int) string

// ProcessData applies a DataProcessor function to each integer in a slice
// and returns the results as a slice of strings.
func ProcessData(numbers []int, processor DataProcessor) []string {
	var results []string
	for _, num := range numbers {
		results = append(results, processor(num))
	}
	return results
}

// AddOneProcessor adds one to the input and converts it to a string.
func AddOneProcessor(num int) string {
	return fmt.Sprintf("%d+1", num)
}

// MultiplyByTwoProcessor multiplies the input by two and converts it to a string.
func MultiplyByTwoProcessor(num int) string {
	return fmt.Sprintf("%d*2", num)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Using AddOneProcessor as the callback
	addOneResults := ProcessData(numbers, AddOneProcessor)
	fmt.Println("Results with AddOneProcessor:")
	for _, result := range addOneResults {
		fmt.Println(result)
	}

	// Using MultiplyByTwoProcessor as the callback
	multiplyTwoResults := ProcessData(numbers, MultiplyByTwoProcessor)
	fmt.Println("\nResults with MultiplyByTwoProcessor:")
	for _, result := range multiplyTwoResults {
		fmt.Println(result)
	}
}
