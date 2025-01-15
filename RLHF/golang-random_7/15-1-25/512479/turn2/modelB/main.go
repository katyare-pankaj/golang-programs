package main

import (
	"fmt"
)

func main() {
	// Format a string using fmt.Sprintf and store it in a variable
	name := "Alice"
	age := 25
	formattedString := fmt.Sprintf("Hello, my name is %s and I am %d years old.", name, age)

	// Modify the formatted string and print the modified result
	modifiedString := fmt.Sprintf("%s I am a software engineer.", formattedString)
	fmt.Println(modifiedString)

	// Use fmt.Printf to directly print a formatted message without storing it
	fmt.Printf("Today is %s, and the weather is %s.\n", "Wednesday", "sunny")
}
