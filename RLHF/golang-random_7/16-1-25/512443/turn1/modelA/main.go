package main

import (
	"fmt"
	"strings"
)

// Endpoint represents an API endpoint with its path and description.
type Endpoint struct {
	Path        string
	Method      string
	Description string
	Version     string
}

// generateDocs generates formatted documentation for a list of endpoints.
func generateDocs(endpoints []Endpoint) string {
	var docs string
	for _, endpoint := range endpoints {
		docs += fmt.Sprintf("\n## API Documentation\n### Version: %s\n### Endpoint: %s\n### Method: %s\n\n%s\n",
			endpoint.Version,
			endpoint.Path,
			endpoint.Method,
			strings.Indent(endpoint.Description, "  "))
	}
	return docs
}

func main() {
	// Define a list of endpoints with their information
	endpoints := []Endpoint{
		{
			Path:        "/v1/users",
			Method:      "GET",
			Description: "Retrieves a list of all users.",
			Version:     "1.0",
		},
		{
			Path:        "/v1/users/{id}",
			Method:      "GET",
			Description: "Retrieves a single user by ID.",
			Version:     "1.0",
		},
		{
			Path:        "/v1/users",
			Method:      "POST",
			Description: "Creates a new user.",
			Version:     "1.0",
		},
	}

	// Generate documentation
	docs := generateDocs(endpoints)

	// Print the generated documentation
	fmt.Println(docs)
}
