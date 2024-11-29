package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the "ENABLE_FEATURE" environment variable is set
	enableFeature := os.Getenv("ENABLE_FEATURE") == "true"

	if enableFeature {
		fmt.Println("Feature is enabled!")
		doSomethingNew()
	} else {
		fmt.Println("Feature is disabled.")
	}
}

func doSomethingNew() {
	fmt.Println("Doing something new and exciting!")
}
