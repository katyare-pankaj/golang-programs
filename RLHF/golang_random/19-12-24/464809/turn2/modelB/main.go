package main

import (
	"fmt"
	"reflect"
)

func sumElements(input interface{}) (float64, error) {
	total := 0.0
	val := reflect.ValueOf(input)

	// Check if the input is a slice
	if val.Kind() != reflect.Slice {
		return 0, fmt.Errorf("unsupported type: %s", val.Kind())
	}

	// Iterate over elements and sum them
	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)
		// Check if the element type is supported (int or float64)
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

func main() {
	// Example usage
	intSlice := []int{1, 2, 3, 4}
	floatSlice := []float64{1.5, 2.5, 3.0}

	sumInt, _ := sumElements(intSlice)
	sumFloat, _ := sumElements(floatSlice)

	fmt.Println("Sum of intSlice:", sumInt)     // Output: Sum of intSlice: 10
	fmt.Println("Sum of floatSlice:", sumFloat) // Output: Sum of floatSlice: 7
}
