package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

// User represents a user object with required fields
type User struct {
	ID    int    `json:"id" required:"true"`
	Name  string `json:"name" required:"true"`
	Email string `json:"email"`
}

// Validate checks if the required fields are present in the struct
func (u *User) Validate() error {
	if u.ID == 0 {
		return fmt.Errorf("field 'id' is required")
	}
	if u.Name == "" {
		return fmt.Errorf("field 'name' is required")
	}
	return nil
}

// ProcessUser takes a JSON string, validates and processes it into a User struct
func ProcessUser(jsonString string) (*User, error) {
	var user User
	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return &user, nil
}

// TestProcessUser tests the ProcessUser function
func TestProcessUser(t *testing.T) {
	// Test normal case
	jsonString := `{"id":1,"name":"John Doe","email":"john@example.com"}`
	user, err := ProcessUser(jsonString)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expectedUser := User{ID: 1, Name: "John Doe", Email: "john@example.com"}
	if *user != expectedUser {
		t.Errorf("Unexpected user: %v, expected %v", user, expectedUser)
	}

	// Test missing ID
	jsonStringMissingID := `{"name":"Jane Doe","email":"jane@example.com"}`
	_, err = ProcessUser(jsonStringMissingID)
	if err == nil {
		t.Fatalf("Expected error, got none")
	}
	if err.Error() != "field 'id' is required" {
		t.Errorf("Unexpected error message: %v, expected 'field 'id' is required'", err.Error())
	}

	// Test missing Name
	jsonStringMissingName := `{"id":2,"email":"jane@example.com"}`
	_, err = ProcessUser(jsonStringMissingName)
	if err == nil {
		t.Fatalf("Expected error, got none")
	}
	if err.Error() != "field 'name' is required" {
		t.Errorf("Unexpected error message: %v, expected 'field 'name' is required'", err.Error())
	}

	// Test invalid JSON
	jsonStringInvalid := `{"id":1,"name":"John Doe,"email":"john@example.com"}`
	_, err = ProcessUser(jsonStringInvalid)
	if err == nil {
		t.Fatalf("Expected error, got none")
	}
	if !strings.Contains(err.Error(), "invalid character") {
		t.Errorf("Unexpected error message: %v, expected error containing 'invalid character'", err.Error())
	}
}
