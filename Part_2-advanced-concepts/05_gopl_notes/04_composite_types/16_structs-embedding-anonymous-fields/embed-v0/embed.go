package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 10
	w.Circle.Radius = 5
	w.Spokes = 10

	fmt.Printf("%#v\n", w)

	w = Wheel{
		Circle: Circle{
			Center: Point{
				X: 5,
				Y: 5,
			},
			Radius: 10,
		},
		Spokes: 5,
	}
	fmt.Printf("%#v\n", w)
}