package main

import (
	"fmt"
	"html"
	"log"
	"strings"
)

// User represents a user with personal details
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// sanitize sanitizes a string by escaping HTML special characters
func sanitize(s string) string {
	if s == "" {
		return ""
	}
	return html.EscapeString(s)
}

// formatUsers formats a list of users with sanitized details
func formatUsers(users []User) (string, error) {
	var result strings.Builder
	for _, user := range users {
		// Sanitize name and age before appending to the result
		sanitizedName := sanitize(user.Name)
		sanitizedAge := sanitize(fmt.Sprintf("%d", user.Age)) // Convert int to string

		// Append formatted user data to the result
		result.WriteString(fmt.Sprintf("Name: %s, Age: %s\n", sanitizedName, sanitizedAge))
	}

	if result.Len() == 0 {
		return "", fmt.Errorf("no users to display")
	}

	return result.String(), nil
}

func main() {
	users := []User{
		{Name: "Alice <script>alert(1)</script>", Age: 25},
		{Name: "Bob", Age: 30},
	}

	formattedUsers, err := formatUsers(users)
	if err != nil {
		log.Fatalf("Error formatting users: %v", err)
	}

	fmt.Println(formattedUsers)
}
