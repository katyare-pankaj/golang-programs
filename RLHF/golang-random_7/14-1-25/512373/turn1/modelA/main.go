package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Survey represents a single survey response
type Survey struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Age   int       `json:"age"`
	Date  time.Time `json:"date"`
}

var surveys []Survey

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/survey", surveyHandler)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/api/surveys", getSurveysHandler)

	log.Println("Survey application listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// homeHandler serves the homepage
func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("home.html")
	t.Execute(w, nil)
}

// surveyHandler serves the survey form
func surveyHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("survey.html")
	t.Execute(w, nil)
}

// submitHandler processes the survey submission
func submitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	survey := Survey{
		ID:    generateID(),
		Name:  r.FormValue("name"),
		Email: r.FormValue("email"),
		Age:   atoi(r.FormValue("age")),
		Date:  time.Now(),
	}
	surveys = append(surveys, survey)
	http.Redirect(w, r, "/api/surveys", http.StatusFound)
}

// getSurveysHandler serves all survey responses
func getSurveysHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(surveys); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// generateID generates a unique ID for a survey response
func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// atoi attempts to convert a string to an integer
func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
