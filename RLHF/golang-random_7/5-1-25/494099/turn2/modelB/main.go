package main

import (
	"fmt"
)

func main() {
	name := "Alice"
	age := 25
	salary := 50000.00
	isEmployee := true

	formattedString := fmt.Sprintf("Hello, my name is %s. I am %d years old and earn %.2f dollars. Am I an employee? %t", name, age, salary, isEmployee)

	fmt.Println(formattedString)
}
