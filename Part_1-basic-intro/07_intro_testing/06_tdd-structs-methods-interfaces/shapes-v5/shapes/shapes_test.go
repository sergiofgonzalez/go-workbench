package shapes_test

import (
	"math"
	"testing"

	"example.com/06_tdd-structs-methods-interfaces/shapes"
)

func TestPerimeter(t *testing.T) {
	tests := map[string]struct {
		shape shapes.Shape
		want  float64
	}{
		"rectangle perimeter": {
			shape: &shapes.Rectangle{
				Width:  10.0,
				Height: 5.0,
			},
			want: 30.0,
		},
		"circle perimeter": {
			shape: &shapes.Circle{
				Radius: 5.0,
			},
			want: 2 * math.Pi * 5.0,
		},
	}

	for testName, testCase := range tests {
		got := testCase.shape.Perimeter()
		if got != testCase.want {
			t.Errorf("%s: got %.2f, but wanted %.2f", testName, got, testCase.want)
		}
	}
}

func TestArea(t *testing.T) {
	tests := map[string]struct {
		shape shapes.Shape
		want  float64
	}{
		"rectangle area": {
			shape: &shapes.Rectangle{
				Width:  10.0,
				Height: 5.0,
			},
			want: 50.0,
		},
		"circle area": {
			shape: &shapes.Circle{
				Radius: 5.0,
			},
			want: math.Pi * 5.0 * 5.0,
		},
		"triangle area": {
			shape: &shapes.Triangle{
				Base: 10.0,
				Height: 5.0,
			},
			want: 25.0,
		},
	}

	for testName, testCase := range tests {
		got := testCase.shape.Area()
		if got != testCase.want {
			t.Errorf("%s: got %.2f, but wanted %.2f", testName, got, testCase.want)
		}
	}
}
