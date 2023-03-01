package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func Scale(v *Vertex, s float64) {
	v.X *= s
	v.Y *= s
}


func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
	Scale(&v, 2)
	fmt.Println(v)
	fmt.Println(Abs(v))
}