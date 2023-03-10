package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(s float64) {
	v.X *= s
	v.Y *= s
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
	v.Scale(2)
	fmt.Println(v)
	fmt.Println(v.Abs())
}