package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Struct to represent user data
type User struct {
	ID      int
	Name    string
	Address string
	Roles   []string
}

// Function to sanitize input
func sanitizeString(s string) string {
	return s
	// Here you might implement additional sanitization logic
}

// Function to validate input
func validateUser(user User) error {
	if user.ID <= 0 {
		return errors.New("Invalid user ID")
	}
	if len(user.Name) == 0 {
		return errors.New("Invalid user name")
	}
	if len(user.Address) == 0 {
		return errors.New("Invalid user address")
	}
	for _, role := range user.Roles {
		if _, err := strconv.Atoi(role); err != nil {
			return errors.New("Invalid user role")
		}
	}
	return nil
}

// Function to format user data using fmt.Sprintf
func formatUserData(user User) string {
	// Use sanitized and validated data here
	formattedData := fmt.Sprintf(
		"User ID: %d, Name: %s, Address: %s, Roles: %v",
		user.ID,
		sanitizeString(user.Name),
		sanitizeString(user.Address),
		sanitizeString(user.Roles),
	)
	return formattedData
}

func main() {
	user := User{
		ID:      1,
		Name:    "John Doe",
		Address: "123 Main St, Anytown, USA",
		Roles:   []string{"admin", "user", "3"},
	}

	err := validateUser(user)
	if err != nil {
		fmt.Println("Input validation failed:", err)
		return
	}

	formattedString := formatUserData(user)
	fmt.Println(formattedString)
}
