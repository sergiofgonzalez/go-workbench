package main

import (
	"fmt"
	"time"
)

func main() {
	// Uncommemt to test with current time
	// now := time.Now()

	now := time.Date(2023, time.February, 5, 17, 1, 0, 0, time.Local)

	switch {
	case now.Hour() < 12:
		fmt.Println("Good morning!")
	case now.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}