package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// User structure for storing user information.
type User struct {
	Username string
	Password string
}

// Session structure for session management.
type Session struct {
	Username string
	Expiry   time.Time
}

var (
	users = map[string]User{
		"admin": {"admin", "password"},
		"user":  {"user", "password"},
	}

	// Lifetime of session before it expires.
	sessionDuration = 5 * time.Minute

	sessionStore = make(map[string]Session)
	sessionsLock sync.Mutex
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/welcome", authMiddleware(welcomeHandler))

	// Minute-based cleanup
	go sessionCleanupTask()

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// loginHandler provides session creation after validating user credentials.
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
		Name:     "session_token",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // This should be true in a production environment with HTTPS
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(sessionDuration),
	})

	w.Write([]byte("Login successful"))
	log.Printf("User %s logged in\n", username)
}

// logoutHandler clears session data, logging the user out.
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err == nil {
		invalidateSession(cookie.Value)
		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Unix(0, 0),
		})
	}

	w.Write([]byte("Logout successful"))
}

// welcomeHandler is a protected endpoint that acknowledges user access.
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("Username")
	w.Write([]byte(fmt.Sprintf("Welcome %s!", username)))
	log.Printf("User %s accessed welcome page\n", username)
}

// Creates a new session and returns its ID.
func createSession(username string) string {
	sessionID := generateSessionID()
	sessionsLock.Lock()
	defer sessionsLock.Unlock()
	sessionStore[sessionID] = Session{
		Username: username,
		Expiry:   time.Now().Add(sessionDuration),
	}
	return sessionID
}

// Invalidates a session by removing it from the session store.
func invalidateSession(sessionID string) {
	sessionsLock.Lock()
	defer sessionsLock.Unlock()
	delete(sessionStore, sessionID)
}

// Generates a session ID. Implement this securely in production environments.
func generateSessionID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// Module responsible for continuously removing expired sessions.
func sessionCleanupTask() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			now := time.Now()
			sessionsLock.Lock()
			for sessionID, session := range sessionStore {
				if session.Expiry.Before(now) {
					delete(sessionStore, sessionID)
				}
			}
			sessionsLock.Unlock()
		}
	}
}

// Middleware to protect routes requiring authentication.
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil || cookie == nil {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		sessionsLock.Lock()
		session, exists := sessionStore[cookie.Value]
		sessionsLock.Unlock()

		if !exists || session.Expiry.Before(time.Now()) {
			http.Error(w, "Session expired or not found", http.StatusUnauthorized)
			return
		}

		// Propagate session and renew session expiry
		sessionsLock.Lock()
		session.Expiry = time.Now().Add(sessionDuration)
		sessionStore[cookie.Value] = session
		sessionsLock.Unlock()

		// Attach user information to request's context or headers
		r.Header.Set("Username", session.Username)
		next(w, r)
	}
}
