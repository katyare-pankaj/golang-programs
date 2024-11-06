package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// State represents the shared state that needs to be synchronized across sites
type State struct {
	mu      sync.Mutex
	value   int
	version int
}

// updateState attempts to update the state with a new value.
// It returns true if the update was successful, false otherwise.
func (s *State) updateState(newValue int, expectedVersion int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.version != expectedVersion {
		// State has been modified since we read it, fail the update
		return false
	}

	s.value = newValue
	s.version++
	return true
}

// getState returns the current state value and version
func (s *State) getState() (int, int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.value, s.version
}

// simulateCrossSiteSynchronization simulates updating the state on different sites with latency
func simulateCrossSiteSynchronization(state *State, siteID int, updateInterval time.Duration, latencyVariability time.Duration) {
	for {
		// Generate a random update value and latency
		updateValue := rand.Intn(10) + 1
		latency := time.Duration(rand.Int63n(int64(latencyVariability))) * time.Millisecond

		// Read the current state
		currentValue, currentVersion := state.getState()

		// Simulate latency
		time.Sleep(latency)

		// Attempt to update the state
		if state.updateState(currentValue+updateValue, currentVersion) {
			fmt.Printf("Site %d updated state to: %d (version: %d)\n", siteID, currentValue+updateValue, currentVersion+1)
		} else {
			fmt.Printf("Site %d failed to update state, conflict detected.\n", siteID)
		}

		// Wait for the next update interval
		time.Sleep(updateInterval)
	}
}

func main() {
	// Initialize the shared state
	sharedState := &State{}

	// Define parameters for simulation
	numSites := 3
	updateInterval := 2 * time.Second
	latencyVariability := 500 * time.Millisecond

	// Start simulation for multiple sites
	for siteID := 1; siteID <= numSites; siteID++ {
		go simulateCrossSiteSynchronization(sharedState, siteID, updateInterval, latencyVariability)
	}

	// Let the simulation run for a while
	time.Sleep(10 * time.Second)
	fmt.Println("Simulation finished.")
}
