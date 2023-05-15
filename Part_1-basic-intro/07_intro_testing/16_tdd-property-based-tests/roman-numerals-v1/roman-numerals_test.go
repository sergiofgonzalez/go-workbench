package main

import (
	"testing"
)

func TestConvertToRoman(t *testing.T) {
	tests := map[string]struct {
		n    int
		want string
	}{
		"Convert 1 to I": {
			n: 1,
			want: "I",
		},
		"Convert 2 to II": {
			n: 2,
			want: "II",
		},
		"Convert 3 to III": {
			n: 3,
			want: "III",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := ConvertToRoman(testCase.n)
			if got != testCase.want {
				t.Errorf("%s: got %q, but wanted %q", testName, got, testCase.want)
			}
		})
	}
}
