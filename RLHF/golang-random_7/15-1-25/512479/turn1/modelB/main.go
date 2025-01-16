package main

import (
	"fmt"
)

func main() {
	num := 42
	str := "Hello, World!"

	// Using fmt.Printf
	fmt.Printf("Using fmt.Printf: The number is %d and the string is %s\n", num, str)

	// Using fmt.Sprintf
	formattedStr := fmt.Sprintf("Using fmt.Sprintf: The number is %d and the string is %s", num, str)
	fmt.Println("Using fmt.Sprintf:", formattedStr)

	// Now we can use the formatted string for further operations
	formattedStr += " This string can be modified and used elsewhere."
	fmt.Println(formattedStr)
}
