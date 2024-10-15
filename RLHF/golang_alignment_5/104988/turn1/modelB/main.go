package modelB

import "fmt"

// Example function with multiple paths
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}

	// Add a check for negative divisor
	if b < 0 {
		return 0, fmt.Errorf("division by a negative number is not allowed")
	}

	return a / b, nil
}
