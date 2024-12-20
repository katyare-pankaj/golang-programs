package main

import (
	"fmt"
)

func main() {
	// Define a map of maps
	mapOfMaps := map[string]map[string]int{
		"apple": {
			"quantity": 10,
			"price":    2,
		},
		"banana": {
			"quantity": 5,
			"price":    1,
		},
	}

	// Iterate over the outer map
	for outerKey, innerMap := range mapOfMaps {
		fmt.Printf("Outer key: %s\n", outerKey)

		// Iterate over the inner map
		for innerKey, value := range innerMap {
			fmt.Printf("  Inner key: %s, Value: %d\n", innerKey, value)
		}

		fmt.Println() // Add a newline for better readability
	}
}
