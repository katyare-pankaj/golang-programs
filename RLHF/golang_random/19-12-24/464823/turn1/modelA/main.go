package main

import (
	"errors"
	"fmt"
)

// validateStringSlice checks that each string in the slice is non-empty and less than 100 characters long.
func validateStringSlice(input []string) error {
	for index, value := range input {
		if len(value) == 0 {
			return errors.Errorf("empty string at index %d", index)
		}
		if len(value) >= 100 {
			return errors.Errorf("string at index %d exceeds maximum length of 99 characters", index)
		}
		// You can add additional validation here (e.g., for special characters)
	}
	return nil
}

func main() {
	input := []string{"hello", "world", "", "too_long_string_here_1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"}

	err := validateStringSlice(input)
	if err != nil {
		fmt.Println("Validation failed:", err)
		return
	}

	fmt.Println("Input slice is valid.")
	// You can safely process the input slice here
}
