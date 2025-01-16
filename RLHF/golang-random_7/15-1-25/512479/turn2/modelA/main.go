package main

import (
	"fmt"
)

func main() {
	// Initialize variables
	name := "Alice"
	age := 30

	// Format a string using fmt.Sprintf and store it in a variable
	formattedString := fmt.Sprintf("Hello, my name is %s and I am %d years old.", name, age)

	// Modify the stored formatted string
	formattedString += " How are you today?"

	// Print the modified string
	fmt.Println(formattedString)

	// Use fmt.Printf to directly print a formatted message
	fmt.Printf("Directly using fmt.Printf: Hello, my name is %s and I am %d years old.", name, age)
}
