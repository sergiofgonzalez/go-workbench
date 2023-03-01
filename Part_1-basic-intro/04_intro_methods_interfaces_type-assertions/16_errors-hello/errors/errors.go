package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %q", e.When, e.What)
}

func run() *MyError {
	return &MyError{
		time.Now(),
		"A fabricated error",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println("An error was found calling run():", err)
	}
}