package gordle

import "fmt"

// Game holds all the information needed to play a game of Gordle
type Game struct{}

// New returns an initialized value for the game
func New() *Game {
	g := &Game{}

	return g
}

// Play initiates a game of Gordle
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	fmt.Printf("Enter a guess: ")
}