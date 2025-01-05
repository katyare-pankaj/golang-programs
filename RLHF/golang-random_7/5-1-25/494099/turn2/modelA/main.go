package main

import (
	"fmt"
)

func formatStringWithSprintf(name string, age int, occupation string) string {
	return fmt.Sprintf("Hello, my name is %s, I am %d years old, and I am a %s.", name, age, occupation)
}

func main() {
	name := "Alice"
	age := 30
	occupation := "Engineer"
	formattedString := formatStringWithSprintf(name, age, occupation)
	fmt.Println(formattedString) // Output: Hello, my name is Alice, I am 30 years old, and I am a Engineer.
}
