package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Mock database
var usersDB = map[int]User{
	1: {ID: 1, Name: "Alice", Email: "alice@example.com"},
	2: {ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// Fetch user by ID (in reality, this might interact with a database or legacy monolith)
func GetUser(userID int, wg *sync.WaitGroup, resultChan chan<- *User, errChan chan<- error) {
	defer wg.Done()
	user, exists := usersDB[userID]
	if !exists {
		errChan <- fmt.Errorf("user not found: %d", userID)
		return
	}
	resultChan <- &user
}

// UserHandler handles requests for a specific user
func UserHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("id")
	userID := 0 // Convert ID from string to int, error handling is omitted for brevity.

	// Run GetUser in a goroutine to handle the request concurrently
	var wg sync.WaitGroup
	resultChan := make(chan *User)
	errChan := make(chan error)

	wg.Add(1)
	go GetUser(userID, &wg, resultChan, errChan)

	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	// Handling responses
	select {
	case user := <-resultChan:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	case err := <-errChan:
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/user", UserHandler)
	fmt.Println("User service is running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
