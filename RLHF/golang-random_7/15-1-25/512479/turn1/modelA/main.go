package main

import (
	"fmt"
	"strings"
)

func main() {
	// Using fmt.Printf
	fmt.Printf("Hello, my name is %s.\n", "Alice")
	fmt.Printf("I am %d years old and I live in %s.\n", 30, "Wonderland")

	// Using fmt.Sprintf
	// Formatting a string and printing it directly
	formattedString := fmt.Sprintf("Hello, my name is %s and I am %d years old.\n", "Bob", 25)
	fmt.Print(formattedString)

	// Using formatted string in other operations
	upperFormattedString := strings.ToUpper(formattedString)
	fmt.Print("Converted to uppercase:", upperFormattedString)

	// Retrieving individual components
	name, ageString := strings.SplitN(upperFormattedString, "AND", 2)
	age, _ := fmt.Sscanf(ageString, " I AM %d YEARS OLD.", &age)

	fmt.Printf("Name: %s, Age: %d\n", strings.TrimSpace(name), age)
}
