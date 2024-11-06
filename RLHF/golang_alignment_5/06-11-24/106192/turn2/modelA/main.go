package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ServerState represents the shared state managed by the server.
type ServerState struct {
	Value int
	mutex sync.Mutex
}

// LocalState manages the local state of a device and buffers changes before synchronizing.
type LocalState struct {
	Value  int
	buffer []int
	mutex  sync.Mutex
}

// Synchronizer orchestrates the buffering and synchronization of local states with the server.
type Synchronizer struct {
	localState  *LocalState
	serverState *ServerState
	stopCh      chan struct{}
	wg          sync.WaitGroup
}

// NewSynchronizer creates a new Synchronizer instance.
func NewSynchronizer(localState *LocalState, serverState *ServerState) *Synchronizer {
	s := &Synchronizer{
		localState:  localState,
		serverState: serverState,
		stopCh:      make(chan struct{}),
	}
	s.wg.Add(2)
	go s.bufferLocalChanges()
	go s.synchronizeWithServer()
	return s
}

// Stop signals the Synchronizer to stop its goroutines.
func (s *Synchronizer) Stop() {
	close(s.stopCh)
	s.wg.Wait()
}

// bufferLocalChanges periodically checks for buffered changes and applies them to the local state
// after a simulated stable connection period.
func (s *Synchronizer) bufferLocalChanges() {
	defer s.wg.Done()
	for {
		select {
		case <-s.stopCh:
			return
		default:
		}

		s.localState.mutex.Lock()
		if len(s.localState.buffer) == 0 {
			s.localState.mutex.Unlock()
			time.Sleep(time.Second) // Simulate checking interval
			continue
		}

		// Apply buffered changes to the local state
		fmt.Println("Buffering changes...")
		for _, change := range s.localState.buffer {
			s.localState.Value += change
		}
		s.localState.buffer = nil
		s.localState.mutex.Unlock()

		// Simulate stable connection period before synchronizing
		fmt.Println("Synchronizing with the server...")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

// synchronizeWithServer synchronizes the local state with the server state periodically.
func (s *Synchronizer) synchronizeWithServer() {
	defer s.wg.Done()
	for {
		select {
		case <-s.stopCh:
			return
		default:
		}

		s.localState.mutex.Lock()
		localValue := s.localState.Value
		s.localState.mutex.Unlock()

		s.serverState.mutex.Lock()
		if localValue != s.serverState.Value {
			fmt.Println("Syncing local state with server state...")
			s.serverState.Value = localValue
		}
		s.serverState.mutex.Unlock()

		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

// updateLocalState updates the local state and adds the change to the buffer.
func updateLocalState(localState *LocalState, value int) {
	localState.mutex.Lock()
	defer localState.mutex.Unlock()
	localState.Value += value
	localState.buffer = append(localState.buffer, value)
}

func main() {
	// Create instances of ServerState and LocalState
	serverState := &ServerState{}
	localState := &LocalState{}

	// Create a Synchronizer to manage buffering and synchronization
	synchronizer := NewSynchronizer(localState, serverState)
	defer synchronizer.Stop()

	// Simulate local updates to the state
	for i := 0; i < 20; i++ {
		updateLocalState(localState, rand.Intn(10)+1)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}

	// Wait for some time to let the synchronization process complete
	time.Sleep(5 * time.Second)

	fmt.Println("Final Local State:", localState.Value)
	fmt.Println("Final Server State:", serverState.Value)
}
