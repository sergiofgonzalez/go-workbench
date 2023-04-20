package shapes_test

import (
	"math"
	"testing"

	"example.com/06_tdd-structs-methods-interfaces/shapes"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, s shapes.Shape, want float64) {
		t.Helper()
		got := s.Perimeter()
		if got != want {
			t.Errorf("got %.2f, but wanted %.2f", got, want)
		}
	}


	t.Run("rectangle perimeter", func(t *testing.T) {
		rect := shapes.Rectangle{
			Width:  10.0,
			Height: 5.0,
		}
		checkPerimeter(t, &rect, 30.0)
	})

	t.Run("circle perimeter", func(t *testing.T) {
		circle := shapes.Circle{
			Radius:  5.0,
		}
		checkPerimeter(t, &circle, 2 * math.Pi * 5.0)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, s shapes.Shape, want float64) {
		t.Helper()
		got := s.Area()
		if got != want {
			t.Errorf("got %.2f, but wanted %.2f", got, want)
		}
	}

	t.Run("rectangle area", func(t *testing.T) {
		rect := shapes.Rectangle{
			Width: 10.0,
			Height: 5.0,
		}
		checkArea(t, &rect, 50.0)
	})

	t.Run("circle area", func(t *testing.T) {
		circle := shapes.Circle{
			Radius: 5.0,
		}
		checkArea(t, &circle, math.Pi * 5.0 * 5.0)
	})
}
