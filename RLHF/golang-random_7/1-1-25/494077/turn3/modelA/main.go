package main

import (
	"fmt"
	"strings"
)

// User represents a user with a username and age.
type User struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

// Users is a slice of User.
type Users []User

// SanitizeUsernames sanitizes a list of usernames and formats ages into a readable string format.
func SanitizeUsernames(users Users) []string {
	var sanitizedUsers []string

	for _, user := range users {
		// Sanitize the username by removing special characters
		sanitizedUsername := strings.Map(func(r rune) rune {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
				return r
			}
			return '_' // Replace special characters with an underscore
		}, user.Username)

		// Format the age to display only the last two digits
		formattedAge := fmt.Sprintf("%02d", user.Age%100)

		// Create the formatted string
		formattedString := fmt.Sprintf("Username: %s, Age: %s", sanitizedUsername, formattedAge)
		sanitizedUsers = append(sanitizedUsers, formattedString)
	}

	return sanitizedUsers
}

func main() {
	users := Users{
		{Username: "Alice!@#$", Age: 25},
		{Username: "Bob_123!", Age: 304},
		{Username: "Charlie%^&", Age: 5},
	}

	sanitizedUsers := SanitizeUsernames(users)

	for _, sanitizedUser := range sanitizedUsers {
		fmt.Println(sanitizedUser)
	}
}
