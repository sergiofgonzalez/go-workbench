package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}


// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of predefined greeting messages. The
// returned message is selected randomly
func randomFormat() string {
	// A slide of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Hello to %v!",
		"Great to see you %v",
	}

	// Return randomly selected message
	return formats[rand.Intn(len(formats))]
}
