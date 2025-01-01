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

	// Create a map to store ages with names as keys
	ages := map[string]int{}
	for _, person := range people {
		ages[person.Name] = person.Age
	}

	// Print the map
	fmt.Println(ages) // Output: map[Alice:25 Bob:30 Charlie:22]
}
