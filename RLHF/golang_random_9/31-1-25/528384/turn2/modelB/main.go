package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Define a struct to hold the template data
type TemplateData struct {
	UserName template.HTML
	Message  string
	Age      int
}

// Create a global variable for the template
var tmpl *template.Template

func init() {
	// Parse the template file
	tmpl = template.Must(template.ParseFiles("template.html"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Create a TemplateData struct to hold the data
	data := TemplateData{}

	// Extract user input from the query string
	userName := r.URL.Query().Get("name")

	// Sanitize user input before passing it to the template
	data.UserName = template.HTMLEscapeString(userName)

	// Validate and handle user input
	ageStr := r.URL.Query().Get("age")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		// Handle invalid age input gracefully
		data.Message = "Invalid age provided."
	} else {
		data.Age = age
	}

	// Execute the template and write the result to the response
	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
