package main

import (
	"fmt"
	"strings"
)

func main() {
	run ("the name", func(s string) {
		fmt.Println("\t>>", s)
	})

	run("another name", func(s string) {
		fmt.Println(strings.ToUpper(s))
	})
}


func run(name string, fn func(s string)) {
	fmt.Println("In run")
	fn(name)
	fmt.Println("Exiting run")
}