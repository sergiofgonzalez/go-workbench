package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Jason Isaacs")

	got := buffer.String()
	want := "Hello to Jason Isaacs!"

	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}
