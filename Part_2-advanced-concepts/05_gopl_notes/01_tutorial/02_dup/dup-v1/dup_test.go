package dup_test

import (
	"bytes"
	"strings"
	"testing"

	"example.com/dup"
	"golang.org/x/exp/maps"
)

func TestGetCounts(t *testing.T) {
	tests := map[string]struct {
		input  string
		wanted map[string]int
	}{
		"empty input": {
			input:  "",
			wanted: make(map[string]int),
		},
		"single line": {
			input: "foo",
			wanted: map[string]int{
				"foo": 1,
			},
		},
		"two different lines": {
			input: "foo\nbar",
			wanted: map[string]int{
				"foo": 1,
				"bar": 1,
			},
		},
		"two identical lines": {
			input: "foobar\nfoobar",
			wanted: map[string]int{
				"foobar": 2,
			},
		},
		"several lines": {
			input: "Foobar\nFoo\nBar\nFoo\n",
			wanted: map[string]int{
				"Foobar": 1,
				"Foo":    2,
				"Bar":    1,
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			r := strings.NewReader(testCase.input)
			gotCounts := make(map[string]int)
			dup.ConfigurableGetCounts(r, gotCounts)
			if !maps.Equal(gotCounts, testCase.wanted) {
				t.Errorf("expected counts to be %v, but was %v", testCase.wanted, gotCounts)
			}
		})
	}
}

func TestWriteDups(t *testing.T) {
	tests := map[string]struct {
		counts map[string]int
		wanted string
	}{
		"no dups empty counts": {
			counts: make(map[string]int),
			wanted: "",
		},
		"no dups nil counts": {
			counts: nil,
			wanted: "",
		},
		"no dups": {
			counts: map[string]int{
				"foo":    1,
				"bar":    1,
				"foobar": 1,
			},
			wanted: "",
		},
		"single dup": {
			counts: map[string]int{
				"foo":    2,
				"bar":    1,
				"foobar": 1,
			},
			wanted: "2\tfoo\n",
		},
		"several dups": {
			// this will make a flaky test because maps do not guarantee order
			counts: map[string]int{
				"foo":    1,
				"bar":    2,
				"foobar": 3,
			},
			wanted: "2\tbar\n3\tfoobar\n",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			buffer := bytes.Buffer{}

			dup.ConfigurableWriteDups(&buffer, testCase.counts)
			got := buffer.String()

			if got != testCase.wanted {
				t.Errorf("wanted %#v, but got %#v", testCase.wanted, got)
			}
		})
	}
}
