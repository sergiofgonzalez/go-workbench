package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it
	message := greetings.Hello("Jason Isaacs")
	fmt.Println(message)
}