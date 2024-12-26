package main

import (
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{Name: "Alice", Age: 17},
		{Name: "Bob", Age: 18}, // Age 18 will now be treated as valid
		{Name: "Charlie", Age: 19},
		{Name: "David", Age: 14},
		{Name: "Emma", Age: 16},
		{Name: "Frank", Age: 25},
	}

	for _, person := range people {
		if person.Age < 18 {
			// Log error and skip processing for individuals under 18
			log.Printf("Error: %s is underage.\n", person.Name)
			continue
		}

		// Valid person aged 18 or above, print success message
		fmt.Printf("Success: %s is %d years old.\n", person.Name, person.Age)
	}
}
