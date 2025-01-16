package main

import (
	"fmt"
)

type Endpoint struct {
	Method     string
	Path       string
	Parameters string
	Response   string
}

func generateEndpointDoc(apiVersion string, endpoint Endpoint) string {
	doc := fmt.Sprintf("### %s %s\n", endpoint.Method, endpoint.Path)
	doc += fmt.Sprintf("**Parameters:**\n  %s\n\n", endpoint.Parameters)
	doc += fmt.Sprintf("**Response:**\n  %s\n\n", endpoint.Response)
	doc += fmt.Sprintf("```json\n%s\n```", endpoint.Response)
	return doc
}

func main() {
	apiVersion := "v1"
	fmt.Printf("# API Documentation: Version %s\n", apiVersion)
	fmt.Println("---")

	// Sample endpoints data
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
			Parameters: "user_name: string, email: string",
			Response:   `{"status": "success", "user_id": 3}`,
		},
		// Add more endpoints as needed
	}

	for _, endpoint := range endpoints {
		fmt.Println(generateEndpointDoc(apiVersion, endpoint))
	}
}
