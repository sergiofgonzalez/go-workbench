package main

import (
	"testing"
)

func TestComma(t *testing.T) {
	tests := map[string]struct {
		s        string
		expected string
	}{
		"number less than 999": {
			s: "999",
			expected: "999",
		},
		"number between 1000 and 999999 (1)": {
				s: "1000",
				expected: "1,000",
		},
		"number between 1000 and 999999 (2)": {
			s: "999999",
			expected: "999,999",
	},
		"number in the millions": {
			s: "1000000",
			expected: "1,000,000",
		},
		"big num": {
			s: "1234567890",
			expected: "1,234,567,890",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := comma(testCase.s)
			if got != testCase.expected {
				t.Errorf("%s: expected %q, but got %q", testName, testCase.expected, got)
			}
		})
	}
}
