package main

import (
	"fmt"
)

func formatStringWithConcatenation(name string, age int, occupation string) string {
	return "Hello, my name is " + name + ", I am " + fmt.Sprint(age) + " years old, and I am a " + occupation + "."
}

func main() {
	name := "Alice"
	age := 30
	occupation := "Engineer"
	formattedString := formatStringWithConcatenation(name, age, occupation)
	fmt.Println(formattedString) // Output: Hello, my name is Alice, I am 30 years old, and I am a Engineer.
}
