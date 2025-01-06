package main

import "fmt"

// Custom error type
type AppError struct {
	error   // Embed the built-in error type
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func Add(a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, &AppError{Code: 1001, Message: "Division by zero not allowed"}
	}
	return a + b, nil
}

func Divide(a, b int) (float64, error) {
	if a == 0 && b == 0 {
		return 0, &AppError{Code: 1001, Message: "Division by zero not allowed"}
	}
	return float64(a) / float64(b), nil
}

func main() {
	// Attempt to add
	sum, errAdd := Add(10, 0)
	if errAdd != nil {
		if appErr, ok := errAdd.(*AppError); ok {
			fmt.Printf("Addition error: %d, %s\n", appErr.Code, appErr.Message)
		} else {
			fmt.Printf("Unexpected error: %v\n", errAdd)
		}
		return
	}
	fmt.Printf("Sum: %d\n", sum)

	// Attempt to divide
	result, errDivide := Divide(0, 10)
	if errDivide != nil {
		if appErr, ok := errDivide.(*AppError); ok {
			fmt.Printf("Division error: %d, %s\n", appErr.Code, appErr.Message)
		} else {
			fmt.Printf("Unexpected error: %v\n", errDivide)
		}
		return
	}
	fmt.Printf("Result: %v\n", result)
}
