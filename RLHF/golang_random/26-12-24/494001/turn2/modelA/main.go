package main

import (
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  int
}

func processPerson(person Person) {
	if person.Age < 18 {
		log.Printf("Error: Person %s is underage, age: %d\n", person.Name, person.Age)
	} else {
		fmt.Printf("Processing successful for %s, age: %d\n", person.Name, person.Age)
	}
}

func main() {
	dataset := []Person{
		{Name: "Alice", Age: 17},
		{Name: "Bob", Age: 20},
		{Name: "Charlie", Age: 15},
		{Name: "David", Age: 25},
		{Name: "Emma", Age: 19},
	}

	for _, person := range dataset {
		processPerson(person)
	}

	fmt.Println("Processing completed.")
}
