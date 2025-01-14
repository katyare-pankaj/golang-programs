package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	cookieName      = "votingSession"
	cookieMaxAge    = 3600 * 24 * 30 // 30 days
	votingSessionID = "votingSessionID"
)

type vote struct {
	ID      int
	Options []string
	Votes   map[string]int
	mutex   sync.RWMutex
	voted   map[string]struct{}
}

func (v *vote) addVote(sessionID string, option string) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	if _, ok := v.voted[sessionID]; ok {
		return
	}

	if _, ok := v.Votes[option]; !ok {
		v.Votes[option] = 0
	}
	v.Votes[option]++
	v.voted[sessionID] = struct{}{}
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
	return fmt.Sprintf("<ul>%s</ul>", strings.Join(results, "<li></li>"))
}

var votingSystem *vote

func main() {
	votingSystem = &vote{
		ID:      1,
		Options: []string{"Option A", "Option B", "Option C"},
		Votes:   make(map[string]int),
		mutex:   sync.RWMutex{},
		voted:   make(map[string]struct{}),
	}

	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/vote", handleVote)
	http.HandleFunc("/results", handleResults)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	sessionID := getSessionID(w, r)

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

	sessionID := getSessionID(w, r)
	votingSystem.addVote(sessionID, option)

	fmt.Fprintf(w, "Thank you for voting for %s!", option)
}
