package main

import (
	"fmt"
	"sync"
	"time"
)

// GameState holds the state of the game
type GameState struct {
	PlayerPositions map[int]int // Map of player ID to their position
	mutex           sync.Mutex  // Protects the game state
}

// NewGameState initializes a new GameState
func NewGameState() *GameState {
	return &GameState{
		PlayerPositions: make(map[int]int),
	}
}

// UpdatePosition updates the position of a player
func (gs *GameState) UpdatePosition(playerID, position int) {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()
	gs.PlayerPositions[playerID] = position
}

// GetState returns a copy of the current game state
func (gs *GameState) GetState() map[int]int {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()
	copyOfState := make(map[int]int)
	for k, v := range gs.PlayerPositions {
		copyOfState[k] = v
	}
	return copyOfState
}

func main() {
	gameState := NewGameState()
	stateUpdateChan := make(chan map[int]int)
	quitChan := make(chan struct{})
	clients := make(map[chan map[int]int]struct{}) // Map of client channels to empty struct

	// Run game loop in a separate goroutine
	go func() {
		for {
			select {
			case newState := <-stateUpdateChan:
				gameState.mutex.Lock()
				for playerID, position := range newState {
					gameState.PlayerPositions[playerID] = position
				}
				gameState.mutex.Unlock()

				// Send the updated state to all clients
				for clientChan := range clients {
					clientChan <- gameState.GetState()
				}
			case <-quitChan:
				return
			default:
				time.Sleep(50 * time.Millisecond) // Control the loop execution frequency
			}
		}
	}()

	// Simulate multiple players updating their positions concurrently
	var wg sync.WaitGroup
	for playerID := 1; playerID <= 3; playerID++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-quitChan:
					return
				default:
					newPosition := time.Now().Second() * id
					fmt.Printf("Simulating player %d moving to position %d\n", id, newPosition)
					stateUpdateChan <- map[int]int{id: newPosition}
					time.Sleep(time.Duration(1+id) * time.Second)
				}
			}
		}(playerID)
	}

	// Simulate client connections
	go func() {
		for {
			select {
			case <-quitChan:
				return
			default:
				// Create a new client channel for each connection
				clientChan := make(chan map[int]int)
				clients[clientChan] = struct{}{} // Add the channel to the clients map
				go func() {
					for state := range clientChan {
						// Send the state update to the client
						fmt.Println("Client received state update:", state)
					}
					// Remove the client channel when it's closed
					delete(clients, clientChan)
				}()
				time.Sleep(2 * time.Second) // Simulate a new client connecting every 2 seconds
			}
		}
	}()

	// Simulate game running for some time
	time.Sleep(10 * time.Second)
	close(quitChan)
	wg.Wait()
}
