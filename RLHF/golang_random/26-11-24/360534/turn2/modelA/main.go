package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define command-line flags for username and password
	username := flag.String("username", "", "Username for login")
	password := flag.String("password", "", "Password for login")

	// Parse command-line flags
	flag.Parse()

	// Predefined credentials for demonstration purposes
	predefinedUsername := "admin"
	predefinedPassword := "secret"

	// Validate credentials
	if *username != predefinedUsername || *password != predefinedPassword {
		fmt.Println("Invalid credentials.")
	} else {
		fmt.Println("Login successful.")
	}
}