package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestAppendInt(t *testing.T) {
	// tests := map[string]struct{
	// 	x []int
	// 	y []int
	// 	expected []int
	// }{
	// 	"single element on empty": {
	// 		x: []int{},
	// 		y: []int{0},
	// 		expected: []int{0},
	// 	},
	// 	"single element on non empty": {
	// 		x: []int{0},
	// 		y: []int{1},
	// 		expected: []int{0, 1},
	// 	},
	// 	"several elements on empty": {
	// 		x: []int{},
	// 		y: []int{0, 1},
	// 		expected: []int{0, 1},
	// 	},
	// 	"several elements on non empty": {
	// 		x: []int{0, 1},
	// 		y: []int{2, 3},
	// 		expected: []int{0, 1, 2, 3},
	// 	},
	// 	"whole array on non empty": {
	// 		x: []int{0, 1, 2},
	// 		y: []int{3, 4, 5},
	// 		expected: []int{0, 1, 2, 3, 4, 5},
	// 	},
	// }

	// testcase 1: single elem
	t.Run("single element on empty slice", func(t *testing.T) {
		var x []int
		got := appendInt(x, 0)
		expected := []int{0}
		if !slices.Equal(got, expected) {
			t.Errorf("got %v, but expected %v", got, expected)
		}
	})

	t.Run("single element on non empty slice", func(t *testing.T) {
		x := []int{0}
		got := appendInt(x, 1)
		expected := []int{0, 1}
		if !slices.Equal(got, expected) {
			t.Errorf("got %v, but expected %v", got, expected)
		}
	})

	t.Run("multiple elements on empty slice", func(t *testing.T) {
		var x []int
		got := appendInt(x, 0, 1)
		expected := []int{0, 1}
		if !slices.Equal(got, expected) {
			t.Errorf("got %v, but expected %v", got, expected)
		}
	})

	t.Run("multiple elements on non-empty slice", func(t *testing.T) {
		x := []int{0, 1}
		got := appendInt(x, 2, 3)
		expected := []int{0, 1, 2, 3}
		if !slices.Equal(got, expected) {
			t.Errorf("got %v, but expected %v", got, expected)
		}
	})

	t.Run("whole slice on empty slice", func(t *testing.T) {
		var x []int
		y := []int{0, 1, 2, 3}
		got := appendInt(x, y...)
		expected := []int{0, 1, 2, 3}
		if !slices.Equal(got, expected) {
			t.Errorf("got %v, but expected %v", got, expected)
		}
	})

	t.Run("whole slice on non-empty slice", func(t *testing.T) {
		x := []int{0, 1}
		y := []int{2, 3, 4}
		got := appendInt(x, y...)
		expected := []int{0, 1, 2, 3, 4}
		if !slices.Equal(got, expected) {
			t.Errorf("got %v, but expected %v", got, expected)
		}
	})
}