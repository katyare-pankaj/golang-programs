package main

import (
	"fmt"
	"strings"
)

// User represents a user with sensitive information.
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Users is a slice of User.
type Users []User

// SanitizedUsers is a slice of sanitized User.
type SanitizedUsers []string

// SanitizeName sanitizes a user's name by removing any non-alphabetic characters.
func SanitizeName(name string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			return r
		}
		return '*'
	}, name)
}

// SanitizeAge sanitizes a user's age by displaying only the last two digits.
func SanitizeAge(age int) string {
	if age < 0 {
		return "--"
	}
	return fmt.Sprintf("%02d", age%100)
}

// FormatUsers safely formats a list of users with sanitized information.
func FormatUsers(users Users) SanitizedUsers {
	sanitizedUsers := make(SanitizedUsers, len(users))
	for i, user := range users {
		sanitizedUsers[i] = fmt.Sprintf("Name: %s, Age: %s", SanitizeName(user.Name), SanitizeAge(user.Age))
	}
	return sanitizedUsers
}

func main() {
	users := Users{
		{Name: "Alice@gmail.com", Age: 25},
		{Name: "Bob!123", Age: 30},
		{Name: "Charlie", Age: -10},
	}

	sanitizedUsers := FormatUsers(users)

	for _, sanitizedUser := range sanitizedUsers {
		fmt.Println(sanitizedUser)
	}
}
