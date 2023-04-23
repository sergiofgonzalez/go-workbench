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
		"struct with two string fields": {
			input: struct{
				Name string
				City string
			}{"Jason Isaacs", "London"},
			expectedCalls: []string{"Jason Isaacs", "London"},
		},
		"struct with non string field": {
			input: struct{
				Name string
				age int
			}{"Jason Isaacs", 58},
			expectedCalls: []string{"Jason Isaacs"},
		},
		"struct with nested fields": {
			input: struct{
				Name string
				Profile struct {
					Age int
					City string
				}
			}{
				Name: "Jason Isaacs",
				Profile: struct{
					Age int
					City string
					}{
					Age: 58,
					City: "London",
				},
			},
			expectedCalls: []string{"Jason Isaacs", "London"},
		},
		"pointer to a struct with nested fields": {
			input: &struct{
				Name string
				Profile struct {
					Age int
					City string
				}
			}{
				Name: "Jason Isaacs",
				Profile: struct{
					Age int
					City string
					}{
					Age: 58,
					City: "London",
				},
			},
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
