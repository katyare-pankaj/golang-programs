package main

import (
	"encoding/json"
	"fmt"
)

// Employee struct generated from the domain model
type Employee struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Skills []Skill `json:"skills"`
}

// Skill struct generated from the domain model
type Skill struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Training struct generated from the domain model
type Training struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Skills      []Skill    `json:"skills"`
	Employees   []Employee `json:"employees"`
}

func main() {
	emp := Employee{
		ID:   "1",
		Name: "Alice",
		Skills: []Skill{{
			ID:   "1",
			Name: "Golang",
		}},
	}

	// Additional code to handle validation, persistence, and scaling
	// ...

	empJSON, err := json.Marshal(emp)
	if err != nil {
		fmt.Println("Error marshalling employee:", err)
		return
	}

	// Print the JSON data
	fmt.Println("Employee JSON:", string(empJSON))

}
