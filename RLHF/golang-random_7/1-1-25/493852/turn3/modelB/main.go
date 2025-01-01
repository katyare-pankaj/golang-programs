package main

import (
	"fmt"
	"regexp"
)

// User represents a user with sensitive information.
type User struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

// SanitizedUsers is a slice of sanitized User.
type SanitizedUsers []string

// sanitizeUsername removes any special characters from the username.
func sanitizeUsername(username string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	return re.ReplaceAllString(username, "")
}

// sanitizeAge formats the age to display only the last two digits.
func sanitizeAge(age int) string {
	return fmt.Sprintf("%02d", age%100)
}

// FormatUsers safely formats a list of users with sanitized information.
func FormatUsers(users []User) SanitizedUsers {
	sanitizedUsers := make(SanitizedUsers, len(users))
	for i, user := range users {
		sanitizedUsername := sanitizeUsername(user.Username)
		sanitizedAge := sanitizeAge(user.Age)
		sanitizedUsers[i] = fmt.Sprintf("Username: %s, Age: %s", sanitizedUsername, sanitizedAge)
	}
	return sanitizedUsers
}

func main() {
	users := []User{
		{Username: "alice!@#$", Age: 25},
		{Username: "bob123", Age: 30},
		{Username: "charles_321", Age: -10},
		{Username: "daisy-54", Age: 100},
	}

	sanitizedUsers := FormatUsers(users)

	for _, sanitizedUser := range sanitizedUsers {
		fmt.Println(sanitizedUser)
	}
}
