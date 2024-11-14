package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jwtgo/jwt"
)

const (
	secretKey = "mySecretKey" // Use a stronger key in production
)

// Token is a struct representing a user token
type Token struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// Claims is a custom claim set for JWT tokens
type Claims struct {
	Token *Token `json:"token"`
	jwt.StandardClaims
}

// JWT contains the token value
type JWT struct {
	Token string `json:"token"`
}

func main() {
	r := mux.NewRouter()

	// Register middleware for logging
	r.Use(logger)

	// Register middleware for token validation
	r.Use(tokenValidator)

	// API endpoint that requires a valid token
	r.HandleFunc("/api/data", dataHandler).Methods("GET")

	// Initialize SSL certificates in real world
	http.ListenAndServe(":8080", r)
}

// logger is a middleware that logs incoming requests
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %s, URL: %s, RemoteAddr: %s", r.Method, r.URL, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// tokenValidator is a middleware that validates the JWT token in the request header
func tokenValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the token from the request header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Token not provided", http.StatusUnauthorized)
			return
		}

		// Split the token string to get the token value
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Method)
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Check if the token is valid
		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Set the user ID in the request context for future use
		ctx := context.WithValue(r.Context(), "userID", claims.Token.ID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// dataHandler is an API endpoint that requires a valid token
func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the user ID from the request context
	userID := r.Context().Value("userID").(string)

	fmt.Fprintf(w, "Hello, User ID: %s!", userID)
}
