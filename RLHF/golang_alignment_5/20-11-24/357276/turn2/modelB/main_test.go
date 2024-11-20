package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// User represents a user object
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ValidateUserJSON validates a JSON string for a User object and returns an error if required fields are missing.
func ValidateUserJSON(jsonString string) error {
	var user User
	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		return err
	}

	if user.ID == 0 {
		return fmt.Errorf("ID is required")
	}

	if user.Name == "" {
		return fmt.Errorf("Name is required")
	}

	if user.Email == "" {
		return fmt.Errorf("Email is required")
	}

	return nil
}

func main() {
	jsonString := `{"id":1,"name":"John Doe","email":"john@example.com"}`
	if err := ValidateUserJSON(jsonString); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("JSON valid")
	}
}

func TestValidateUserJSON(t *testing.T) {
	testCases := []struct {
		input     string
		wantError string
	}{
		{
			input:     `{"id":1,"name":"John Doe","email":"john@example.com"}`,
			wantError: "",
		},
		{
			input:     `{"name":"Jane Doe","email":"jane@example.com"}`,
			wantError: "ID is required",
		},
		{
			input:     `{"id":2,"email":"jane@example.com"}`,
			wantError: "Name is required",
		},
		{
			input:     `{"id":3,"name":"Jane Doe"}`,
			wantError: "Email is required",
		},
		{
			input:     `{"id":1,"name":"John Doe,"email":"john@example.com"}`,
			wantError: "invalid character ',' after object key:value pair",
		},
	}

	for _, tc := range testCases {
		err := ValidateUserJSON(tc.input)
		if err != nil {
			if tc.wantError == "" {
				t.Errorf("Unexpected error for %s: %v", tc.input, err)
			} else if err.Error() != tc.wantError {
				t.Errorf("Unexpected error message for %s. Got '%v', wanted '%v'", tc.input, err.Error(), tc.wantError)
			}
		} else if tc.wantError != "" {
			t.Errorf("Expected error for %s, but got nil", tc.input)
		}
	}
}
