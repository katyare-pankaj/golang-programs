package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// Define your database structure
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	SessionID string `json:"session_id"`
}

type Session struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	SessionID string `json:"session_id"`
}

var db *sql.DB

func main() {
	// Initialize database
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/database_name")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create new router
	r := mux.NewRouter()

	// Login endpoint
	r.HandleFunc("/api/login", loginHandler).Methods("POST")

	// Middleware to check for valid session and RBAC
	r.Use(authMiddleware)

	// Protected endpoints with RBAC
	r.HandleFunc("/api/admin", adminHandler).Methods("GET").Queries("role", "admin")
	r.HandleFunc("/api/user", userHandler).Methods("GET").Queries("role", "user")

	// Logout endpoint
	r.HandleFunc("/api/logout", logoutHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

// User authentication
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate credentials
	var dbUser User
	row := db.QueryRow("SELECT id, username, password, role FROM users WHERE username = ?", user.Username)
	if err := row.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password, &dbUser.Role); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Compare password hash
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate new session
	sessionID, err := generateRandomString(32)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Store session in database
	_, err = db.Exec("INSERT INTO sessions (user_id, session_id) VALUES (?, ?)", dbUser.ID, sessionID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set session cookie
	cookie := http.Cookie{
		Name:  "session",
		Value: sessionID,
		Path:  "/",
	}
	http.SetCookie(w, &cookie)

	// Return user information
	user.Password = ""
	user.Role = dbUser.Role
	user.SessionID = sessionID
	json.NewEncoder(w).Encode(user)
}
