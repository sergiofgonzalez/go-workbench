package main

import (
	"fmt"
	"math"
)

type Vector2D struct {
	X, Y float64
}

func (v Vector2D) Dist() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := Vector2D{3, 4}

	fmt.Println(v.Dist())
}