package main

import (
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

	// Create a dynamic API response using fmt.Sprintf
	response := fmt.Sprintf("User found: { ID: %d, Name: %s, Age: %d }", user.ID, user.Name, user.Age)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/user", userHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
