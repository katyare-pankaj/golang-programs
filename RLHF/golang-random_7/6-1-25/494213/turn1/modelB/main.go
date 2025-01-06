package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func getUserInfo(userID int) (user *User, err error) {
	// Simulate fetching user information
	if userID != 1 { // Example error condition
		return nil, errors.New("User not found")
	}

	// If no error, return a sample user
	return &User{
		ID:   userID,
		Name: "John Doe",
		Age:  30,
	}, nil
}

func main() {
	user, err := getUserInfo(1)
	if err != nil || user == nil { // Check for sentinel value
		// Handle error
		fmt.Println("Error fetching user information:", err)
		return
	}

	fmt.Println("User:", user.Name)
}
