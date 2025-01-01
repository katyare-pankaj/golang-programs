package main

import (
	"fmt"
	"strings"
)

// User represents a user with basic information
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Users is a slice of User
type Users []User

// String method for Users type
func (u Users) String() string {
	var sb strings.Builder
	for _, user := range u {
		sb.WriteString(fmt.Sprintf("User Name: %s, User Age: %d\n", user.Name, user.Age))
	}
	return sb.String()
}

func main() {
	users := Users{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
	}

	// Use the custom String() method to safely format the Users
	safeOutput := users.String()
	fmt.Println(safeOutput)

	// Validate user input and sanitize before using it in fmt.Sprintf
	// ...
}
