package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// User struct for representing user data
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: "1", Name: "Alice", Email: "alice@example.com"},
	{ID: "2", Name: "Bob", Email: "bob@example.com"},
}

// validateEmail function to validate email address
func validateEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("id")

	for _, user := range users {
		if user.ID == userID {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
}

func main() {
	r := httprouter.New()

	r.GET("/users/:id", getUser)

	log.Fatal(http.ListenAndServe(":8080", r))
}
