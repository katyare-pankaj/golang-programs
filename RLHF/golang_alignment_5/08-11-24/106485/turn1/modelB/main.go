package main

import (
	"fmt"
	"sync"
)

// Define the SessionState struct to represent the state of a session
type SessionState struct {
	// You can add any fields here that represent the state of the session.
	// For example, you might have count of events for a specific session,
	// timestamps, and other data relevant to your CEP logic.
	count                int64
	latestEventTimestamp int64
}

type SessionStateManager struct {
	// Mutex to protect concurrent access to the sessionStateMap
	mu sync.Mutex
	// Use a map to store session states. The key can be any type that can be hashed.
	// For example, you could use a string, int, or a custom struct.
	sessionStateMap map[string]SessionState
}

func (ssm *SessionStateManager) GetOrCreateSessionState(sessionKey string) *SessionState {
	ssm.mu.Lock()
	defer ssm.mu.Unlock()

	state, exists := ssm.sessionStateMap[sessionKey]
	if !exists {
		// Create a new session state if it doesn't exist
		state = SessionState{}
		ssm.sessionStateMap[sessionKey] = state
	}
	return &state
}

func (ssm *SessionStateManager) UpdateSessionState(sessionKey string, updateFunc func(*SessionState)) {
	ssm.mu.Lock()
	defer ssm.mu.Unlock()

	state, exists := ssm.sessionStateMap[sessionKey]
	if exists {
		// Update the session state using the provided update function
		updateFunc(&state)
		ssm.sessionStateMap[sessionKey] = state
	}
}

func (ssm *SessionStateManager) RemoveSessionState(sessionKey string) {
	ssm.mu.Lock()
	defer ssm.mu.Unlock()

	delete(ssm.sessionStateMap, sessionKey)
}

func main() {
	// Create an instance of the SessionStateManager
	ssm := SessionStateManager{sessionStateMap: make(map[string]SessionState)}

	// Simulate processing events for different sessions
	for eventTimestamp := 0; eventTimestamp < 100; eventTimestamp++ {
		// Get or create the session state for a specific session
		sessionKey := fmt.Sprintf("session-%d", eventTimestamp%5) // For simplicity, we'll use a modulo-based session key
		state := ssm.GetOrCreateSessionState(sessionKey)

		// Update the session state with the event data
		ssm.UpdateSessionState(sessionKey, func(s *SessionState) {
			s.count++
			s.latestEventTimestamp = int64(eventTimestamp)
		})

		// Perform any CEP processing based on the updated session state
		// For example, you might check if the state meets certain conditions to trigger an action.
		if state.count == 3 {
			fmt.Printf("Session %s triggered at timestamp %d\n", sessionKey, eventTimestamp)
		}
	}
}
