package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type User struct {
	Username string
	Password string
	Role     string
}

type Session struct {
	Username string
	Expiry   time.Time
}

var (
	users = map[string]User{
		"admin": {"admin", "password", "admin"},
		"user":  {"user", "password", "user"},
	}

	sessionStore = make(map[string]Session)
	mu           sync.Mutex
)

const sessionDuration = 30 * time.Minute

func createSession(username string) string {
	sessionID := generateSessionID()
	mu.Lock()
	defer mu.Unlock()
	sessionStore[sessionID] = Session{Username: username, Expiry: time.Now().Add(sessionDuration)}
	return sessionID
}

func getSession(sessionID string) (Session, bool) {
	mu.Lock()
	defer mu.Unlock()
	session, exists := sessionStore[sessionID]
	if !exists || session.Expiry.Before(time.Now()) {
		delete(sessionStore, sessionID)
		return Session{}, false
	}
	return session, true
}

func generateSessionID() string {
	// Implementation-dependent: Consider using a UUID or cryptographic random generator
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, exists := users[username]
	if !exists || user.Password != password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	sessionID := createSession(username)
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionID,
		Expires: time.Now().Add(sessionDuration),
	})
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil || cookie == nil {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		session, valid := getSession(cookie.Value)
		if !valid {
			http.Error(w, "Session expired or invalid", http.StatusUnauthorized)
			return
		}

		r.Header.Set("Username", session.Username)
		next(w, r)
	}
}

func roleMiddleware(allowedRoles ...string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			username := r.Header.Get("Username")
			user, exists := users[username]
			if !exists {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			for _, role := range allowedRoles {
				if user.Role == role {
					next(w, r)
					return
				}
			}

			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	}
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome, Admin!"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome, User!"))
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/admin", authMiddleware(roleMiddleware("admin")(adminHandler)))
	http.HandleFunc("/user", authMiddleware(roleMiddleware("user", "admin")(userHandler)))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
