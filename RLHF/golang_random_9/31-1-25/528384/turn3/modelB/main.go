package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// handler responds to HTTP requests, including form submissions.
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Serve the HTML form for GET requests
		http.ServeFile(w, r, "index.html")
		return
	}

	// Parse the form data for POST requests
	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Retrieve the 'name' parameter from the form
	userInput := r.FormValue("name")

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
