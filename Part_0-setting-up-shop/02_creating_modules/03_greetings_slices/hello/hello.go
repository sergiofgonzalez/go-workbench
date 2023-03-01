package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set the properties of the predefined Logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0);		// disable time, source file and line no

	// Request a greeting message and print it
	message, err := greetings.Hello("Jason Isaacs")

	// If an error is returned we call log.Fatal which will print
	// the error and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// Print the message if no error was returned
	fmt.Println(message)
}