package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Printf("Type a word: ")
	input, err := readInput()
	if err != nil {
		fmt.Printf("Program failed: %v\n", err)
		if  errors.Is(err, errInvalidLength) {
			fmt.Println("\t>> the wrapped error was because the string was an invalid length")
		}
		return
	}

	fmt.Printf("You entered: %v\n", input)
}

func readInput() (string, error) {
	reader  := bufio.NewReader(os.Stdin)
	lineBytes, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}

	input := string(lineBytes)
	err = checkLen(input)
	if err != nil {
		return "", fmt.Errorf("Could not read input: checkLen failed: %w\n", err)
	}

	firstRune := rune(input[0])
	if unicode.IsLower(firstRune) {
		return "", fmt.Errorf("Could not read input: first character is not a capital letter\n")
	}

	return input, nil
}


var errInvalidLength error = fmt.Errorf("string has an invalid length")

func checkLen(str string) error {
	if len(str) != 5 {
		return fmt.Errorf("invalid length: expected 5, got %v: %w\n", len(str), errInvalidLength)
	}
	return nil
}