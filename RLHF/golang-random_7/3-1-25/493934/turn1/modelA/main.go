package main

import (
	"fmt"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	user := User{ID: 1, Name: "John Doe", Email: "john@example.com"}

	// Original approach with fmt.Sprintf
	formattedString := fmt.Sprintf("User ID: %d, Name: %s, Email: %s", user.ID, user.Name, user.Email)
	fmt.Println(formattedString)
}
