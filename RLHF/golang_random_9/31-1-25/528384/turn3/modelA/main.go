package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// formHandler serves the HTML form
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Serve a simple HTML form
	htmlForm := `
		<html>
		<body>
			<form action="/submit" method="post">
				Name: <input type="text" name="name">
				<input type="submit" value="Submit">
			</form>
		</body>
		</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, htmlForm)
}

// submitHandler processes form submissions
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Retrieve and sanitize the 'name' input
	userInput := r.FormValue("name")
	safeInput := template.HTMLEscapeString(userInput)

	// Generate the response HTML
	htmlContent := fmt.Sprintf("<html><body><h1>Hello, %s!</h1></body></html>", safeInput)

	// Send the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, htmlContent)
}

func main() {
	// Route to serve the form
	http.HandleFunc("/", formHandler)

	// Route to handle form submissions
	http.HandleFunc("/submit", submitHandler)

	// Start the web server on port 8080
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
