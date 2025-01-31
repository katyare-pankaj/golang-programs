package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// handler responds to HTTP requests by generating dynamic HTML content.
func handler(w http.ResponseWriter, r *http.Request) {
	// Sanitize user input to prevent XSS (Cross-Site Scripting) attacks
	userInput := r.URL.Query().Get("name")
	safeInput := sanitizeInput(userInput)

	// Use fmt.Sprintf to generate dynamic HTML content
	htmlContent := fmt.Sprintf("<html><body><h1>Hello, %s!</h1></body></html>", safeInput)

	// Write the generated HTML content to the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, htmlContent)
}

// sanitizeInput is a utility function to escape potentially harmful input
func sanitizeInput(input string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
		"'", "&#39;",
	)
	return replacer.Replace(input)
}

func main() {
	// Set up the HTTP handler for the root path
	http.HandleFunc("/", handler)

	// Start the web server on port 8080 and log any potential errors
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
