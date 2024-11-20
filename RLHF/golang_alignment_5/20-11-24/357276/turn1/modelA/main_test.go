package main

import (
	"encoding/json"
	"testing"
)

// User represents a user object
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UnmarshalUser takes a JSON string and unmarshals it into a User struct
func UnmarshalUser(jsonString string) (*User, error) {
	var user User
	err := json.Unmarshal([]byte(jsonString), &user)
	return &user, err
}

// TestUnmarshalUser tests the UnmarshalUser function
func TestUnmarshalUser(t *testing.T) {
	// Test normal case
	jsonString := `{"id":1,"name":"John Doe","email":"john@example.com"}`
	user, err := UnmarshalUser(jsonString)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expectedUser := User{ID: 1, Name: "John Doe", Email: "john@example.com"}
	if *user != expectedUser {
		t.Errorf("Unexpected user: %v, expected %v", user, expectedUser)
	}

	// Test missing ID
	jsonStringMissingID := `{"name":"Jane Doe","email":"jane@example.com"}`
	user, err = UnmarshalUser(jsonStringMissingID)
	if err == nil {
		t.Fatalf("Expected error, got none")
	}

	// Test missing Name
	jsonStringMissingName := `{"id":2,"email":"jane@example.com"}`
	user, err = UnmarshalUser(jsonStringMissingName)
	if err == nil {
		t.Fatalf("Expected error, got none")
	}

	// Test missing Email
	jsonStringMissingEmail := `{"id":3,"name":"Jane Doe"}`
	user, err = UnmarshalUser(jsonStringMissingEmail)
	if err == nil {
		t.Fatalf("Expected error, got none")
	}

	// Test invalid JSON
	jsonStringInvalid := `{"id":1,"name":"John Doe,"email":"john@example.com"}`
	user, err = UnmarshalUser(jsonStringInvalid)
	if err == nil {
		t.Fatalf("Expected error, got none")
	}
}
