package shapes_test

import (
	"math"
	"testing"

	"example.com/06_tdd-structs-methods-interfaces/shapes"
)

func TestPerimeter(t *testing.T) {
	t.Run("rectangle perimeter", func(t *testing.T) {
		rect := shapes.Rectangle{
			Width:  10.0,
			Height: 5.0,
		}
		got := rect.Perimeter()
		want := 30.0

		if got != want {
			t.Errorf("got %.2f, but wanted %.2f", got, want)
		}
	})

	t.Run("circle perimeter", func(t *testing.T) {
		circle := shapes.Circle{
			Radius:  5.0,
		}
		got := circle.Perimeter()
		want := 2 * math.Pi * 5.0

		if got != want {
			t.Errorf("got %.2f, but wanted %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("rectangle area", func(t *testing.T) {
		rect := shapes.Rectangle{
			Width: 10.0,
			Height: 5.0,
		}
		got := rect.Area()
		want := 50.0

		if got != want {
			t.Errorf("got %.2f, but wanted %.2f", got, want)
		}
	})

	t.Run("circle area", func(t *testing.T) {
		circle := shapes.Circle{
			Radius: 5.0,
		}
		got := circle.Area()
		want := math.Pi * 5.0 * 5.0

		if got != want {
			t.Errorf("got %.2f, but wanted %.2f", got, want)
		}
	})
}
