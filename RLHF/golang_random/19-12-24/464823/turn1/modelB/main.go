package main

import (
	"fmt"
	"strconv"
)

func validateIntSliceInput(input []string) []int {
	validSlice := make([]int, 0, len(input)) // Preallocate slice with the same capacity as input

	for _, value := range input {
		// Convert string to int using strconv.Atoi
		if num, err := strconv.Atoi(value); err == nil {
			validSlice = append(validSlice, num)
		} else {
			// Handle invalid input (e.g., non-numeric value)
			fmt.Println("Invalid input:", value)
			return nil
		}
	}
	return validSlice
}

func main() {
	unsafeInput := []string{"1", "2", "three", "4", "5"}
	safeInput := []string{"1", "2", "3", "4", "5"}

	unsafeValidated := validateIntSliceInput(unsafeInput)
	safeValidated := validateIntSliceInput(safeInput)

	if unsafeValidated != nil {
		fmt.Println("Unsafe input validation passed:", unsafeValidated)
	}

	if safeValidated != nil {
		fmt.Println("Safe input validation passed:", safeValidated)
	}
}
