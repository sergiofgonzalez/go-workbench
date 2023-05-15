package main

import (
	"testing"
)


func TestConvertToRoman(t *testing.T) {
	want := "I"
	got := ConvertToRoman(1)

	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}