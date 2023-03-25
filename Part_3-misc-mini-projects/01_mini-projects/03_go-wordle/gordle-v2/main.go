package main

import (
	"fmt"
	"os"

	"example.com/gordle/gordle"
)

func main() {
	g := gordle.New(os.Stdin)

	g.Play()

	fmt.Println("\n-- done!")
}