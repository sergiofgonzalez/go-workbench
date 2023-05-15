package main

import (
	"fmt"

	"example.com/shapes/shapes/rectangle"
)

func main() {
	r := rectangle.New(3, 4)
	fmt.Println(r)

	enlarged := r.Scale(2)
	fmt.Println(enlarged)

	enlarged = enlarged.Scale(0.5)

	fmt.Println(r == enlarged) 			// not the same address
	fmt.Println(r.Equal(enlarged))	// but have the same value

	// like a static function
	square := rectangle.Square(4)
	fmt.Println(square)	// but it is a Rectangle underneath
}