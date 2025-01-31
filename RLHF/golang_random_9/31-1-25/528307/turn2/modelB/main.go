package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// GameState holds the state of the game
type GameState struct {
	PlayerPositions map[int]int32 // Use int32 for atomic operations
	mutex           sync.Mutex
}

func NewGameState() *GameState {
	return &GameState{
		PlayerPositions: make(map[int]int32),
	}
}

// UpdatePosition updates the position of a player safely
func (gs *GameState) UpdatePosition(playerID int, position int32) {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()
	gs.PlayerPositions[playerID] = position
}

// GetState returns a copy of the current game state
func (gs *GameState) GetState() map[int]int32 {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()
	copyOfState := make(map[int]int32)
	for k, v := range gs.PlayerPositions {
		copyOfState[k] = v
	}
	return copyOfState
}

// Player represents a player in the game
type Player struct {
	ID       int
	Position int32
	mutex    sync.Mutex
}

// Move advances the player's position by one
func (p *Player) Move(state *GameState) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.Position++

	// Use atomic.StoreInt32 to safely update the game state
	atomic.StoreInt32(&state.PlayerPositions[p.ID], p.Position)
}

func main() {
	gameState := NewGameState()
	stateUpdateChan := make(chan map[int]int32)
	quitChan := make(chan int)

	// Number of players
	const numPlayers = 10
	players := make([]*Player, numPlayers)

	// Initialize players
	for i := 0; i < numPlayers; i++ {
		players[i] = &Player{ID: i, Position: 0}
		gameState.UpdatePosition(i, players[i].Position)
	}

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
				BroadcastState(gameState.GetState())
			case <-quitChan:
				return
			default:
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	// Simulate player movement concurrently
	for _, player := range players {
		go func(p *Player) {
			for {
				select {
				case <-quitChan:
					return
				default:
					p.Move(gameState)
					time.Sleep(time.Duration(100+p.ID) * time.Millisecond)
				}
			}
		}(player)
	}

	// Simulate game running for some time
	time.Sleep(10 * time.Second)
	close(quitChan)
}

// BroadcastState sends the game state to all clients
func BroadcastState(state map[int]int32) {
	fmt.Println("Broadcasting game state to clients:", state)
}
