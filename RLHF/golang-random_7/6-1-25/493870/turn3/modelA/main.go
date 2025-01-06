package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("# HTTP Status Codes and Descriptions")

	// Loop over a range of standard HTTP status codes
	for code := http.StatusContinue; code < http.StatusBadGateway; code++ {
		// Retrieve the status description using http.StatusText
		description := http.StatusText(code)
		if description != "" {
			// Print the status code and its description
			fmt.Printf("- %d: %s\n", code, description)
		}
	}
}
