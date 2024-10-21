package main

import (
	"fmt"
	"math/rand"
	"time"
)

// LazyValue represents a value that will be computed lazily.
type LazyValue struct {
	compute func() int
	value   int
	done    bool
}

// Get retrieves the computed value of the LazyValue.
func (l *LazyValue) Get() int {
	if !l.done {
		l.value = l.compute()
		l.done = true
	}
	return l.value
}

// newLazyValue creates a new LazyValue that will compute the value using the given function.
func newLazyValue(compute func() int) *LazyValue {
	return &LazyValue{compute: compute}
}

// API represents a mock API for retrieving game data.
type API struct{}

// GetScore fetches the player's score from the API.
func (a *API) GetScore() int {
	// Simulate an expensive API call.
	time.Sleep(time.Second)
	return rand.Intn(1000)
}

// Game represents the main game logic.
type Game struct {
	api        *API
	score      *LazyValue
	highScore  *LazyValue
	difficulty *LazyValue
}

// NewGame creates a new Game instance.
func NewGame() *Game {
	api := &API{}
	return &Game{
		api:        api,
		score:      newLazyValue(api.GetScore),
		highScore:  newLazyValue(func() int { return api.GetScore() }),
		difficulty: newLazyValue(func() int { return 5 }),
	}
}

// Update updates the game state.
func (g *Game) Update() {
	// Update the game logic here.
	// For demonstration purposes, we'll update the difficulty randomly.
	if rand.Intn(10) == 0 {
		g.difficulty = newLazyValue(func() int { return rand.Intn(10) + 1 })
	}
}

// GetScore retrieves the player's current score.
func (g *Game) GetScore() int {
	return g.score.Get()
}

// GetHighScore retrieves the player's high score.
func (g *Game) GetHighScore() int {
	return g.highScore.Get()
}

// GetDifficulty retrieves the current game difficulty.
func (g *Game) GetDifficulty() int {
	return g.difficulty.Get()
}

func main() {
	game := NewGame()

	fmt.Println("Initial Score:", game.GetScore())           // Will fetch score from API lazily
	fmt.Println("Initial High Score:", game.GetHighScore())  // Will fetch high score from API lazily
	fmt.Println("Initial Difficulty:", game.GetDifficulty()) // Will return the default difficulty lazily

	// Simulate game update loop
	for i := 0; i < 10; i++ {
		game.Update()
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Updated Score:", game.GetScore())
		fmt.Println("Updated High Score:", game.GetHighScore())
		fmt.Println("Updated Difficulty:", game.GetDifficulty())
	}
}
