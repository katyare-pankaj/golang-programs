package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type vote struct {
	option string
	count  int
}

var votes []vote
var voteMutex sync.Mutex

func init() {
	votes = []vote{
		{option: "Red", count: 0},
		{option: "Green", count: 0},
		{option: "Blue", count: 0},
	}
}
func handleVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	option := r.FormValue("option")
	voteMutex.Lock()
	defer voteMutex.Unlock()

	for i, v := range votes {
		if v.option == option {
			votes[i].count++
			break
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func handleShowVotes(w http.ResponseWriter, r *http.Request) {
	voteMutex.Lock()
	defer voteMutex.Unlock()

	fmt.Fprintf(w, "<h1>Vote Results:</h1><ul>")
	for _, v := range votes {
		fmt.Fprintf(w, "<li>%s: %d</li>", v.option, v.count)
	}
	fmt.Fprintf(w, "</ul>")
}

func main() {
	http.HandleFunc("/vote", handleVote)
	http.HandleFunc("/", handleShowVotes)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
