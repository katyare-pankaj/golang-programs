package main

import (
	"fmt"
	"sync"
)

// define a callback function type
type Callback func()

// game state struct
type GameState struct {
	playerPositions map[string]int
	mux             sync.Mutex
	callbacks       []Callback
}

// Initialize game state
func NewGameState() *GameState {
	return &GameState{
		playerPositions: make(map[string]int),
		callbacks:       make([]Callback, 0),
	}
}

// Register a callback function
func (g *GameState) RegisterCallback(cb Callback) {
	g.mux.Lock()
	defer g.mux.Unlock()
	g.callbacks = append(g.callbacks, cb)
}

// Update game state and notify all callbacks
func (g *GameState) UpdatePlayerPosition(playerID string, newPosition int) {
	g.mux.Lock()
	defer g.mux.Unlock()
	g.playerPositions[playerID] = newPosition
	for _, cb := range g.callbacks {
		cb()
	}
}

// Example callback function that prints the updated game state
func printGameStateCallback(gs *GameState) Callback {
	return func() {
		gs.mux.Lock()
		defer gs.mux.Unlock()
		fmt.Println("Updated Game State:")
		for playerID, position := range gs.playerPositions {
			fmt.Printf("Player %s: Position %d\n", playerID, position)
		}
	}
}
func main() {
	gs := NewGameState()
	// Register the callback function
	gs.RegisterCallback(printGameStateCallback(gs))
	// Update game state from multiple goroutines
	go func() {
		for i := 0; i < 5; i++ {
			gs.UpdatePlayerPosition("Player1", i)
		}
	}()
	go func() {
		for i := 0; i < 3; i++ {
			gs.UpdatePlayerPosition("Player2", i)
		}
	}()
	// Wait for the game to finish
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		<-make(chan struct{}) // Wait for player1 updates to finish
	}()
	go func() {
		defer wg.Done()
		<-make(chan struct{}) // Wait for player2 updates to finish
	}()
	wg.Wait()
	fmt.Println("Game finished!")
}
