package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate retrieving a user from a database or another data source
	user := User{ID: 1, Name: "John Doe", Age: 30}

	// Marshal the user struct to JSON
	userBytes, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshaling user data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userBytes)
}

func main() {
	http.HandleFunc("/user", userHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
