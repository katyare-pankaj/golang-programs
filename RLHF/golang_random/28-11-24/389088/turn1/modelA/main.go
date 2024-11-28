package main

import (
	"fmt"
)

// Define custom error types for better handling and structuring
type UserExistsError struct {
	message string
	userId  int
}

func (err *UserExistsError) Error() string {
	return err.message
}

type AdminUserError struct {
	message string
	userId  int
}

func (err *AdminUserError) Error() string {
	return err.message
}

// CreateUser function with error handling for administrators and end-users
func CreateUser(userId int) error {
	// Simulate checking if a user already exists
	if userId == 1 {
		return &UserExistsError{message: "User already exists.", userId: userId}
	}

	// Simulate a random error that only administrators see
	if userId == 2 {
		return &AdminUserError{message: "Internal server error.", userId: userId}
	}

	return nil
}

// HandleUserError function to tailor error messages based on user role
func HandleUserError(err error, isAdmin bool) string {
	if err == nil {
		return "User created successfully."
	}

	switch err.(type) {
	case *UserExistsError:
		userError := err.(*UserExistsError)
		if isAdmin {
			return fmt.Sprintf("User with ID %d already exists.", userError.userId)
		} else {
			return "The user you're trying to create already exists."
		}
	case *AdminUserError:
		adminError := err.(*AdminUserError)
		if isAdmin {
			return fmt.Sprintf("Internal server error when creating user with ID %d.", adminError.userId)
		} else {
			return "There was an error creating your account. Please try again later."
		}
	default:
		if isAdmin {
			return fmt.Sprintf("Unexpected error: %s", err.Error())
		} else {
			return "Something went wrong. Please contact support for assistance."
		}
	}
}

func main() {
	// End-user creating a user
	fmt.Println(HandleUserError(CreateUser(1), false)) // User with ID 1 already exists.
	fmt.Println(HandleUserError(CreateUser(3), false)) // User created successfully.

	// Admin creating a user
	fmt.Println(HandleUserError(CreateUser(1), true)) // User with ID 1 already exists.
	fmt.Println(HandleUserError(CreateUser(2), true)) // Internal server error when creating user with ID 2.
}
