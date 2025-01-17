package main

import (
	"fmt"
	"net/http"
)

// User struct represents user data
type User struct {
	Name  string
	Age   int
	Email string
	Role  string
}

// handler function processes HTTP requests and constructs dynamic responses using fmt.Sprintf
func handler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	email := r.URL.Query().Get("email")
	role := r.URL.Query().Get("role")

	// Create a User struct with the parsed parameters
	user := User{
		Name:  name,
		Age:   int(parseAge(age)),
		Email: email,
		Role:  role,
	}

	// Construct a dynamic response using fmt.Sprintf
	response := fmt.Sprintf("User Details:\nName: %s\nAge: %d\nEmail: %s\nRole: %s\n",
		user.Name, user.Age, user.Email, user.Role)

	// Send the response
	w.Write([]byte(response))
}

// parseAge helper function to parse age parameter
func parseAge(age string) int {
	if age == "" {
		return 0
	}
	return int(fmt.Sprintf("%d", age))
}

func main() {
	// Handle GET requests to /user
	http.HandleFunc("/user", handler)

	// Start the HTTP server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
