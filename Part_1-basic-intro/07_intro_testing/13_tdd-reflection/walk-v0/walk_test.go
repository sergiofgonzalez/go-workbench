package main

import (
	"testing"
)

func TestWalk(t *testing.T) {
	x := struct {
		Name string
	}{"Jason Isaacs"}

	var got []string

	walk(x, func(fld string) {
		got = append(got, fld)
	})

	if len(got) != 1 {
		t.Errorf("expected num fields to be %d, but was %d", 1, len(got))
	}
	if got[0] != "Jason Isaacs" {
		t.Errorf("expected %q, but was %q", "Jason Isaacs", got[0])
	}
}