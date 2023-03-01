package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	p := Point{3, 2}
	fmt.Println(p)

	p.X *= 2
	fmt.Printf("p=(%v, %v)", p.X, p.Y)
}