package main

import (
	"bytes"
	"testing"

	"golang.org/x/exp/slices"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdown struct {
	Calls []string
}

func (s *SpyCountdown) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func (s *SpyCountdown) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return
}

func TestCountdown(t *testing.T) {
	t.Run("confirm content", func(t *testing.T) {
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
	})

	t.Run("confirm test sequence", func(t *testing.T) {
		spyCountdown := SpyCountdown{}
		want := []string{"write", "sleep", "write", "sleep", "write", "sleep", "write"}

		Countdown(&spyCountdown, &spyCountdown)
		if !slices.Equal(want, spyCountdown.Calls) {
			t.Errorf("got %v, but wanted %v", spyCountdown.Calls, want)
		}
	})
}