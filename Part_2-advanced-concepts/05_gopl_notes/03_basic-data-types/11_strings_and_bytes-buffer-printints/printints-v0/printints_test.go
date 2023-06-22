package main

import "testing"

func TestIntsToString(t *testing.T) {
	tests := map[string]struct {
		ints     []int
		expected string
	}{
		"zero value": {
			expected: "[]",
		},
		"empty array": {
			ints: []int{},
			expected: "[]",
		},
		"single elem": {
			ints: []int{1},
			expected: "[1]",
		},
		"several elems": {
			ints: []int{1, 2, 3},
			expected: "[1, 2, 3]",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := intsToString(testCase.ints)
			if got != testCase.expected {
				t.Errorf("%s: expected %q, but got %q", testName, testCase.expected, got)
			}
		})
	}
}
