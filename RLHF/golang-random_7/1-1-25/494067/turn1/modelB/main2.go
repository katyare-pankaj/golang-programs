package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Define a slice of Person structs
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 22},
	}

	// Calculate the total age
	totalAge := 0
	for _, person := range people {
		totalAge += person.Age
	}

	// Print the total age
	fmt.Println(totalAge) // Output: 77
}
