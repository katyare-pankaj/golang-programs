package main

import (
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  int
}

func processPerson(person Person) error {
	// Simulate an error for the age 25
	if person.Age == 25 {
		return fmt.Errorf("Age 25 is not allowed: %s", person.Name)
	}

	// Process the person's data (e.g., validation, calculation, etc.)
	fmt.Printf("Processing %s: Age %d\n", person.Name, person.Age)
	return nil
}

func main() {
	dataset := []Person{
		{Name: "Alice", Age: 20},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 30},
		{Name: "David", Age: 22},
		{Name: "Emma", Age: 27},
	}

	for _, person := range dataset {
		if err := processPerson(person); err != nil {
			// Log the error and continue with the next iteration
			log.Printf("Error processing person: %v\n", err)
			continue
		}
	}

	fmt.Println("Processing completed.")
}
