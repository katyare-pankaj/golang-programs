package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// User represents the input user data
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// TransformedUser represents the output user data
type TransformedUser struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Domain   string `json:"domain"`
}

func main() {
	// Sample user data in JSON format
	jsonData := `[
		{"id": 1, "username": "john_doe", "email": "john.doe@example.com"},
		{"id": 2, "username": "jane_smith", "email": "jane.smith@example.com"}
	]`

	// Step 2: Unmarshal the JSON data into User struct
	var users []User
	if err := json.Unmarshal([]byte(jsonData), &users); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Step 3: Transform the data
	var transformedUsers []TransformedUser
	for _, user := range users {
		domain := user.Email[strings.LastIndex(user.Email, "@")+1:] // Extract domain from email
		transformedUser := TransformedUser{
			UserID:   user.ID,
			UserName: user.Username,
			Domain:   domain,
		}
		transformedUsers = append(transformedUsers, transformedUser)
	}

	// Step 4: Marshal the transformed data back to JSON
	output, err := json.Marshal(transformedUsers)
	if err != nil {
		log.Fatalf("Error marshalling transformed data: %v", err)
	}

	// Output the transformed data
	fmt.Println(string(output))
}
