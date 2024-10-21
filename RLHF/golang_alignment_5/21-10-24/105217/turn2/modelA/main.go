package main

import "fmt"

// transformData applies a transformation (doubling each element) to the input data and returns the resulting list.
func transformData(data []int) ([]int, error) {
	// Add error handling for nil input array
	if data == nil {
		return nil, fmt.Errorf("input data cannot be nil")
	}

	transformedData := make([]int, len(data))

	// Use a for-range loop with index to avoid confusion
	for index, value := range data {
		transformedData[index] = value * 2
	}

	return transformedData, nil
}

func main() {
	inputData := []int{1, 2, 3, 4, 5}
	result, err := transformData(inputData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Transformed data:", result)

	// Test for nil input
	_, err = transformData(nil)
	if err != nil {
		fmt.Println("Test for nil input passed:", err)
	}
}
