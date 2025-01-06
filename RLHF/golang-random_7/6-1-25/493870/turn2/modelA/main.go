package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("# HTTP Status Codes")

	// Iterate over the HTTP status codes and descriptions
	for code := 100; code < 600; code++ {
		description := http.StatusText(code)
		if description != "" {
			fmt.Printf("- %d: %s\n", code, description)
		}
	}
}
