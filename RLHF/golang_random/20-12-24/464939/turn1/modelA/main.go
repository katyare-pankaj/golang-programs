package main

import (
	"fmt"
)

func main() {
	// Define a nested map
	nestedMap := map[string]map[string]int{
		"Fruits": map[string]int{
			"Apple":  5,
			"Banana": 3,
			"Orange": 4,
		},
		"Vegetables": map[string]int{
			"Carrot": 2,
			"Tomato": 6,
		},
	}

	// Iterate over the first level of the map
	for outerKey, innerMap := range nestedMap {
		fmt.Println("Category:", outerKey)

		// Iterate over the second level of the map
		for innerKey, value := range innerMap {
			fmt.Printf("  %s: %d\n", innerKey, value)
		}
	}
}
