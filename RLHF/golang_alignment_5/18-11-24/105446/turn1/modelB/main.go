package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{
		Username: "admin",
		Password: "$2b$10$IhPvOgUMU9/PlY.m/vYUmOxEu6TZl0h7/P/5a9ljjxG/TPjZ3Kx12", // Hashed password for "admin"
	},
}

func main() {
	http.HandleFunc("/login", loginHandler)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var user User
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Input validation
	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Authentication
	for _, u := range users {
		if u.Username == user.Username {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			if bcrypt.CompareHashAndPassword([]byte(u.Password), hash) == nil {
				w.Write([]byte("Login successful"))
				return
			}
		}
	}

	http.Error(w, "Invalid username or password", http.StatusUnauthorized)
}
