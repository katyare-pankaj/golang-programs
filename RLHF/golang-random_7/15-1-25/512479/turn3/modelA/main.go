package main

import "fmt"

func main() {
	// Define the variables to be used in the formatted string
	name := "Bob"
	age := 30
	profession := "Engineer"

	// Format the string using fmt.Sprintf and store it in a variable
	formattedString := fmt.Sprintf("%s is a %d-year-old %s.", name, age, profession)

	// Print the initial formatted string
	fmt.Println(formattedString)

	// Append additional information to the formatted string and print it
	modifiedString := formattedString + " He has been working at TechCorp for 5 years."
	fmt.Println(modifiedString)
}
