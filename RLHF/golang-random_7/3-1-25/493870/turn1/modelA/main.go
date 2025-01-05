package main

import (
	"fmt"
)

func main() {
	name := "Alice"
	age := 30
	result := fmt.Sprintf("Hello, my name is %s and I am %d years old.", name, age)
	fmt.Println(result) // Output: Hello, my name is Alice and I am 30 years old.
}
