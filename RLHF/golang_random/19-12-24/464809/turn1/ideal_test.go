package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Function to sum elements of a slice or array
func sumElements(input interface{}) (float64, error) {
	total := 0.0
	val := reflect.ValueOf(input)

	// Ensure the input is a slice or an array
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		return 0, fmt.Errorf("unsupported type: %s", val.Kind())
	}

	// Iterate over elements using reflection
	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)
		switch elem.Kind() {
		case reflect.Int:
			total += float64(elem.Int())
		case reflect.Float64:
			total += elem.Float()
		default:
			return 0, fmt.Errorf("unsupported element type: %s", elem.Kind())
		}
	}
	return total, nil
}

// Test function using Go's testing package
func TestSumElements(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected float64
	}{
		{"Sum of integers", []int{1, 2, 3, 4}, 10},
		{"Sum of floats", []float64{1.5, 2.5, 3.0}, 7},
		{"Empty slice", []int{}, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := sumElements(tc.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func main() {
	// Example usage of sumElements
	sum, err := sumElements([]interface{}{1, 2, 3.5})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sum:", sum)
	}
}
