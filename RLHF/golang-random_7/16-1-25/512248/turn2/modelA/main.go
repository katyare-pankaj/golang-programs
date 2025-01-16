package main

import "fmt"

// Endpoint represents the details of an API endpoint
type Endpoint struct {
	Method     string
	Path       string
	Parameters string
	Response   string
}

func main() {
	// Define a list of endpoints
	endpoints := []Endpoint{
		{
			Method:     "GET",
			Path:       "/users",
			Parameters: "user_id: integer (optional)",
			Response:   "user_list: [user object]",
		},
		{
			Method:     "GET",
			Path:       "/users/{user_id}",
			Parameters: "",
			Response:   "user_object: {user details}",
		},
		{
			Method:     "POST",
			Path:       "/users/{user_id}",
			Parameters: "user_data: {user object}",
			Response:   "updated_user_object: {updated user details}",
		},
		{
			Method:     "DELETE",
			Path:       "/users/{user_id}",
			Parameters: "",
			Response:   "message: User deleted successfully",
		},
	}

	// Create a title for the API documentation
	apiVersion := "v1"
	fmt.Printf("# API Documentation: Version %s\n", apiVersion)
	fmt.Println("---")

	// Iterate over each endpoint and print its details
	for _, endpoint := range endpoints {
		endpointString := fmt.Sprintf("### %s %s\n", endpoint.Method, endpoint.Path)
		endpointString += fmt.Sprintf("**Parameters:**\n  %s\n\n", endpoint.Parameters)
		endpointString += fmt.Sprintf("**Response:**\n  %s\n\n", endpoint.Response)
		fmt.Println(endpointString)
	}
}
