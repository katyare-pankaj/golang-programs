package main

import (
	"fmt"
)

func main() {
	// Iterating over a slice
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Iterating over a map
	personAges := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 35}
	for name, age := range personAges {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}
}
