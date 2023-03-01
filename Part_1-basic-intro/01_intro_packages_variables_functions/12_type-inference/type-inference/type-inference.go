package main

import "fmt"

func main() {
	x := 3
	y := 5.6789
	z := 1 + 3i

	fmt.Printf("%T\n%T\n%T\n", x, y, z)
}