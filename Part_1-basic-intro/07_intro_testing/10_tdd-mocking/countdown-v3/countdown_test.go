package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	spySleeper := SpySleeper{}
	want := `3
2
1
Go!
`
	wantCalls := 3

	Countdown(&buffer, &spySleeper)
	got := buffer.String()
	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
	if spySleeper.Calls != wantCalls {
		t.Errorf("unexpected calls: got %d, but wanted %d", spySleeper.Calls, wantCalls)
	}
}