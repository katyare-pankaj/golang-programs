package main

import (
	"fmt"
)

// StringProcessor defines the interface for callbacks that convert ints to strings.
type StringProcessor func(int) string

// ProcessData takes an array of integers and a callback of type StringProcessor.
// It applies the callback to each integer and returns the results as a slice of strings.
func ProcessData(numbers []int, processor StringProcessor) []string {
	var results []string
	for _, num := range numbers {
		results = append(results, processor(num))
	}
	return results
}

// ToUpperCase converts an integer to a string representation of its value followed by "UPPER".
func ToUpperCase(num int) string {
	return fmt.Sprintf("%dUPPER", num)
}

// ToLowerCase converts an integer to a string representation of its value followed by "LOWER".
func ToLowerCase(num int) string {
	return fmt.Sprintf("%dLOWER", num)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Using ToUpperCase as the callback
	upperCasedResults := ProcessData(numbers, ToUpperCase)
	fmt.Println("Results with ToUpperCase:")
	for _, result := range upperCasedResults {
		fmt.Println(result)
	}

	// Using ToLowerCase as the callback
	lowerCasedResults := ProcessData(numbers, ToLowerCase)
	fmt.Println("\nResults with ToLowerCase:")
	for _, result := range lowerCasedResults {
		fmt.Println(result)
	}
}
