package main

import (
	"fmt"
)

func main() {
	name := "James"
	planet := "Mars"
	time := "22:15"

	// Using fmt.Sprintf to format a string with multiple variables and store it in a variable
	formattedString := fmt.Sprintf("Greetings, %s! It's currently %s on %s.", name, time, planet)

	// Modifying the formatted string by appending additional information
	modifiedString := fmt.Sprintf("%s I hope you are doing well. I just completed a mission to %s.", formattedString, planet)

	// Printing the modified string using fmt.Println
	fmt.Println(modifiedString)
}
