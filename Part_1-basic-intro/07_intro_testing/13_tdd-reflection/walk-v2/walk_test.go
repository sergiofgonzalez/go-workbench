package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestWalk(t *testing.T) {
	tests := map[string]struct {
		input         interface{}
		expectedCalls []string
	}{
		"struct with one string field": {
			input: struct{
				Name string
			}{"Jason Isaacs"},
			expectedCalls: []string{"Jason Isaacs"},
		},
		"strict with two string fields": {
			input: struct{
				Name string
				City string
			}{"Jason Isaacs", "London"},
			expectedCalls: []string{"Jason Isaacs", "London"},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			var got []string
			walk(testCase.input, func(fld string) {
				got = append(got, fld)
			})
			if !slices.Equal(got, testCase.expectedCalls) {
				t.Errorf("want %v, but got %v", testCase.expectedCalls, got)
			}
		})
	}
}
