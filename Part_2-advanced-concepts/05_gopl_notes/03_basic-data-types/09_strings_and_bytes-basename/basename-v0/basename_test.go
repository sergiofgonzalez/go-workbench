package main

import (
	"testing"
)

func TestBasename(t *testing.T) {
	tests := map[string]struct {
		s        string
		expected string
	}{
		"path to file with ext": {
			s: "a/b/c.go",
			expected: "c",
		},
		"filename with dot and ext": {
				s: "c.d.go",
				expected: "c.d",
		},
		"file with no ext": {
			s: "abc",
			expected: "abc",
		},
		"filename wo dot and ext": {
			s: "def.go",
			expected: "def",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := basename(testCase.s)
			if got != testCase.expected {
				t.Errorf("%s: expected %q, but got %q", testName, testCase.expected, got)
			}
		})
	}
}
