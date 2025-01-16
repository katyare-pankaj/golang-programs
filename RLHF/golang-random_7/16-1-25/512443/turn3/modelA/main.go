package main

import (
	"fmt"
)

// Define a struct to hold the details of an API endpoint
type Endpoint struct {
	Method     string
	Path       string
	Parameters string
	Response   string
}

// Function to generate markdown documentation for a given endpoint
func generateEndpointDoc(endpoint Endpoint) string {
	doc := fmt.Sprintf("### %s %s\n", endpoint.Method, endpoint.Path)
	doc += fmt.Sprintf("**Parameters:**\n  %s\n\n", endpoint.Parameters)
	doc += fmt.Sprintf("**Response:**\n  %s\n\n", endpoint.Response)
	doc += fmt.Sprintf("```json\n%s\n```\n\n", endpoint.Response)
	return doc
}

func main() {
	// List of endpoints
	endpoints := []Endpoint{
		{
			Method:     "GET",
			Path:       "/users",
			Parameters: "user_id: integer (optional), limit: integer (default=10), page: integer (default=1)",
			Response:   `{"users": [{"user_id": 1, "name": "Alice"}, {"user_id": 2, "name": "Bob"}]}`,
		},
		{
			Method:     "POST",
			Path:       "/users",
			Parameters: "user_name: string, email: string, password: string",
			Response:   `{"status": "success", "user_id": 3}`,
		},
		{
			Method:     "GET",
			Path:       "/users/{user_id}",
			Parameters: "",
			Response:   `{"user": {"user_id": 1, "name": "Alice", "email": "alice@example.com"}}`,
		},
		{
			Method:     "PUT",
			Path:       "/users/{user_id}",
			Parameters: "user_name: string, email: string, password: string",
			Response:   `{"status": "updated", "user_id": 1}`,
		},
		{
			Method:     "DELETE",
			Path:       "/users/{user_id}",
			Parameters: "",
			Response:   `{"status": "deleted"}`,
		},
	}

	// Title of the API documentation
	fmt.Printf("# API Documentation\n")
	fmt.Println("---")

	// Generate and print documentation for each endpoint
	for _, endpoint := range endpoints {
		fmt.Println(generateEndpointDoc(endpoint))
	}
}
