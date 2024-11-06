package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define the ServerState struct to hold the server's state
type ServerState struct {
	Value int
}

// Define the LocalState struct to hold the local state and a buffer for pending changes
type LocalState struct {
	Value  int
	buffer []int
	mutex  sync.Mutex
}

func main() {
	// Create an instance of the ServerState and LocalState
	serverState := &ServerState{Value: 0}
	localState := &LocalState{Value: 0}
	// Start a goroutine for buffering local changes
	go bufferLocalChanges(localState)
	// Start a goroutine for synchronization with the server
	go synchronizeWithServer(localState, serverState)
	// Simulate local updates to the state
	for i := 0; i < 10; i++ {
		updateLocalState(localState, rand.Intn(10)+1)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

// Update the local state
func updateLocalState(localState *LocalState, value int) {
	localState.mutex.Lock()
	defer localState.mutex.Unlock()
	localState.Value += value
	localState.buffer = append(localState.buffer, value)
}

// Buffer local changes and apply them to the server state when the connection is stable
func bufferLocalChanges(localState *LocalState) {
	for {
		// Check if the buffer is empty
		localState.mutex.Lock()
		if len(localState.buffer) == 0 {
			localState.mutex.Unlock()
			time.Sleep(time.Second) // Wait for a second before checking again
			continue
		}
		// Buffer is not empty, so apply the changes to the server state
		fmt.Println("Buffering changes...")
		for _, change := range localState.buffer {
			localState.Value += change
		}
		localState.buffer = nil
		localState.mutex.Unlock()
		// Simulate a synchronization with the server
		fmt.Println("Synchronizing with the server...")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

// Synchronize the local state with the server state
func synchronizeWithServer(localState *LocalState, serverState *ServerState) {
	for {
		// Synchronize the local state with the server state
		localState.mutex.Lock()
		if localState.Value != serverState.Value {
			fmt.Println("Syncing local state with server state...")
			serverState.Value = localState.Value
		}
		localState.mutex.Unlock()
		// Wait for a random period before synchronizing again
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}
