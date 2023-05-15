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
		"Convert 9 to IX": {
			n: 9,
			want: "IX",
		},
		"Convert 10 to X": {
			n: 10,
			want: "X",
		},
		"Convert 14 to XIV": {
			n: 14,
			want: "XIV",
		},
		"Convert 18 to XVIII": {
			n: 18,
			want: "XVIII",
		},
		"Convert 20 to XX": {
			n: 20,
			want: "XX",
		},
		"Convert 30 to XXX": {
			n: 30,
			want: "XXX",
		},
		"Convert 39 to XXXIX": {
			n: 39,
			want: "XXXIX",
		},
		"Convert 40 to XL": {
			n: 40,
			want: "XL",
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
