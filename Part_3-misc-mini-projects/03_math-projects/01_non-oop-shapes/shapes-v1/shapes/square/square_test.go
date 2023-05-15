package square_test

import (
	"testing"

	"example.com/shapes/shapes/square"
)

func TestNew(t *testing.T) {
	s := square.New(3)
	wantSide := 3.0

	if s == nil {
		t.Error("constructor expected to return a square but was nil")
	}
	if s.Side != wantSide {
		t.Errorf("expected side to be %.2f, but was %2.f", wantSide, s.Side)
	}
}

func TestArea(t *testing.T) {
	s := square.New(3)
	want := 9.0

	got := s.Area()

	if got != want {
		t.Errorf("expected area to be %.2f, but was %2.f", want, got)
	}
}

func TestEquals(t *testing.T) {
	t.Run("equality", func(t *testing.T) {
		s1 := square.New(3)
		s2 := square.New(3)

		if !s1.Equal(s2) {
			t.Errorf("expected squares to be equal but weren't: %v != %v", s1, s2)
		}
	})

	t.Run("inequality", func(t *testing.T) {
		s1 := square.New(3)
		s2 := square.New(4)

		if s1.Equal(s2) {
			t.Errorf("expected squares to be equal but weren't: %v != %v", s1, s2)
		}
	})
}

func TestScale(t *testing.T) {
	s := square.New(3)
	want := square.New(6)

	got := s.Scale(2.0)
	if !got.Equal(want) {
		t.Errorf("expected %v, but got %v", want, got)
	}
}