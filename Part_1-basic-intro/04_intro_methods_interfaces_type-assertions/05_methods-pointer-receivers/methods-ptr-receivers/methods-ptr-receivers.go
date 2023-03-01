package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale(s float64) {
	v.X *= s
	v.Y *= s
}

func main() {
	v := Vertex{4, 3}
	fmt.Println(v.Abs())
	v.Scale(2)
	fmt.Println(v)
}