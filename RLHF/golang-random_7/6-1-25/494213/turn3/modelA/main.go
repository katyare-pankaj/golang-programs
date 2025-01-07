package main

import (
	"errors"
	"fmt"
)

// Function returning multiple values with an error
func readData() (data string, size int, err error) {
	// Simulated error
	if true {
		return "", 0, errors.New("Error reading data")
	}
	// Simulated valid result
	return "Sample data", 1024, nil
}

func main() {
	// Calling the function
	data, size, err := readData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Data: %s, Size: %d\n", data, size)
}
