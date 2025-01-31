package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// handler responds to HTTP requests and generates dynamic HTML content.
func handler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the 'name' parameter from the URL query
	userInput := r.URL.Query().Get("name")

	// Use template.HTMLEscapeString to escape potentially dangerous characters
	safeInput := template.HTMLEscapeString(userInput)

	// Generate the HTML using fmt.Sprintf
	htmlContent := fmt.Sprintf("<html><body><h1>Hello, %s!</h1></body></html>", safeInput)

	// Set the content type and write the generated HTML to the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, htmlContent)
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
