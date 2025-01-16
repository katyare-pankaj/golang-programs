package main

import (
	"fmt"
)

type Endpoint struct {
	Method         string
	Path           string
	Parameters     string
	SampleResponse string
}

func generateEndpointDoc(apiVersion string, endpoint Endpoint) string {
	doc := fmt.Sprintf("### %s %s\n", endpoint.Method, endpoint.Path)
	doc += fmt.Sprintf("**Parameters:**\n  %s\n\n", endpoint.Parameters)
	doc += fmt.Sprintf("**Sample Response:**\n  %s\n\n", endpoint.SampleResponse)
	doc += fmt.Sprintf("```json\n%s\n```", endpoint.SampleResponse)
	return doc
}

func main() {
	apiVersion := "v1"
	fmt.Printf("# API Documentation: Version %s\n", apiVersion)
	fmt.Println("---")

	// Sample endpoints data
	endpoints := []Endpoint{
		{
			Method:         "GET",
			Path:           "/users",
			Parameters:     "user_id: integer (optional), limit: integer (default=10), page: integer (default=1)",
			SampleResponse: `{"users": [{"user_id": 1, "name": "Alice", "email": "alice@example.com"}, {"user_id": 2, "name": "Bob", "email": "bob@example.com"}]}`,
		},
		{
			Method:         "POST",
			Path:           "/users",
			Parameters:     "user_name: string, email: string, password: string",
			SampleResponse: `{"status": "success", "user_id": 3}`,
		},
		{
			Method:         "GET",
			Path:           "/users/{user_id}",
			Parameters:     "user_id: integer (required)",
			SampleResponse: `{"user": {"user_id": 1, "name": "Alice", "email": "alice@example.com"}}`,
		},
		{
			Method:         "PUT",
			Path:           "/users/{user_id}",
			Parameters:     "user_id: integer (required), user_data: {name: string, email: string}",
			SampleResponse: `{"status": "success", "message": "User updated successfully"}`,
		},
		{
			Method:         "DELETE",
			Path:           "/users/{user_id}",
			Parameters:     "user_id: integer (required)",
			SampleResponse: `{"status": "success", "message": "User deleted successfully"}`,
		},
	}

	for _, endpoint := range endpoints {
		fmt.Println(generateEndpointDoc(apiVersion, endpoint))
	}
}
