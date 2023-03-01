package main

import (
	"fmt"

	"example.com/01_hello-go/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
}