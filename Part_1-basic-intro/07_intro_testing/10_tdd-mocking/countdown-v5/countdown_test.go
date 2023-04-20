package main

import (
	"bytes"
	"testing"
	"time"

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


type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
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

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := SpyTime{}
	sleeper := ConfigurableTimeSleeper{
		duration: sleepTime,
		sleep: spyTime.Sleep,
	}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v, but slept for %v", sleepTime, spyTime.durationSlept)
	}
}