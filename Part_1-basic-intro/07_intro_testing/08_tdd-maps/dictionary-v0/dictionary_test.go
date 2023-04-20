package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{
		"test": "this is the value associated to test",
	}

	got := Search(dictionary, "test")
	want := "this is the value associated to test"
	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}