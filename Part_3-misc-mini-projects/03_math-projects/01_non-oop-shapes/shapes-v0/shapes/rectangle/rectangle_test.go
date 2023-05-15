package rectangle_test

import (
	"testing"

	"example.com/shapes/shapes/rectangle"
)

func TestNew(t *testing.T) {
	r := rectangle.New(3, 4)
	wantWidth := 3.0
	wantHeight := 4.0

	if r == nil {
		t.Error("constructor expected to return a rectangle but was nil")
	}
	if r.Width != wantWidth {
		t.Errorf("expected width to be %.2f, but was %2.f", wantWidth, r.Width)
	}
	if r.Height != wantHeight {
		t.Errorf("expected height to be %.2f, but was %2.f", wantHeight, r.Height)
	}
}

func TestArea(t *testing.T) {
	r := rectangle.New(3, 4)
	want := 12.0

	got := r.Area()

	if got != want {
		t.Errorf("expected area to be %.2f, but was %2.f", want, got)
	}
}

func TestEquals(t *testing.T) {
	t.Run("equality", func(t *testing.T) {
		r1 := rectangle.New(3, 4)
		r2 := rectangle.New(3, 4)

		if !r1.Equal(r2) {
			t.Errorf("expected rectangles to be equal but weren't: %v != %v", r1, r2)
		}
	})

	t.Run("inequality", func(t *testing.T) {
		r1 := rectangle.New(3, 4)
		r2 := rectangle.New(4, 3)

		if r1.Equal(r2) {
			t.Errorf("expected rectangles to be equal but weren't: %v != %v", r1, r2)
		}
	})
}

func TestScale(t *testing.T) {
	r := rectangle.New(3, 4)
	want := rectangle.New(6, 8)

	got := r.Scale(2.0)
	if !got.Equal(want) {
		t.Errorf("expected %v, but got %v", want, got)
	}
}

func TestSquare(t *testing.T) {
	sq := rectangle.Square(4)
	want := rectangle.New(4, 4)

	if sq.Equal(want) {
		t.Errorf("expected %v to be equal to %v but wasn't", sq, want)
	}
}