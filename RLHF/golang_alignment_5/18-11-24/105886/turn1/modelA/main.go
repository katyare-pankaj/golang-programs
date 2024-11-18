package main

import (
	"html/template"
	"log"
	"net/http"
)

// NewTemplate returns a new template.
func NewTemplate() *template.Template {
	tpl := template.Must(template.ParseGlob("templates/*.html"))
	// Regular expression-based func for parsing pagination URLs
	tpl.Funcs(template.FuncMap{
		"pagination": func(page int, itemsPerPage int, totalItems int) string {
			// Construct pagination URL logic
			return ""
		},
	})
	// PageTitle    function
	tpl.Funcs(template.FuncMap{"PageTitle": PageTitle})
	return tpl
}
func PageTitle(title string) string {
	return "MyApp - " + title
}

// HandleRequest handles all incoming requests
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Hello World",
		Message: "This is a Go website!",
	}

	tpl := NewTemplate()
	err := tpl.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		log.Fatal("Error executing template:", err)
	}
}
func main() {
	http.HandleFunc("/", HandleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
