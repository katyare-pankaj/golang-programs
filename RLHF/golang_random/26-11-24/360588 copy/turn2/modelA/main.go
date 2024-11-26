package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username  string    `json:"username"`
	ExpiresAt time.Time `json:"expires_at"`
	jwt.StandardClaims
}

const jwtSecret = "your_secret_key" // Use a secure secret key in production

func authHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body
	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Simulate password verification (replace this with your authentication logic)
	if payload.Username == "example_user" && payload.Password == "example_password" {
		// Create claims
		claims := &Claims{
			Username:  payload.Username,
			ExpiresAt: time.Now().Add(15 * time.Minute), // Token expires after 15 minutes
			StandardClaims: jwt.StandardClaims{
				Issuer:    "example.com",
				IssuedAt:  time.Now().Unix(),
				Audience:  "your_application",
				NotBefore: time.Now().Unix(),
			},
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(jwtSecret))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Respond with the token
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"access_token": tokenString})
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func main() {
	http.HandleFunc("/auth/token", authHandler)
	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
