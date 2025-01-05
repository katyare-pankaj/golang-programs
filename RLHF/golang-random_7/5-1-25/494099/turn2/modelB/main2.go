package main

import (
	"fmt"
)

func main() {
	name := "Bob"
	age := 30
	salary := 75000.00
	isEmployee := false

	formattedString := "Hello, my name is " + name + ". I am " + fmt.Sprint(age) + " years old and earn " + fmt.Sprintf("%.2f", salary) + " dollars. Am I an employee? " + fmt.Sprint(isEmployee)

	fmt.Println(formattedString)
}
