package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	// Set the session expiration duration in minutes
	sessionExpirationDuration = time.Minute * 10
)

type Session struct {
	Data   interface{}
	Expiry time.Time
}

func main() {
	// Create a sync.Map to store sessions
	sessions := &sync.Map{}

	// Start the session expiration goroutine
	go manageSessionExpiration(sessions, sessionExpirationDuration)

	// Simulate some sessions
	simulateSessions(sessions)
}

// Simulate session creation and updates
func simulateSessions(sessions *sync.Map) {
	for sessionID := 0; sessionID < 10; sessionID++ {
		newSession := Session{
			Data:   map[string]interface{}{"user": "user" + fmt.Sprint(sessionID)},
			Expiry: time.Now().Add(sessionExpirationDuration),
		}
		sessions.Store(fmt.Sprint(sessionID), newSession)
		fmt.Printf("Session %d created.\n", sessionID)

		// Modify the session data after some time to demonstrate updates
		time.Sleep(time.Duration(sessionID*2) * time.Second)
		updateSession, _ := sessions.Load(fmt.Sprint(sessionID))
		if updateSession != nil {
			updatedSession := updateSession.(Session)
			updatedSession.Data.(map[string]interface{})["key"] = "value" + fmt.Sprint(sessionID)
			updatedSession.Expiry = time.Now().Add(sessionExpirationDuration)
			sessions.Store(fmt.Sprint(sessionID), updatedSession)
			fmt.Printf("Session %d updated.\n", sessionID)
		}
	}
}

// Manage session expiration in a separate goroutine
func manageSessionExpiration(sessions *sync.Map, expirationDuration time.Duration) {
	for {
		// Check for expired sessions every 30 seconds
		time.Sleep(30 * time.Second)
		now := time.Now()

		sessions.Range(func(key, value interface{}) bool {
			session := value.(Session)
			if session.Expiry.Before(now) {
				// Session expired, remove it from the map
				sessions.Delete(key)
				fmt.Printf("Session %s expired.\n", key)
			}
			return true // Continue the iteration
		})
	}
}
