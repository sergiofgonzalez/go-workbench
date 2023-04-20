package countdown

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	want := "3"

	Countdown(&buffer)
	got := buffer.String()
	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}