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
	quitChan := make(chan int)

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

				// Send the updated state back to the connected clients
				BroadcastState(gameState.GetState())
			case <-quitChan:
				return
			default:
				time.Sleep(50 * time.Millisecond) // Control the loop execution frequency
			}
		}
	}()

	// Simulate network connections updating game state
	go func() {
		playerID := 1
		for {
			select {
			case <-quitChan:
				return
			default:
				newPosition := time.Now().Second()
				fmt.Printf("Simulating player %d moving to position %d\n", playerID, newPosition)
				stateUpdateChan <- map[int]int{playerID: newPosition}
				time.Sleep(2 * time.Second)
			}
		}
	}()

	// Simulate game running for some time
	time.Sleep(10 * time.Second)
	close(quitChan)
}

// BroadcastState sends the game state to all clients
func BroadcastState(state map[int]int) {
	fmt.Println("Broadcasting game state to clients:", state)
}
