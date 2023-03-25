package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const wordLength = 5

// Game holds all the information needed to play a game of Gordle
type Game struct {
	reader *bufio.Reader
}

// New returns an initialized value for the game
func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}

	return g
}

// Play initiates a game of Gordle
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	fmt.Printf("Enter a guess: ")
}

// ask reads input from the user until a valid suggestion is made
func (g *Game) ask() []rune {
	fmt.Printf("Type your guess (%d characters): ", wordLength)

	for {
		line, _, err := g.reader.ReadLine()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid input: %v", err)
			continue
		}
		guess := []rune(string(line))

		if len(guess) != wordLength {
			fmt.Fprintf(os.Stderr, "Invalid length, guess must be %d characters long", wordLength)
			continue
		}
		return guess
	}
}
