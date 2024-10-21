package main

import "fmt"

func main() {
	// This line has a typo (should be 'printf')
	fmt.Print("Hello, World!")

	// This variable is unused
	unusedVar := 42

	// This function call is missing
	myFunction(unusedVar)
}

func myFunction(n int) {
	fmt.Println(n)
}
