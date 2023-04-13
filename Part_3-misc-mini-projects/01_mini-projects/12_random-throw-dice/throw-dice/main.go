package main

import (
	"fmt"

	"example.com/12_random-throw-dice/experiment"
)


func main() {
	e := experiment.New(10000, 2, 6)
	err := e.Run()
	if err != nil {
		fmt.Printf("Could not complete: %v\n", err)
		return
	}
	e.Show()
}

