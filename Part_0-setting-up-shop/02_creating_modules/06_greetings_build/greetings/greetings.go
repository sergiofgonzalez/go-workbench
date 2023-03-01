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

	// Comment/Uncomment to see a failing test
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat())

	return message, nil
}

//Hellos returns a map that associates each of the named people with a greeting
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	// Loop through the received slice of names, calling the Hello()
	// function to get a greeting for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// If everything went well, associate the retrieved message with the name
		messages[name] = message
	}

	return messages, nil
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
