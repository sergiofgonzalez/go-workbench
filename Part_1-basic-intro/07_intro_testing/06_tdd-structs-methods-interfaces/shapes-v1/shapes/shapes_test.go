package shapes_test

import (
	"testing"

	"example.com/06_tdd-structs-methods-interfaces/shapes"
)

func TestPerimeter(t *testing.T) {
	rect := shapes.Rectangle{
		Width:  10.0,
		Height: 5.0,
	}
	got := shapes.Perimeter(rect)
	want := 30.0

	if got != want {
		t.Errorf("got %.2f, but wanted %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rect := shapes.Rectangle{
		Width: 10.0,
		Height: 5.0,
	}
	got := shapes.Area(rect)
	want := 50.0

	if got != want {
		t.Errorf("got %.2f, but wanted %.2f", got, want)
	}
}
