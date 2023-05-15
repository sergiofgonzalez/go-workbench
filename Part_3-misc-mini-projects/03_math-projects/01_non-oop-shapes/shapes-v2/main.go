package main

import (
	"fmt"

	"example.com/shapes/shapes"
	"example.com/shapes/shapes/rectangle"
	"example.com/shapes/shapes/square"
)

func main() {
	// r := rectangle.New(3, 4)
	// fmt.Println(r)

	// enlarged := r.Scale(2)
	// fmt.Println(enlarged)

	// enlarged = enlarged.Scale(0.5)

	// fmt.Println(r == enlarged) 			// not the same address
	// fmt.Println(r.Equal(enlarged))	// but have the same value

	// // like a static function
	// square := rectangle.Square(4)
	// fmt.Println(square)	// but it is a Rectangle underneath

	s := square.New(4)
	r := rectangle.New(4, 3)

	fmt.Printf("Area(%v)=%.2f\n", r, shapes.Area(s))
	fmt.Printf("Area(%v)=%.2f\n", r, shapes.Area(r))

	fmt.Printf("Equal(%v, %v)=%v\n", s, r, shapes.Equal(s, r))

	fmt.Printf("Scale(%v, 0.5)=%v\n", s, shapes.Scale(s, 0.5))
	fmt.Printf("Scale(%v, 2)=%v\n", s, shapes.Scale(r, 2))
}