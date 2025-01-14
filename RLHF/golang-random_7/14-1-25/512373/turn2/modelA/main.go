package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type vote struct {
	ID      int
	Options []string
	Votes   map[string]int
	mutex   sync.RWMutex
}

func (v *vote) addVote(option string) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	if _, ok := v.Votes[option]; !ok {
		v.Votes[option] = 0
	}
	v.Votes[option]++
}

func (v *vote) getResults() string {
	v.mutex.RLock()
	defer v.mutex.RUnlock()

	results := make([]string, len(v.Votes))
	i := 0
	for option, count := range v.Votes {
		results[i] = fmt.Sprintf("%s: %d", option, count)
		i++
	}
	return fmt.Sprintf("%s", results)
}

var votingSystem *vote

func main() {
	votingSystem = &vote{
		ID:      1,
		Options: []string{"Option A", "Option B", "Option C"},
		Votes:   make(map[string]int),
		mutex:   sync.RWMutex{},
	}

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/vote", handleVote)
	http.HandleFunc("/results", handleResults)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Voting System</h1>")
	fmt.Fprintf(w, "<form action='/vote' method='post'>")
	for _, option := range votingSystem.Options {
		fmt.Fprintf(w, "<input type='radio' name='option' value='%s'> %s<br>", option, option)
	}
	fmt.Fprintf(w, "<input type='submit' value='Vote'>")
	fmt.Fprintf(w, "</form>")
}

func handleVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	option := r.FormValue("option")
	if option == "" {
		http.Error(w, "Please select an option", http.StatusBadRequest)
		return
	}

	votingSystem.addVote(option)
	fmt.Fprintf(w, "Thank you for voting for %s!", option)
}

func handleResults(w http.ResponseWriter, r *http.Request) {
	results := votingSystem.getResults()
	fmt.Fprintf(w, "<h1>Voting Results</h1>")
	fmt.Fprintf(w, "%s", results)
}
