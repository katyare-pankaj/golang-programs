package main

import (
	"fmt"
)

func main() {
	nestedMap := map[string]map[string]int{
		"employee1": {"age": 30, "salary": 5000},
		"employee2": {"age": 28, "salary": 4500},
		"employee3": {"age": 35, "salary": 6000},
	}

	for outerKey, innerMap := range nestedMap {
		fmt.Printf("\nEmployee Name: %s\n", outerKey)
		for innerKey, value := range innerMap {
			fmt.Printf("  %s: %d\n", innerKey, value)
		}
	}
}
