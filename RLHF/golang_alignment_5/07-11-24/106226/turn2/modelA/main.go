package main

import (
	"fmt"
	"sync"
	"time"
)

// Session represents a single session with an expiration time
type Session struct {
	Data       interface{}
	Expiration time.Time
}

// SessionManager manages sessions using a sync.Map
type SessionManager struct {
	sessions sync.Map
	duration time.Duration
	ticker   *time.Ticker
	quit     chan struct{}
}

// NewSessionManager creates a new SessionManager with a specified expiration duration
func NewSessionManager(duration time.Duration) *SessionManager {
	sm := &SessionManager{
		duration: duration,
		quit:     make(chan struct{}),
	}
	sm.ticker = time.NewTicker(time.Minute)
	go sm.run()
	return sm
}

// Stop the session manager
func (sm *SessionManager) Stop() {
	close(sm.quit)
	sm.ticker.Stop()
}

// run is a background goroutine that expires sessions periodically
func (sm *SessionManager) run() {
	for {
		select {
		case <-sm.ticker.C:
			sm.expireSessions()
		case <-sm.quit:
			return
		}
	}
}

// expireSessions removes all sessions that have expired
func (sm *SessionManager) expireSessions() {
	now := time.Now()
	sm.sessions.Range(func(key, value interface{}) bool {
		session := value.(*Session)
		if session.Expiration.Before(now) {
			sm.sessions.Delete(key)
		}
		return true
	})
}

// SetSession sets a new session or updates an existing one
func (sm *SessionManager) SetSession(key string, data interface{}) {
	expiration := time.Now().Add(sm.duration)
	session := &Session{Data: data, Expiration: expiration}
	sm.sessions.Store(key, session)
}

// GetSession retrieves a session by key
func (sm *SessionManager) GetSession(key string) (interface{}, bool) {
	value, ok := sm.sessions.Load(key)
	if !ok {
		return nil, false
	}
	session := value.(*Session)
	if session.Expiration.Before(time.Now()) {
		sm.sessions.Delete(key)
		return nil, false
	}
	return session.Data, true
}

// DeleteSession deletes a session by key
func (sm *SessionManager) DeleteSession(key string) {
	sm.sessions.Delete(key)
}

func main() {
	// Create a session manager with a session expiration of 10 seconds
	sessionManager := NewSessionManager(10 * time.Second)
	defer sessionManager.Stop()

	// Set some sessions
	sessionManager.SetSession("user1", "some data for user1")
	sessionManager.SetSession("user2", "some data for user2")

	// Retrieve and print sessions
	data, ok := sessionManager.GetSession("user1")
	if ok {
		fmt.Println("User1 session data:", data)
	} else {
		fmt.Println("User1 session expired")
	}

	// Wait for sessions to expire
	time.Sleep(15 * time.Second)

	data, ok = sessionManager.GetSession("user1")
	if ok {
		fmt.Println("User1 session data:", data)
	} else {
		fmt.Println("User1 session expired")
	}
}
