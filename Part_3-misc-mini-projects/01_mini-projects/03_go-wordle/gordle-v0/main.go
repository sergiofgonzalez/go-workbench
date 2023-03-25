package main

import (
	"fmt"

	"example.com/gordle/gordle"
)

func main() {
	g := gordle.New()

	g.Play()

	fmt.Println("\n-- done!")
}