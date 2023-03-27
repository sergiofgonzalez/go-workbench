package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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

	guess := g.ask()

	fmt.Printf("Your guess is: %q\n", string(guess))
}

// ask reads input from the user until a valid suggestion is made
func (g *Game) ask() []rune {

	for {
		fmt.Printf("Type your guess (%d characters): ", wordLength)

		line, _, err := g.reader.ReadLine()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid input: %v", err)
			continue
		}
		guess := splitToUppercaseCharacters(string(line))
		err = g.validateGuess(guess)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s.\n", err.Error())
			continue
		}

		return guess
	}
}

var errInvalidLengthWord error = fmt.Errorf("invalid guess: word must have the same number of characters of the solution")


// validateGuess takes a guess and returns nil if it matches, or an error otherwise
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != wordLength {
		return fmt.Errorf("Invalid length: expected %v, got %v, %w", wordLength, len(guess), errInvalidLengthWord)
	}

	return nil
}

// splitToUppercaseCharacters takes an input string and returns a slice of runes corresponding
// to the given string in which all the characters have been capitalized
func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}