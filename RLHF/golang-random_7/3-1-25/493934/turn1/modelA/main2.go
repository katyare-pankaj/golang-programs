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

	// Optimized approach with string concatenation
	formattedString := "User ID: " + fmt.Sprint(user.ID) + ", Name: " + user.Name + ", Email: " + user.Email
	fmt.Println(formattedString)
}
