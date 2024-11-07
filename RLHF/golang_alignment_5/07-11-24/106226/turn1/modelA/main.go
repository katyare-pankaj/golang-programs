package main

import (
	"fmt"
	"sync"
	"time"
)

// Session represents a session state
type Session struct {
	data    map[string]interface{}
	expires time.Time
}

// NewSession creates a new session with a specified expiration duration
func NewSession(duration time.Duration) *Session {
	return &Session{
		data:    make(map[string]interface{}),
		expires: time.Now().Add(duration),
	}
}

// IsExpired checks if the session has expired
func (s *Session) IsExpired() bool {
	return time.Now().After(s.expires)
}

// Set stores a key-value pair in the session
func (s *Session) Set(key string, value interface{}) {
	s.data[key] = value
}

// Get retrieves a value from the session by key
func (s *Session) Get(key string) interface{} {
	return s.data[key]
}

// Delete removes a key-value pair from the session
func (s *Session) Delete(key string) {
	delete(s.data, key)
}

// SessionManager manages a map of sessions with expiration
type SessionManager struct {
	sessions        sync.Map
	duration        time.Duration
	cleanupInterval time.Duration
}

// NewSessionManager creates a new SessionManager
func NewSessionManager(duration, cleanupInterval time.Duration) *SessionManager {
	m := &SessionManager{
		duration:        duration,
		cleanupInterval: cleanupInterval,
	}
	go m.startCleanup()
	return m
}

// startCleanup periodically removes expired sessions
func (m *SessionManager) startCleanup() {
	ticker := time.NewTicker(m.cleanupInterval)
	defer ticker.Stop()
	for range ticker.C {
		m.sessions.Range(func(key, value interface{}) bool {
			session := value.(*Session)
			if session.IsExpired() {
				m.sessions.Delete(key)
			}
			return true
		})
	}
}

// GetSession retrieves a session by ID
func (m *SessionManager) GetSession(sessionID string) (*Session, bool) {
	session, ok := m.sessions.Load(sessionID)
	if !ok {
		return nil, false
	}
	return session.(*Session), true
}

// CreateSession creates a new session and stores it
func (m *SessionManager) CreateSession() string {
	sessionID := fmt.Sprintf("%x", time.Now().UnixNano())
	session := NewSession(m.duration)
	m.sessions.Store(sessionID, session)
	return sessionID
}

func main() {
	// Create a session manager with a session expiration of 1 minute and a cleanup interval of 30 seconds
	sessionManager := NewSessionManager(time.Minute, 30*time.Second)

	// Create
