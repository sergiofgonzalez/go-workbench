package stringfns_test

import (
	"testing"

	"example.com/stringfns"
)

func TestHasPrefix(t *testing.T) {
	tests := map[string]struct{
		s string
		prefix string
		expected bool
	}{
		"prefix found": {
			s: "foobar",
			prefix: "foo",
			expected: true,
		},
		"prefix not found": {
			s: "foobar",
			prefix: "fox",
			expected: false,
		},
		"prefix equal s": {
			s: "foo",
			prefix: "foo",
			expected: true,
		},
		"prefix larger than s": {
			s: "foo",
			prefix: "foos",
			expected: false,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := stringfns.HasPrefix(testCase.s, testCase.prefix)
			if got != testCase.expected {
				t.Errorf("%s: got %v, expected %v", testName, got, testCase.expected)
			}
		})
	}
}

func TestHasSuffix(t *testing.T) {
	tests := map[string]struct{
		s string
		suffix string
		expected bool
	}{
		"prefix found": {
			s: "foobar",
			suffix: "bar",
			expected: true,
		},
		"suffix not found": {
			s: "foobar",
			suffix: "baz",
			expected: false,
		},
		"suffix equal s": {
			s: "foo",
			suffix: "foo",
			expected: true,
		},
		"suffix larger than s": {
			s: "foo",
			suffix: "sfoo",
			expected: false,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := stringfns.HasSuffix(testCase.s, testCase.suffix)
			if got != testCase.expected {
				t.Errorf("%s: got %v, expected %v", testName, got, testCase.expected)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := map[string]struct{
		s string
		substr string
		expected bool
	}{
		"prefix found": {
			s: "foobarbaz",
			substr: "foo",
			expected: true,
		},
		"mid str found": {
			s: "foobarbaz",
			substr: "bar",
			expected: true,
		},
		"suffix str found": {
			s: "foobarbaz",
			substr: "baz",
			expected: true,
		},
		"suffix not found": {
			s: "foobar",
			substr: "ooz",
			expected: false,
		},
		"equals": {
			s: "foo",
			substr: "foo",
			expected: true,
		},
		"substr larger than s": {
			s: "foo",
			substr: "foos",
			expected: false,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := stringfns.Contains(testCase.s, testCase.substr)
			if got != testCase.expected {
				t.Errorf("%s: got %v, expected %v", testName, got, testCase.expected)
			}
		})
	}
}