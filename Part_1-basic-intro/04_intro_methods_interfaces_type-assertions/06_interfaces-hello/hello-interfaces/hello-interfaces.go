package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}


func main() {
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	var a Abser

	a = f		// OK, MyFloat type implements Abser because it provides an `Abs() float64` method
	a = &v	// OK, *Vertex type implements Absert

	// a = v   // Error: Vertex type does not implement Absert

	fmt.Println(a.Abs())
}