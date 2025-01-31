package main

import (
	"fmt"
	"sync"
	"time"
)

// GameState holds the state of the game
type GameState struct {
	PlayerPositions map[int]int
	mutex           sync.Mutex
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
	broadcastChan := make(chan map[int]int)

	// Simulate clients that listen for broadcasted game state updates
	clientCount := 3
	for i := 0; i < clientCount; i++ {
		go func(clientID int) {
			for {
				select {
				case updatedState := <-broadcastChan:
					fmt.Printf("Client %d received updated game state: %v\n", clientID, updatedState)
				case <-quitChan:
					return
				}
			}
		}(i + 1)
	}

	// Run game loop in a separate goroutine
	go func() {
		for {
			select {
			case newState := <-stateUpdateChan:
				// Update game state
				gameState.mutex.Lock()
				for playerID, position := range newState {
					gameState.PlayerPositions[playerID] = position
				}
				gameState.mutex.Unlock()

				// Broadcast the updated state to all clients
				broadcastChan <- gameState.GetState()

			case <-quitChan:
				return

			default:
				time.Sleep(50 * time.Millisecond)
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
					newPosition := (time.Now().Second() * id) % 100 // Sample formula for positioning
					stateUpdateChan <- map[int]int{id: newPosition}
					time.Sleep(time.Second)
				}
			}
		}(playerID)
	}

	// Simulate game running for some time
	time.Sleep(10 * time.Second)
	close(quitChan)
	wg.Wait()
}
