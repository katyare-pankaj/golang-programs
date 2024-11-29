// handlers/user.go
package handlers

import (
	"encoding/json"
	"go-programs/RLHF/golang_random/29-11-24/389143/turn1/modelA/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{
		{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
		},
		{
			ID:    2,
			Name:  "Jane Smith",
			Email: "jane@example.com",
		},
	}

	json.NewEncoder(w).Encode(users)
}

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users", ListUsers).Methods("GET")

	// added listener fo testing
	http.ListenAndServe(":8080", r)
}
