package main

import (
	"fmt"
	"os"

	"example.com/gordle/gordle"
)

func main() {
	corpus, err := gordle.ReadCorpus("./corpus/english.txt")
	if err != nil {
		panic(err)
	}

	g, err := gordle.New(os.Stdin, corpus, 5)
	if err != nil {
		panic(err)
	}

	g.Play()

	fmt.Println("\n-- done!")
}