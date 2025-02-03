package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string
	Role     string
}

var (
	users = map[string]User{
		"admin": {"admin", "$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi", "admin"}, // Hashed password for "password"
		"user":  {"user", "$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi", "user"},  // Hashed password for "password"
	}
	redisClient *redis.Client
)

const (
	sessionDuration     = 30 * time.Minute // Session duration
	sessionCheckInterval = 10 * time.Minute // Interval to check for expired sessions
	redisHost            = "localhost"     // Redis server host
	redisPort            = "6379"           // Redis server port
)

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "", // No password set
		DB:       0,  // Use default DB
	})

	// Start a goroutine to check for expired sessions
	go checkExpiredSessions()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("GET")
	r.HandleFunc("/dashboard", dashboardHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, exists := users[username]
	if !exists {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	sessionID := createSession(username)
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionID,
		Expires: time.Now().Add(sessionDuration),
		HttpOnly: true,
		Secure:   true,
	})

	w.Write([]byte("Login successful"))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Invalid session", http.StatusBadRequest)
		return
	}

	err = redisClient.Del(r.Context(), cookie.Value).Err()
	if err != nil {
		log.Printf("Error deleting session: %v", err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",