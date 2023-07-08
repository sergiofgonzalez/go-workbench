package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestReverse(t *testing.T) {
	tests := map[string]struct{
		s []int
		expected []int
	}{
		"nil slice": {
			expected: nil,
		},
		"empty slice": {
			s: []int{},
			expected: []int{},
		},
		"one elem": {
			s: []int{1},
			expected: []int{1},
		},
		"two elems": {
			s: []int{1, 2},
			expected: []int{2, 1},
		},
		"three elems": {
			s: []int{1, 2, 3},
			expected: []int{3, 2, 1},
		},
		"four elems": {
			s: []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			reverse(testCase.s)
			if !slices.Equal(testCase.s, testCase.expected) {
				t.Errorf("%s: expected %v, but got %v", testName, testCase.expected, testCase.s)
			}
		})
	}
}