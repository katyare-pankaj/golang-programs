package main

import (
	"fmt"
	"net/http"
)

type apiResponse struct {
	StatusCode int
	Message    string
}

func main() {
	// Create a map of API endpoints and their corresponding responses
	apiEndpoints := map[string]apiResponse{
		"/user/{id}": {
			StatusCode: http.StatusOK,
			Message:    "User data retrieved successfully.",
		},
		"/posts": {
			StatusCode: http.StatusCreated,
			Message:    "Post created successfully.",
		},
		"/comments/{post_id}": {
			StatusCode: http.StatusNotFound,
			Message:    "Comments for the specified post not found.",
		},
	}

	generateAPIDocs(apiEndpoints)
}

func generateAPIDocs(endpoints map[string]apiResponse) {
	fmt.Println("# API Documentation")
	fmt.Println("## Endpoints")

	for endpoint, response := range endpoints {
		fmt.Println("###", endpoint)
		fmt.Println("- **Response Status:**", response.StatusCode)
		fmt.Println("- **Response Message:**", response.Message)
		fmt.Println("- **HTTP Status Code Description:**", http.StatusText(response.StatusCode))
		fmt.Println()
	}
}
