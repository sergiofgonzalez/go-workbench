package shapes_test

import (
	"testing"

	"example.com/06_tdd-structs-methods-interfaces/shapes"
)

func TestPerimeter(t *testing.T) {
	want := 30.0
	got := shapes.Perimeter(10.0, 5.0)
	if got != want {
		t.Errorf("got %.2f, but wanted %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	want := 50.0
	got := shapes.Area(10.0, 5.0)
	if got != want {
		t.Errorf("got %.2f, but wanted %.2f", got, want)
	}
}