package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type survey struct {
	ID        int
	Name      string
	Questions []question
	Results   map[string]map[string]int
	mutex     sync.RWMutex
}

type question struct {
	ID      int
	Text    string
	Options []string
}

func (s *survey) addResponse(qID string, option string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.Results[qID]; !ok {
		s.Results[qID] = make(map[string]int)
	}

	s.Results[qID][option]++
}

var surveyApp *survey

func main() {
	surveyApp = &survey{
		ID:   1,
		Name: "Test Survey",
		Questions: []question{
			{
				ID:      1,
				Text:    "What is your favorite programming language?",
				Options: []string{"Go", "Python", "JavaScript", "Java"},
			},
			{
				ID:      2,
				Text:    "How many years of experience do you have?",
				Options: []string{"Less than 1 year", "1-3 years", "3-5 years", "More than 5 years"},
			},
		},
		Results: make(map[string]map[string]int),
	}

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/survey", handleSurvey)
	http.HandleFunc("/results", handleResults)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Survey App!")
}

func handleSurvey(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for qID, options := range r.PostForm {
		for _, option := range options {
			surveyApp.addResponse(qID, option)
		}
	}

	fmt.Fprintf(w, "Thank you for submitting your response!")
}

func handleResults(w http.ResponseWriter, r *http.Request) {
	surveyApp.mutex.RLock()
	defer surveyApp.mutex.RUnlock()

	fmt.Fprintf(w, "Survey Results:")
	for _, q := range surveyApp.Questions {
		fmt.Fprintf(w, "\n\nQuestion %d: %s", q.ID, q.Text)
		for _, option := range q.Options {
			count := surveyApp.Results[fmt.Sprintf("%d", q.ID)][option]
			fmt.Fprintf(w, "\n\t%s: %d", option, count)
		}
	}
}
