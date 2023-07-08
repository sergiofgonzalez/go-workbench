package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel

	w.X = 8
	w.Y = 10
	w.Radius = 5
	w.Spokes = 10

	fmt.Printf("%#v\n", w)

	// verbose mode also possible: field names derived from type name
	w.Circle.Point.X = 8
	w.Circle.Point.Y = 10
	w.Circle.Radius = 5
	w.Spokes = 10

	fmt.Printf("%#v\n", w)

	w = Wheel{
		Circle: Circle{
			Point: Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w)
}