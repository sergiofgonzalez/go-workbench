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
		"Convert 4 to IV": {
			n: 4,
			want: "IV",
		},
		"Convert 5 to V": {
			n: 5,
			want: "V",
		},
		"Convert 6 to VI": {
			n: 6,
			want: "VI",
		},
		"Convert 7 to VII": {
			n: 7,
			want: "VII",
		},
		"Convert 8 to VIII": {
			n: 8,
			want: "VIII",
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
