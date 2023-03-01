package main

import "fmt"

func main() {
	actors := [4]string{
		"Isaacs",
		"Pugh",
		"Elba",
		"Robbie",
	}

	a := actors[0:2]
	b := actors[1:3]

	b[0] = "Watson"

	fmt.Println(a, b)
	fmt.Println(actors)
}